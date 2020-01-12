// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/chewxy/math32"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/dnd"
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/oswin/mouse"
	"github.com/goki/ki/ki"
)

//go:generate stringer -type=EventPris

// EventPris for different queues of event signals, processed in priority order
type EventPris int32

const (
	// HiPri = high priority -- event receivers processed first -- can be used
	// to override default behavior
	HiPri EventPris = iota

	// RegPri = default regular priority -- most should be here
	RegPri

	// LowPri = low priority -- processed last -- typically for containers /
	// dialogs etc
	LowPri

	// LowRawPri = unfiltered (raw) low priority -- ignores whether the event
	// was already processed.
	LowRawPri

	EventPrisN

	// AllPris = -1 = all priorities (for delete cases only)
	AllPris EventPris = -1
)

// EventMgr is an event manager that handles distributing events to nodes.
// It relies on the EventMaster for a few things outside of its scope.
type EventMgr struct {
	Master          EventMaster                             `json:"-" xml:"-" view:"-" desc:"master of this event mangager -- handles broader scope issues"`
	EventSigs       [oswin.EventTypeN][EventPrisN]ki.Signal `json:"-" xml:"-" view:"-" desc:"signals for communicating each type of event, organized by priority"`
	EventMu         sync.Mutex                              `json:"-" xml:"-" view:"-" desc:"mutex that protects event sending"`
	TimerMu         sync.Mutex                              `json:"-" xml:"-" view:"-" desc:"mutex that protects timer variable updates (e.g., hover AfterFunc's)"`
	Dragging        ki.Ki                                   `json:"-" xml:"-" desc:"node receiving mouse dragging events -- not for DND but things like sliders -- anchor to same"`
	Scrolling       ki.Ki                                   `json:"-" xml:"-" desc:"node receiving mouse scrolling events -- anchor to same"`
	DNDData         mimedata.Mimes                          `json:"-" xml:"-" desc:"drag-n-drop data -- if non-nil, then DND is taking place"`
	DNDSource       ki.Ki                                   `json:"-" xml:"-" desc:"drag-n-drop source node"`
	DNDFinalEvent   *dnd.Event                              `json:"-" xml:"-" view:"-" desc:"final event for DND which is sent if a finalize is received"`
	Focus           ki.Ki                                   `json:"-" xml:"-" desc:"node receiving keyboard events -- use SetFocus, CurFocus"`
	FocusMu         sync.RWMutex                            `json:"-" xml:"-" view:"-" desc:"mutex that protects focus updating"`
	FocusStack      []ki.Ki                                 `json:"-" xml:"-" desc:"stack of focus"`
	startDrag       *mouse.DragEvent
	dragStarted     bool
	startDND        *mouse.DragEvent
	dndStarted      bool
	startHover      *mouse.MoveEvent
	curHover        *mouse.MoveEvent
	hoverStarted    bool
	hoverTimer      *time.Timer
	startDNDHover   *mouse.DragEvent
	curDNDHover     *mouse.DragEvent
	dndHoverStarted bool
	dndHoverTimer   *time.Timer
}

// WinEventRecv is used to hold info about widgets receiving event signals to
// given function, used for sorting and delayed sending.
type WinEventRecv struct {
	Recv ki.Ki
	Func ki.RecvFunc
	Data int
}

// Set sets the recv and fun
func (we *WinEventRecv) Set(r ki.Ki, f ki.RecvFunc, data int) {
	we.Recv = r
	we.Func = f
	we.Data = data
}

// Call calls the function on the recv with the args
func (we *WinEventRecv) Call(send ki.Ki, sig int64, data interface{}) {
	if EventTrace {
		fmt.Printf("calling event: %v method on: %v\n", data, we.Recv.Path())
	}
	we.Func(we.Recv, send, sig, data)
}

type WinEventRecvList []WinEventRecv

func (wl *WinEventRecvList) Add(recv ki.Ki, fun ki.RecvFunc, data int) {
	rr := WinEventRecv{recv, fun, data}
	*wl = append(*wl, rr)
}

func (wl *WinEventRecvList) AddDepth(recv ki.Ki, fun ki.RecvFunc, par ki.Ki) {
	wl.Add(recv, fun, recv.ParentLevel(par))
}

// ConnectEvent adds a Signal connection for given event type and
// priority to given receiver
func (em *EventMgr) ConnectEvent(recv ki.Ki, et oswin.EventType, pri EventPris, fun ki.RecvFunc) {
	if et >= oswin.EventTypeN {
		log.Printf("EventMgr ConnectEvent type: %v is not a known event type\n", et)
		return
	}
	em.EventSigs[et][pri].Connect(recv, fun)
}

// DisconnectEvent removes Signal connection for given event type to given
// receiver -- pri is priority -- pass AllPris for all priorities
func (em *EventMgr) DisconnectEvent(recv ki.Ki, et oswin.EventType, pri EventPris) {
	if et >= oswin.EventTypeN {
		log.Printf("EventMgr DisconnectEvent type: %v is not a known event type\n", et)
		return
	}
	if pri == AllPris {
		for p := HiPri; p < EventPrisN; p++ {
			em.EventSigs[et][p].Disconnect(recv)
		}
	} else {
		em.EventSigs[et][pri].Disconnect(recv)
	}
}

// DisconnectAllEvents disconnect node from all event signals -- pri is
// priority -- pass AllPris for all priorities
func (em *EventMgr) DisconnectAllEvents(recv ki.Ki, pri EventPris) {
	if pri == AllPris {
		for et := oswin.EventType(0); et < oswin.EventTypeN; et++ {
			for p := HiPri; p < EventPrisN; p++ {
				em.EventSigs[et][p].Disconnect(recv)
			}
		}
	} else {
		for et := oswin.EventType(0); et < oswin.EventTypeN; et++ {
			em.EventSigs[et][pri].Disconnect(recv)
		}
	}
}

// SendEventSignal sends given event signal to all receivers that want it --
// note that because there is a different EventSig for each event type, we are
// ONLY looking at nodes that have registered to receive that type of event --
// the further filtering is just to ensure that they are in the right position
// to receive the event (focus, popup filtering, etc).  If popup is true, then
// only items on popup are in scope, otherwise items NOT on popup are in scope
// (if no popup, everything is in scope).
func (em *EventMgr) SendEventSignal(evi oswin.Event, popup bool) {
	et := evi.Type()
	if et > oswin.EventTypeN || et < 0 {
		return // can't handle other types of events here due to EventSigs[et] size
	}

	em.EventMu.Lock()

	send := em.Master.EventTopNode()

	// fmt.Printf("got event type: %v\n", et)
	for pri := HiPri; pri < EventPrisN; pri++ {
		if pri != LowRawPri && evi.IsProcessed() { // someone took care of it
			continue
		}

		// we take control of signal process to sort elements by depth, and
		// dispatch to inner-most one first
		rvs := make(WinEventRecvList, 0, 10)

		esig := &em.EventSigs[et][pri]

		esig.Mu.RLock()
		for recv, fun := range esig.Cons {
			if recv.IsDestroyed() {
				// fmt.Printf("ki.Signal deleting destroyed receiver: %v type %T\n", recv.Name(), recv)
				delete(esig.Cons, recv)
				continue
			}
			if recv.IsDeleted() {
				continue
			}
			esig.Mu.RUnlock()
			cont := em.SendEventSignalFunc(evi, popup, &rvs, recv, fun)
			esig.Mu.RLock()
			if !cont {
				break
			}
		}
		esig.Mu.RUnlock()

		if len(rvs) == 0 {
			continue
		}

		// deepest first
		sort.Slice(rvs, func(i, j int) bool {
			return rvs[i].Data > rvs[j].Data
		})

		for _, rr := range rvs {
			switch evi.(type) {
			case *mouse.DragEvent:
				if em.Dragging == nil {
					rr.Recv.SetFlag(int(NodeDragging)) // PROVISIONAL!
				}
			}
			em.EventMu.Unlock()
			rr.Call(send, int64(et), evi) // could call further event loops..
			em.EventMu.Lock()
			if pri != LowRawPri && evi.IsProcessed() { // someone took care of it
				switch evi.(type) { // only grab events if processed
				case *mouse.DragEvent:
					if em.Dragging == nil {
						em.Dragging = rr.Recv
						rr.Recv.SetFlag(int(NodeDragging))
					}
				case *mouse.ScrollEvent:
					if em.Scrolling == nil {
						em.Scrolling = rr.Recv
					}
				}
				break
			} else {
				switch evi.(type) {
				case *mouse.DragEvent:
					if em.Dragging == nil {
						rr.Recv.ClearFlag(int(NodeDragging)) // clear provisional
					}
				}
			}
		}
	}
	em.EventMu.Unlock()
}

// SendEventSignalFunc is the inner loop of the SendEventSignal -- needed to deal with
// map iterator locking logic in a cleaner way.  Returns true to continue, false to break
func (em *EventMgr) SendEventSignalFunc(evi oswin.Event, popup bool, rvs *WinEventRecvList, recv ki.Ki, fun ki.RecvFunc) bool {
	nii, ni := KiToNode2D(recv)
	if ni != nil {
		if !em.Master.IsInScope(ni, popup) {
			return true
		}
		if evi.OnFocus() {
			if !nii.HasFocus2D() {
				return true
			}
			if !em.Master.IsFocusActive() { // reactivate on keyboard input
				em.Master.SetFocusActiveState(true)
				// fmt.Printf("set foc active: %v\n", ni.PathUnique())
				nii.FocusChanged2D(FocusActive)
			}
		}
	}
	top := em.Master.EventTopNode()
	// remainder is done using generic node interface, for 2D and 3D
	gni := recv.(Node)
	gn := gni.AsGiNode()
	// todo: need a focus concept for 3D
	if evi.HasPos() {
		pos := evi.Pos()
		switch evi.(type) {
		case *mouse.DragEvent:
			if em.Dragging != nil {
				if em.Dragging == gn.This() {
					rvs.Add(recv, fun, 10000)
					return false
				} else {
					return true
				}
			} else {
				if pos.In(gn.WinBBox) {
					rvs.AddDepth(recv, fun, top)
					return false
				}
				return true
			}
		case *mouse.ScrollEvent:
			if em.Scrolling != nil {
				if em.Scrolling == gn.This() {
					rvs.Add(recv, fun, 10000)
				} else {
					return true
				}
			} else {
				if pos.In(gn.WinBBox) {
					rvs.AddDepth(recv, fun, top)
					return false
				}
				return true
			}
		default:
			if em.Dragging == gn.This() { // dragger always gets it
				rvs.Add(recv, fun, 10000) // top priority -- can't steal!
				return false
			}
			if !pos.In(gn.WinBBox) {
				return true
			}
		}
	}
	rvs.AddDepth(recv, fun, top)
	return true
}

// SendSig directly calls SendSig from given recv, sender for given event
// across all priorities.
func (em *EventMgr) SendSig(recv, sender ki.Ki, evi oswin.Event) {
	et := evi.Type()
	for pri := HiPri; pri < EventPrisN; pri++ {
		em.EventSigs[et][pri].SendSig(recv, sender, int64(et), evi)
	}
}

///////////////////////////////////////////////////////////////////////////
//  Mouse event processing

// MouseEvents processes mouse drag and move events
func (em *EventMgr) MouseEvents(evi oswin.Event) {
	et := evi.Type()
	if et == oswin.MouseDragEvent {
		em.MouseDragEvents(evi)
	} else if et != oswin.KeyEvent { // allow modifier keypress
		em.ResetMouseDrag()
	}

	if et == oswin.MouseMoveEvent {
		em.MouseMoveEvents(evi)
	} else {
		em.ResetMouseMove()
	}
}

// MouseEventReset resets state for "catch" events (Dragging, Scrolling)
func (em *EventMgr) MouseEventReset(evi oswin.Event) {
	et := evi.Type()
	if em.Dragging != nil && et != oswin.MouseDragEvent {
		em.Dragging.ClearFlag(int(NodeDragging))
		em.Dragging = nil
	}
	if em.Scrolling != nil && et != oswin.MouseScrollEvent {
		em.Scrolling = nil
	}
}

// MouseDragEvents processes MouseDragEvent to Detect start of drag and DND.
// These require timing and delays, e.g., due to minor wiggles when pressing
// the mouse button
func (em *EventMgr) MouseDragEvents(evi oswin.Event) {
	now := time.Now()
	if !em.dragStarted {
		if em.startDrag == nil {
			em.startDrag = evi.(*mouse.DragEvent)
		} else {
			if em.DoInstaDrag(em.startDrag, !em.Master.CurPopupIsTooltip()) {
				em.dragStarted = true
				em.startDrag = nil
			} else {
				delayMs := int(now.Sub(em.startDrag.Time()) / time.Millisecond)
				if delayMs >= DragStartMSec {
					dst := int(math32.Hypot(float32(em.startDrag.Where.X-evi.Pos().X), float32(em.startDrag.Where.Y-evi.Pos().Y)))
					if dst >= DragStartPix {
						em.dragStarted = true
						em.startDrag = nil
					}
				}
			}
		}
	}
	if em.Dragging == nil && !em.dndStarted {
		if em.startDND == nil {
			em.startDND = evi.(*mouse.DragEvent)
		} else {
			delayMs := int(now.Sub(em.startDND.Time()) / time.Millisecond)
			if delayMs >= DNDStartMSec {
				dst := int(math32.Hypot(float32(em.startDND.Where.X-evi.Pos().X), float32(em.startDND.Where.Y-evi.Pos().Y)))
				if dst >= DNDStartPix {
					em.dndStarted = true
					em.DNDStartEvent(em.startDND)
					em.startDND = nil
				}
			}
		}
	} else { // em.dndStarted
		em.TimerMu.Lock()
		if !em.dndHoverStarted {
			em.dndHoverStarted = true
			em.startDNDHover = evi.(*mouse.DragEvent)
			em.curDNDHover = em.startDNDHover
			em.dndHoverTimer = time.AfterFunc(time.Duration(HoverStartMSec)*time.Millisecond, func() {
				em.TimerMu.Lock()
				hoe := em.curDNDHover
				em.dndHoverStarted = false
				em.startDNDHover = nil
				em.curDNDHover = nil
				em.dndHoverTimer = nil
				em.TimerMu.Unlock()
				em.SendDNDHoverEvent(hoe)
			})
		} else {
			dst := int(math32.Hypot(float32(em.startDNDHover.Where.X-evi.Pos().X), float32(em.startDNDHover.Where.Y-evi.Pos().Y)))
			if dst > HoverMaxPix {
				em.dndHoverTimer.Stop()
				em.dndHoverStarted = false
				em.startDNDHover = nil
				em.dndHoverTimer = nil
			} else {
				em.curDNDHover = evi.(*mouse.DragEvent)
			}
		}
		em.TimerMu.Unlock()
	}
}

// ResetMouseDrag resets all the mouse dragging variables after last drag
func (em *EventMgr) ResetMouseDrag() {
	em.dragStarted = false
	em.startDrag = nil
	em.dndStarted = false
	em.startDND = nil

	em.TimerMu.Lock()
	em.dndHoverStarted = false
	em.startDNDHover = nil
	em.curDNDHover = nil
	if em.dndHoverTimer != nil {
		em.dndHoverTimer.Stop()
		em.dndHoverTimer = nil
	}
	em.TimerMu.Unlock()
}

// MouseMoveEvents processes MouseMoveEvent to detect start of hover events.
// These require timing and delays
func (em *EventMgr) MouseMoveEvents(evi oswin.Event) {
	em.TimerMu.Lock()
	if !em.hoverStarted {
		em.hoverStarted = true
		em.startHover = evi.(*mouse.MoveEvent)
		em.curHover = em.startHover
		em.hoverTimer = time.AfterFunc(time.Duration(HoverStartMSec)*time.Millisecond, func() {
			em.TimerMu.Lock()
			hoe := em.curHover
			em.hoverStarted = false
			em.startHover = nil
			em.curHover = nil
			em.hoverTimer = nil
			em.TimerMu.Unlock()
			em.SendHoverEvent(hoe)
		})
	} else {
		dst := int(math32.Hypot(float32(em.startHover.Where.X-evi.Pos().X), float32(em.startHover.Where.Y-evi.Pos().Y)))
		if dst > HoverMaxPix {
			em.hoverTimer.Stop()
			em.hoverStarted = false
			em.startHover = nil
			em.hoverTimer = nil
			em.Master.DeleteTooltip()
		} else {
			em.curHover = evi.(*mouse.MoveEvent)
		}
	}
	em.TimerMu.Unlock()
}

// ResetMouseMove resets all the mouse moving variables after last move
func (em *EventMgr) ResetMouseMove() {
	em.TimerMu.Lock()
	em.hoverStarted = false
	em.startHover = nil
	em.curHover = nil
	if em.hoverTimer != nil {
		em.hoverTimer.Stop()
		em.hoverTimer = nil
	}
	em.TimerMu.Unlock()
}

// GenMouseFocusEvents processes mouse.MoveEvent to generate mouse.FocusEvent
// events -- returns true if any such events were sent.  If popup is true,
// then only items on popup are in scope, otherwise items NOT on popup are in
// scope (if no popup, everything is in scope).
func (em *EventMgr) GenMouseFocusEvents(mev *mouse.MoveEvent, popup bool) bool {
	fe := mouse.FocusEvent{Event: mev.Event}
	pos := mev.Pos()
	ftyp := oswin.MouseFocusEvent
	updated := false
	updt := false
	send := em.Master.EventTopNode()
	for pri := HiPri; pri < EventPrisN; pri++ {
		em.EventSigs[ftyp][pri].EmitFiltered(send, int64(ftyp), &fe, func(k ki.Ki) bool {
			if k.IsDeleted() { // destroyed is filtered upstream
				return false
			}
			_, ni := KiToNode2D(k)
			if ni != nil {
				if !em.Master.IsInScope(ni, popup) {
					return false
				}
				in := pos.In(ni.WinBBox)
				if in {
					if !ni.HasFlag(int(MouseHasEntered)) {
						fe.Action = mouse.Enter
						ni.SetFlag(int(MouseHasEntered))
						if !updated {
							updt = send.UpdateStart()
							updated = true
						}
						return true // send event
					} else {
						return false // already in
					}
				} else { // mouse not in object
					if ni.HasFlag(int(MouseHasEntered)) {
						fe.Action = mouse.Exit
						ni.ClearFlag(int(MouseHasEntered))
						if !updated {
							updt = send.UpdateStart()
							updated = true
						}
						return true // send event
					} else {
						return false // already out
					}
				}
			} else {
				// todo: 3D
				return false
			}
		})
	}
	if updated {
		send.UpdateEnd(updt)
	}
	return updated
}

// DoInstaDrag tests whether the given mouse DragEvent is on a widget marked
// with InstaDrag
func (em *EventMgr) DoInstaDrag(me *mouse.DragEvent, popup bool) bool {
	et := me.Type()
	for pri := HiPri; pri < EventPrisN; pri++ {
		esig := &em.EventSigs[et][pri]
		for recv, _ := range esig.Cons {
			if recv.IsDestroyed() {
				delete(esig.Cons, recv)
				continue
			}
			if recv.IsDeleted() {
				continue
			}
			_, ni := KiToNode2D(recv)
			if ni != nil {
				if !em.Master.IsInScope(ni, popup) {
					continue
				}
				pos := me.Pos()
				if pos.In(ni.WinBBox) {
					if ni.IsInstaDrag() {
						em.Dragging = ni.This()
						ni.SetFlag(int(NodeDragging))
						return true
					}
				}
			}
		}
	}
	return false
}

// SendHoverEvent sends mouse hover event, based on last mouse move event
func (em *EventMgr) SendHoverEvent(e *mouse.MoveEvent) {
	he := mouse.HoverEvent{Event: e.Event}
	he.Processed = false
	he.Action = mouse.Hover
	em.SendEventSignal(&he, true) // popup = true by default
}

// DNDStartEvent handles drag-n-drop start events.
func (em *EventMgr) DNDStartEvent(e *mouse.DragEvent) {
	de := dnd.Event{EventBase: e.EventBase, Where: e.Where, Modifiers: e.Modifiers}
	de.Processed = false
	de.Action = dnd.Start
	de.DefaultMod()                // based on current key modifiers
	em.SendEventSignal(&de, false) // popup = false: ignore any popups
	// now up to receiver to call StartDragNDrop if they want to..
}

// SendDNDHoverEvent sends DND hover event, based on last mouse move event
func (em *EventMgr) SendDNDHoverEvent(e *mouse.DragEvent) {
	he := dnd.FocusEvent{Event: dnd.Event{EventBase: e.EventBase, Where: e.Where, Modifiers: e.Modifiers}}
	he.Processed = false
	he.Action = dnd.Hover
	em.SendEventSignal(&he, false) // popup = false by default
}

// SendDNDMoveEvent sends DND move event
func (em *EventMgr) SendDNDMoveEvent(e *mouse.DragEvent) *dnd.MoveEvent {
	// todo: when e.Where goes negative, transition to OS DND
	// todo: send move / enter / exit events to anyone listening
	de := &dnd.MoveEvent{Event: dnd.Event{EventBase: e.Event.EventBase, Where: e.Event.Where, Modifiers: e.Event.Modifiers}, From: e.From, LastTime: e.LastTime}
	de.Processed = false
	de.DefaultMod() // based on current key modifiers
	de.Action = dnd.Move
	em.SendEventSignal(de, false) // popup = false: ignore any popups
	em.GenDNDFocusEvents(de, false)
	return de
}

// SendDNDDropEvent sends DND drop event
func (em *EventMgr) SendDNDDropEvent(e *mouse.Event) {
	de := dnd.Event{EventBase: e.EventBase, Where: e.Where, Modifiers: e.Modifiers}
	de.Processed = false
	de.DefaultMod()
	de.Action = dnd.DropOnTarget
	de.Data = em.DNDData
	de.Source = em.DNDSource
	em.DNDSource.ClearFlag(int(NodeDragging))
	em.Dragging = nil
	em.SendEventSignal(&de, false) // popup = false: ignore any popups
	em.DNDFinalEvent = &de
	e.SetProcessed()
}

// ClearDND clears DND state
func (em *EventMgr) ClearDND() {
	em.DNDSource = nil
	em.DNDData = nil
	em.Dragging = nil
}

// GenDNDFocusEvents processes mouse.MoveEvent to generate dnd.FocusEvent
// events -- returns true if any such events were sent.  If popup is true,
// then only items on popup are in scope, otherwise items NOT on popup are in
// scope (if no popup, everything is in scope).  Extra work is done to ensure
// that Exit from prior widget is always sent before Enter to next one.
func (em *EventMgr) GenDNDFocusEvents(mev *dnd.MoveEvent, popup bool) bool {
	fe := dnd.FocusEvent{Event: mev.Event}
	pos := mev.Pos()
	ftyp := oswin.DNDFocusEvent

	// first pass is just to get all the ins and outs
	var ins, outs WinEventRecvList

	send := em.Master.EventTopNode()
	for pri := HiPri; pri < EventPrisN; pri++ {
		esig := &em.EventSigs[ftyp][pri]
		for recv, fun := range esig.Cons {
			if recv.IsDeleted() { // destroyed is filtered upstream
				continue
			}
			_, ni := KiToNode2D(recv)
			if ni != nil {
				if !em.Master.IsInScope(ni, popup) {
					continue
				}
				in := pos.In(ni.WinBBox)
				if in {
					if !ni.HasFlag(int(DNDHasEntered)) {
						ni.SetFlag(int(DNDHasEntered))
						ins.Add(recv, fun, 0)
					}
				} else { // mouse not in object
					if ni.HasFlag(int(DNDHasEntered)) {
						ni.ClearFlag(int(DNDHasEntered))
						outs.Add(recv, fun, 0)
					}
				}
			} else {
				// todo: 3D
			}
		}
	}
	if len(outs)+len(ins) > 0 {
		updt := send.UpdateStart()
		// now send all the exits before the enters..
		fe.Action = dnd.Exit
		for i := range outs {
			outs[i].Call(send, int64(ftyp), &fe)
		}
		fe.Action = dnd.Enter
		for i := range ins {
			ins[i].Call(send, int64(ftyp), &fe)
		}
		send.UpdateEnd(updt)
		return true
	}
	return false
}

///////////////////////////////////////////////////////////////////
//   Key events

// SendKeyChordEvent sends a KeyChord event with given values.  If popup is
// true, then only items on popup are in scope, otherwise items NOT on popup
// are in scope (if no popup, everything is in scope).
func (em *EventMgr) SendKeyChordEvent(popup bool, r rune, mods ...key.Modifiers) {
	ke := key.ChordEvent{}
	ke.SetTime()
	ke.SetModifiers(mods...)
	ke.Rune = r
	ke.Action = key.Press
	em.SendEventSignal(&ke, popup)
}

// SendKeyFunEvent sends a KeyChord event with params from the given KeyFun.
// If popup is true, then only items on popup are in scope, otherwise items
// NOT on popup are in scope (if no popup, everything is in scope).
func (em *EventMgr) SendKeyFunEvent(kf KeyFuns, popup bool) {
	chord := ActiveKeyMap.ChordForFun(kf)
	if chord == "" {
		return
	}
	r, mods, err := chord.Decode()
	if err != nil {
		return
	}
	ke := key.ChordEvent{}
	ke.SetTime()
	ke.Modifiers = mods
	ke.Rune = r
	ke.Action = key.Press
	em.SendEventSignal(&ke, popup)
}

// CurFocus gets the current focus node under mutex protection
func (em *EventMgr) CurFocus() ki.Ki {
	em.FocusMu.RLock()
	defer em.FocusMu.RUnlock()
	return em.Focus
}

// setFocusPtr JUST sets the focus pointer under mutex protection --
// use SetFocus for end-user setting of focus
func (em *EventMgr) setFocusPtr(k ki.Ki) {
	em.FocusMu.Lock()
	em.Focus = k
	em.FocusMu.Unlock()
}

// SetFocus sets focus to given item -- returns true if focus changed.
func (em *EventMgr) SetFocus(k ki.Ki) bool {
	cfoc := em.CurFocus()
	if cfoc == k {
		if k != nil {
			_, ni := KiToNode2D(k)
			if ni != nil && ni.This() != nil {
				ni.SetFocusState(true) // ensure focus flag always set
			}
		}
		return false
	}

	top := em.Master.EventTopNode()
	updt := top.UpdateStart()
	defer top.UpdateEnd(updt)

	if cfoc != nil {
		nii, ni := KiToNode2D(cfoc)
		if ni != nil && ni.This() != nil {
			ni.SetFocusState(false)
			// fmt.Printf("clear foc: %v\n", ni.PathUnique())
			nii.FocusChanged2D(FocusLost)
		}
	}
	em.setFocusPtr(k)
	if k == nil {
		return true
	}
	nii, ni := KiToNode2D(k)
	if ni == nil || ni.This() == nil { // only 2d for now
		em.setFocusPtr(nil)
		return false
	}
	ni.SetFocusState(true)
	em.Master.SetFocusActiveState(true)
	// fmt.Printf("set foc: %v\n", ni.PathUnique())
	em.ClearNonFocus(k) // shouldn't need this but actually sometimes do
	nii.FocusChanged2D(FocusGot)
	return true
}

// 	FocusNext sets the focus on the next item that can accept focus after the
// given item (can be nil) -- returns true if a focus item found.
func (em *EventMgr) FocusNext(foc ki.Ki) bool {
	gotFocus := false
	focusNext := false // get the next guy
	if foc == nil {
		focusNext = true
	}

	focRoot := em.Master.FocusTopNode()

	for i := 0; i < 2; i++ {
		focRoot.FuncDownMeFirst(0, focRoot, func(k ki.Ki, level int, d interface{}) bool {
			if gotFocus {
				return false
			}
			_, ni := KiToNode2D(k)
			if ni == nil || ni.This() == nil {
				return true
			}
			if foc == k { // current focus can be a non-can-focus item
				focusNext = true
				return true
			}
			if !focusNext {
				return true
			}
			if !ni.CanFocus() {
				return true
			}
			em.SetFocus(k)
			gotFocus = true
			return false // done
		})
		if gotFocus {
			return true
		}
		focusNext = true // this time around, just get the first one
	}
	return gotFocus
}

// FocusOnOrNext sets the focus on the given item, or the next one that can
// accept focus -- returns true if a new focus item found.
func (em *EventMgr) FocusOnOrNext(foc ki.Ki) bool {
	cfoc := em.CurFocus()
	if cfoc == foc {
		return true
	}
	_, ni := KiToNode2D(foc)
	if ni == nil || ni.This() == nil {
		return false
	}
	if ni.CanFocus() {
		em.SetFocus(foc)
		return true
	}
	return em.FocusNext(foc)
}

// FocusOnOrPrev sets the focus on the given item, or the previous one that can
// accept focus -- returns true if a new focus item found.
func (em *EventMgr) FocusOnOrPrev(foc ki.Ki) bool {
	cfoc := em.CurFocus()
	if cfoc == foc {
		return true
	}
	_, ni := KiToNode2D(foc)
	if ni == nil || ni.This() == nil {
		return false
	}
	if ni.CanFocus() {
		em.SetFocus(foc)
		return true
	}
	return em.FocusPrev(foc)
}

// FocusPrev sets the focus on the previous item before the given item (can be nil)
func (em *EventMgr) FocusPrev(foc ki.Ki) bool {
	if foc == nil { // must have a current item here
		em.FocusLast()
		return false
	}

	gotFocus := false
	var prevItem ki.Ki

	focRoot := em.Master.FocusTopNode()

	focRoot.FuncDownMeFirst(0, focRoot, func(k ki.Ki, level int, d interface{}) bool {
		if gotFocus {
			return false
		}
		// todo: see about 3D guys
		_, ni := KiToNode2D(k)
		if ni == nil || ni.This() == nil {
			return true
		}
		if foc == k {
			gotFocus = true
			return false
		}
		if !ni.CanFocus() {
			return true
		}
		prevItem = k
		return true
	})
	if gotFocus && prevItem != nil {
		em.SetFocus(prevItem)
		return true
	} else {
		return em.FocusLast()
	}
}

// FocusLast sets the focus on the last item in the tree -- returns true if a
// focusable item was found
func (em *EventMgr) FocusLast() bool {
	var lastItem ki.Ki

	focRoot := em.Master.FocusTopNode()

	focRoot.FuncDownMeFirst(0, focRoot, func(k ki.Ki, level int, d interface{}) bool {
		// todo: see about 3D guys
		_, ni := KiToNode2D(k)
		if ni == nil || ni.This() == nil {
			return true
		}
		if !ni.CanFocus() {
			return true
		}
		lastItem = k
		return true
	})
	em.SetFocus(lastItem)
	if lastItem == nil {
		return false
	}
	return true
}

// ClearNonFocus clears the focus of any non-w.Focus item.
func (em *EventMgr) ClearNonFocus(foc ki.Ki) {
	focRoot := em.Master.FocusTopNode()

	top := em.Master.EventTopNode()
	updated := false
	updt := false

	focRoot.FuncDownMeFirst(0, focRoot, func(k ki.Ki, level int, d interface{}) bool {
		if k == focRoot { // skip top-level
			return true
		}
		// todo: see about 3D guys
		nii, ni := KiToNode2D(k)
		if ni == nil || ni.This() == nil {
			return true
		}
		if foc == k {
			return true
		}
		if ni.HasFocus() {
			// fmt.Printf("ClearNonFocus: %v\n", ni.PathUnique())
			if !updated {
				updated = true
				updt = top.UpdateStart()
			}
			ni.ClearFlag(int(HasFocus))
			nii.FocusChanged2D(FocusLost)
		}
		return true
	})
	if updated {
		top.UpdateEnd(updt)
	}
}

// PushFocus pushes current focus onto stack and sets new focus.
func (em *EventMgr) PushFocus(p ki.Ki) {
	em.FocusMu.Lock()
	if em.FocusStack == nil {
		em.FocusStack = make([]ki.Ki, 0, 50)
	}
	em.FocusStack = append(em.FocusStack, em.Focus)
	em.Focus = nil // don't un-focus on prior item when pushing
	em.FocusMu.Unlock()
	em.FocusOnOrNext(p)
}

// PopFocus pops off the focus stack and sets prev to current focus.
func (em *EventMgr) PopFocus() {
	em.FocusMu.Lock()
	if em.FocusStack == nil || len(em.FocusStack) == 0 {
		em.Focus = nil
		return
	}
	sz := len(em.FocusStack)
	em.Focus = nil
	nxtf := em.FocusStack[sz-1]
	_, ni := KiToNode2D(nxtf)
	if ni != nil && ni.This() != nil {
		em.FocusMu.Unlock()
		em.SetFocus(nxtf)
		em.FocusMu.Lock()
	}
	em.FocusStack = em.FocusStack[:sz-1]
	em.FocusMu.Unlock()
}

///////////////////////////////////////////////////////////////////
//   Master interface

// EventMaster provides additional control methods for the
// event manager, for things beyond its immediate scope
type EventMaster interface {
	// EventTopNode returns the top-level node for this event scope.
	// This is also the node that serves as the event sender.
	// By default it is the Window
	EventTopNode() ki.Ki

	// FocusTopNode returns the top-level node for key event focusing.
	FocusTopNode() ki.Ki

	// IsInScope returns whether given node is in scope for receiving events
	IsInScope(node *Node2DBase, popup bool) bool

	// CurPopupIsTooltip returns true if current popup is a tooltip
	CurPopupIsTooltip() bool

	// DeleteTooltip deletes any tooltip popup (called when hover ends)
	DeleteTooltip()

	// IsFocusActive returns true if focus is active in this master
	IsFocusActive() bool

	// SetFocusActiveState sets focus active state
	SetFocusActiveState(active bool)
}
