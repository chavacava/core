// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"image"
	"slices"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/enums"
	"cogentcore.org/core/events"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/system"
	"cogentcore.org/core/tree"
)

// Scene contains a [Widget] tree, rooted in an embedded [Frame] layout,
// which renders into its [Scene.Pixels] image. The [Scene] is set in a
// [Stage], which the [Scene] has a pointer to.
//
// Each [Scene] contains state specific to its particular usage
// within a given [Stage] and overall rendering context, representing the unit
// of rendering in the Cogent Core framework.
type Scene struct { //core:no-new
	Frame

	// Body provides the main contents of scenes that use control Bars
	// to allow the main window contents to be specified separately
	// from that dynamic control content.  When constructing scenes using
	// a [Body], you can operate directly on the [Body], which has wrappers
	// for most major Scene functions.
	Body *Body `json:"-" xml:"-" set:"-"`

	// WidgetInit is a function called on every newly created [Widget].
	// This can be used to set global configuration and styling for all
	// widgets in conjunction with [App.SceneInit].
	WidgetInit func(w Widget) `json:"-" xml:"-" edit:"-"`

	// Bars contains functions for constructing the control bars for this Scene,
	// attached to different sides of a Scene (e.g., TopAppBar at Top,
	// NavBar at Bottom, etc).  Functions are called in forward order
	// so first added are called first.
	Bars styles.Sides[BarFuncs] `json:"-" xml:"-"`

	// AppBars contains functions for making the plan for the top app bar.
	// If there are no such functions, no top app bar will be made.
	AppBars []func(p *tree.Plan) `json:"-" xml:"-"`

	// Data is the optional data value being represented by this scene.
	// Used e.g., for recycling views of a given item instead of creating new one.
	Data any

	// Size and position relative to overall rendering context.
	SceneGeom math32.Geom2DInt `edit:"-" set:"-"`

	// paint context for rendering
	PaintContext paint.Context `copier:"-" json:"-" xml:"-" display:"-" set:"-"`

	// live pixels that we render into
	Pixels *image.RGBA `copier:"-" json:"-" xml:"-" display:"-" set:"-"`

	// event manager for this scene
	Events Events `copier:"-" json:"-" xml:"-" set:"-"`

	// current stage in which this Scene is set
	Stage *Stage `copier:"-" json:"-" xml:"-" set:"-"`

	// renderBBoxes indicates to render colored bounding boxes for all of the widgets
	// in the scene. This is enabled by the [Inspector] in select element mode.
	renderBBoxes bool

	// renderBBoxHue is current hue for rendering bounding box in [Scene.RenderBBoxes] mode.
	renderBBoxHue float32

	// selectedWidget is the currently selected/hovered widget through the [Inspector] selection mode
	// that should be highlighted with a background color.
	selectedWidget Widget

	// selectedWidgetChan is the channel on which the selected widget through the inspect editor
	// selection mode is transmitted to the inspect editor after the user is done selecting.
	selectedWidgetChan chan Widget `json:"-" xml:"-"`

	// lastRender captures key params from last render.
	// If different then a new ApplyStyleScene is needed.
	lastRender renderParams

	// showIter counts up at start of showing a Scene
	// to trigger Show event and other steps at start of first show
	showIter int

	// directRenders are widgets that render directly to the [RenderWindow]
	// instead of rendering into the Scene Pixels image.
	directRenders []Widget

	// flags are atomic bit flags for [Scene] state.
	flags sceneFlags
}

// sceneFlags are atomic bit flags for [Scene] state.
// They must be atomic to prevent race conditions.
type sceneFlags int64 //enums:bitflag -trim-prefix scene

const (
	// sceneHasShown is whether this scene has been shown.
	// This is used to ensure that [events.Show] is only sent once.
	sceneHasShown sceneFlags = iota

	// sceneUpdating means the Scene is in the process of sceneUpdating.
	// It is set for any kind of tree-level update.
	// Skip any further update passes until it goes off.
	sceneUpdating

	// sceneNeedsRender is whether anything in the Scene needs to be re-rendered
	// (but not necessarily the whole scene itself).
	sceneNeedsRender

	// sceneNeedsLayout is whether the Scene needs a new layout pass.
	sceneNeedsLayout

	// sceneImageUpdated indicates that the Scene's image has been updated
	// e.g., due to a render or a resize. This is reset by the
	// global [RenderWindow] rendering pass, so it knows whether it needs to
	// copy the image up to the GPU or not.
	sceneImageUpdated

	// scenePrefSizing means that this scene is currently doing a
	// PrefSize computation to compute the size of the scene
	// (for sizing window for example); affects layout size computation
	// only for Over
	scenePrefSizing
)

// hasFlag returns whether the given flag is set.
func (sc *Scene) hasFlag(f sceneFlags) bool {
	return sc.flags.HasFlag(f)
}

// setFlag sets the given flags to the given value.
func (sc *Scene) setFlag(on bool, f ...enums.BitFlag) {
	sc.flags.SetFlag(on, f...)
}

// newBodyScene creates a new Scene for use with an associated Body that
// contains the main content of the Scene (e.g., a Window, Dialog, etc).
// It will be constructed from the Bars-configured control bars on each
// side, with the given Body as the central content.
func newBodyScene(body *Body) *Scene {
	sc := NewScene(body.Name + " scene")
	sc.Body = body
	// need to set parent immediately so that SceneInit works,
	// but can not add it yet because it may go elsewhere due
	// to app bars
	tree.SetParent(body, sc)
	return sc
}

// NewScene creates a new [Scene] object without a [Body], e.g., for use
// in a Menu, Tooltip or other such simple popups or non-control-bar Scenes.
func NewScene(name ...string) *Scene {
	sc := tree.New[Scene]()
	if len(name) > 0 {
		sc.SetName(name[0])
	}
	sc.Events.scene = sc
	return sc
}

func (sc *Scene) Init() {
	sc.Scene = sc
	sc.Frame.Init()
	sc.AddContextMenu(sc.standardContextMenu)
	sc.Styler(func(s *styles.Style) {
		s.Cursor = cursors.Arrow
		s.Background = colors.Scheme.Background
		s.Color = colors.Scheme.OnBackground
		// we never want borders on scenes
		s.MaxBorder = styles.Border{}
		s.Direction = styles.Column
		s.Overflow.Set(styles.OverflowAuto) // screen is always scroller of last resort

		// insets and minimum window padding
		if sc.Stage == nil {
			return
		}
		if sc.Stage.Type.isPopup() || (sc.Stage.Type == DialogStage && !sc.Stage.FullWindow) {
			return
		}

		s.Padding.Set(units.Dp(8))
	})
	sc.OnShow(func(e events.Event) {
		currentRenderWindow.SetStageTitle(sc.Stage.Title)
	})
	sc.OnClose(func(e events.Event) {
		sm := sc.Stage.Mains
		if sm == nil {
			return
		}
		sm.mu.RLock()
		defer sm.mu.RUnlock()

		if sm.stack.Len() < 2 {
			return
		}
		// the stage that will be visible next
		st := sm.stack.ValueByIndex(sm.stack.Len() - 2)
		currentRenderWindow.SetStageTitle(st.Title)
	})
	if TheApp.SceneInit != nil {
		TheApp.SceneInit(sc)
	}
}

// renderContext returns the current render context.
// This will be nil prior to actual rendering.
func (sc *Scene) renderContext() *renderContext {
	if sc.Stage == nil {
		return nil
	}
	sm := sc.Stage.Mains
	if sm == nil {
		return nil
	}
	return sm.renderContext
}

// RenderWindow returns the current render window for this scene.
// In general it is best to go through [renderContext] instead of the window.
// This will be nil prior to actual rendering.
func (sc *Scene) RenderWindow() *renderWindow {
	if sc.Stage == nil {
		return nil
	}
	sm := sc.Stage.Mains
	if sm == nil {
		return nil
	}
	return sm.renderWindow
}

// fitInWindow fits Scene geometry (pos, size) into given window geom.
// Calls resize for the new size.
func (sc *Scene) fitInWindow(winGeom math32.Geom2DInt) {
	geom := sc.SceneGeom
	// full offscreen windows ignore any window geometry constraints
	// because they must be unbounded by any previous window sizes
	if TheApp.Platform() != system.Offscreen || !sc.Stage.FullWindow {
		geom = geom.FitInWindow(winGeom)
	}
	sc.resize(geom)
	sc.SceneGeom.Pos = geom.Pos
	// fmt.Println("win", winGeom, "geom", geom)
}

// resize resizes the scene, creating a new image; updates Geom
func (sc *Scene) resize(geom math32.Geom2DInt) {
	if geom.Size.X <= 0 || geom.Size.Y <= 0 {
		return
	}
	if sc.PaintContext.State == nil {
		sc.PaintContext.State = &paint.State{}
	}
	if sc.PaintContext.Paint == nil {
		sc.PaintContext.Paint = &styles.Paint{}
	}
	sc.SceneGeom.Pos = geom.Pos
	if sc.Pixels == nil || sc.Pixels.Bounds().Size() != geom.Size {
		sc.Pixels = image.NewRGBA(image.Rectangle{Max: geom.Size})
	}
	sc.PaintContext.Init(geom.Size.X, geom.Size.Y, sc.Pixels)
	sc.SceneGeom.Size = geom.Size // make sure

	sc.updateScene()
	sc.applyStyleScene()
	// restart the multi-render updating after resize, to get windows to update correctly while
	// resizing on Windows (OS) and Linux (see https://github.com/cogentcore/core/issues/584), to get
	// windows on Windows (OS) to update after a window snap (see https://github.com/cogentcore/core/issues/497),
	// and to get FillInsets to overwrite mysterious black bars that otherwise are rendered on both iOS
	// and Android in different contexts.
	// TODO(kai): is there a more efficient way to do this, and do we need to do this on all platforms?
	sc.showIter = 0
	sc.NeedsLayout()
}

// Close closes the [Stage] associated with this [Scene].
// This only works for main stages (windows and dialogs).
// It returns whether the [Stage] was successfully closed.
func (sc *Scene) Close() bool {
	if sc == nil {
		return true
	}
	e := &events.Base{Typ: events.Close}
	e.Init()
	sc.WidgetWalkDown(func(cw Widget, cwb *WidgetBase) bool {
		cw.AsWidget().HandleEvent(e)
		return tree.Continue
	})
	// if they set the event as handled, we do not close the scene
	if e.IsHandled() {
		return false
	}
	mm := sc.Stage.Mains
	if mm == nil {
		return false // todo: needed, but not sure why
	}
	mm.deleteStage(sc.Stage)
	if sc.Stage.NewWindow && !TheApp.Platform().IsMobile() && !mm.renderWindow.closing && !mm.renderWindow.stopEventLoop && !TheApp.IsQuitting() {
		mm.renderWindow.closeReq()
	}
	return true
}

func (sc *Scene) ApplyScenePos() {
	sc.Frame.ApplyScenePos()
	if sc.Parts == nil {
		return
	}

	mvi := sc.Parts.ChildByName("move", 1)
	if mvi == nil {
		return
	}
	mv := mvi.(Widget).AsWidget()

	sc.Parts.Geom.Pos.Total.Y = math32.Ceil(0.5 * mv.Geom.Size.Actual.Total.Y)
	sc.Parts.Geom.Size.Actual = sc.Geom.Size.Actual
	sc.Parts.setContentPosFromPos()
	sc.Parts.setBBoxesFromAllocs()
	sc.Parts.applyScenePosChildren()

	psz := sc.Parts.Geom.Size.Actual.Content

	mv.Geom.RelPos.X = 0.5*psz.X - 0.5*mv.Geom.Size.Actual.Total.X
	mv.Geom.RelPos.Y = 0
	mv.setPosFromParent()
	mv.setBBoxesFromAllocs()

	rszi := sc.Parts.ChildByName("resize", 1)
	if rszi == nil {
		return
	}
	rsz := rszi.(Widget).AsWidget()
	rsz.Geom.RelPos.X = psz.X // - 0.5*rsz.Geom.Size.Actual.Total.X
	rsz.Geom.RelPos.Y = psz.Y // - 0.5*rsz.Geom.Size.Actual.Total.Y
	rsz.setPosFromParent()
	rsz.setBBoxesFromAllocs()
}

func (sc *Scene) AddDirectRender(w Widget) {
	sc.directRenders = append(sc.directRenders, w)
}

func (sc *Scene) DeleteDirectRender(w Widget) {
	idx := slices.Index(sc.directRenders, w)
	if idx >= 0 {
		sc.directRenders = slices.Delete(sc.directRenders, idx, idx+1)
	}
}
