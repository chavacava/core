// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"bytes"
	"fmt"
	"image"
	"io"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/goosi"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keyfun"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/states"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/svg"
	"cogentcore.org/core/units"
)

// TheApp is the current [App]; only one is ever in effect.
var TheApp = &App{App: goosi.TheApp}

// App represents a Cogent Core app. It extends [goosi.App] to provide both system-level
// and high-level data and functions to do with the currently running application. The
// single instance of it is [TheApp], which embeds [goosi.TheApp].
type App struct { //gti:add -setters
	goosi.App `set:"-"`

	// Icon specifies the app icon, which is passed to [goosi.Window.SetIcon].
	// It should typically be set using [App.SetIconSVG].
	Icon []image.Image

	// AppBarConfig is the function that configures the AppBar,
	// typically put in the [Scene.Bars.Top] (i.e., a TopAppBar).
	// It is set to StdAppBarConfig by default, which makes the
	// standard AppBar behavior. If this is nil, then no AppBar
	// will be created by default.
	AppBarConfig func(parent Widget)

	// SceneConfig is the function called on every newly created [gi.Scene]
	// to configure it, if it is non-nil. This can be used to set global
	// configuration and styling for all widgets using the OnWidgetAdded
	// method.
	SceneConfig func(sc *Scene)
}

// SetIconSVG sets the icon of the app to the given SVG icon.
// It automatically logs any errors.
func (a *App) SetIconSVG(r io.Reader) *App {
	a.Icon = make([]image.Image, 3)

	sv := svg.NewSVG(16, 16)
	sv.Color = colors.C(colors.FromRGB(66, 133, 244)) // Google Blue (#4285f4)
	err := sv.ReadXML(r)
	if grr.Log(err) != nil {
		return a
	}

	sv.Render()
	a.Icon[0] = sv.Pixels

	sv.Resize(image.Pt(32, 32))
	sv.Render()
	a.Icon[1] = sv.Pixels

	sv.Resize(image.Pt(48, 48))
	sv.Render()
	a.Icon[2] = sv.Pixels
	return a
}

// SetIconBytes sets the icon of the app to the given SVG icon bytes.
// It automatically logs any errors.
func (a *App) SetIconBytes(b []byte) *App {
	return a.SetIconSVG(bytes.NewReader(b))
}

// Quit closes all windows and exits the program.
func Quit() {
	if !goosi.TheApp.IsQuitting() {
		goosi.TheApp.Quit()
	}
}

//////////////////////////////////////////////////////////////////////////////
//		AppBar

// StandardAppBarConfig is the standard impl for a [App.AppBarConfig].
// It adds a Back navigation buttons and the AppChooser,
// followed by the [Widget.ConfigToolbar] for the current FullWindow
// Scene being viewed, along with [StandardOverflowMenu] items.
// and calls AddDefaultOverflowMenu to provide default menu items,
// which will appear below any other OverflowMenu items added.
func StandardAppBarConfig(parent Widget) {
	tb := RecycleToolbar(parent)
	StandardAppBarStart(tb)
	StandardOverflowMenu(tb)
	CurrentWindowAppBar(tb)
}

// StandardAppBarStart adds standard items to start of an AppBar:
// [StandardAppBarBack] and [StandardAppBarChooser]
func StandardAppBarStart(tb *Toolbar) {
	StandardAppBarBack(tb)
	StandardAppBarChooser(tb)
}

// StandardAppBarBack adds a back button
func StandardAppBarBack(tb *Toolbar) *Button {
	bt := NewButton(tb, "back").SetIcon(icons.ArrowBack).SetTooltip("Back").SetKey(keyfun.HistPrev)
	// TODO(kai/abc): app bar back button disabling
	// bt.StyleFirst(func(s *styles.Style) {
	// 	if tb.Scene.Stage.MainMgr == nil {
	// 		return
	// 	}
	// 	s.SetState(tb.Scene.Stage.MainMgr.Stack.Len() <= 1 && len(AllRenderWins) <= 1, states.Disabled)
	// })
	bt.OnClick(func(e events.Event) {
		if slen := tb.Scene.Stage.MainMgr.Stack.Len(); slen > 1 {
			if tb.Scene.Stage.CloseOnBack {
				tb.Scene.Close()
			} else {
				tb.Scene.Stage.MainMgr.Stack.ValueByIndex(slen - 2).Raise()
			}
			return
		}
		if wlen := len(AllRenderWins); wlen > 1 {
			if tb.Scene.Stage.CloseOnBack {
				CurRenderWin.CloseReq()
			}
			AllRenderWins[wlen-2].Raise()
		}
	})
	return bt
}

// StandardAppBarChooser adds an AppChooser
func StandardAppBarChooser(tb *Toolbar) *Chooser {
	return ConfigAppChooser(NewChooser(tb, "app-chooser"), tb)
}

// todo: use CurrentMainScene instead?

// CurrentWindowAppBar calls ConfigToolbar functions registered on
// the Scene to which the given toolbar belongs.
func CurrentWindowAppBar(tb *Toolbar) {
	tb.Scene.AppBars.Call(tb)
}

// StandardOverflowMenu adds the standard overflow menu function.
func StandardOverflowMenu(tb *Toolbar) {
	tb.OverflowMenus = append(tb.OverflowMenus, tb.StandardOverflowMenu)
}

// note: must be a method on toolbar to get scene

// StandardOverflowMenu adds standard overflow menu items.
func (tb *Toolbar) StandardOverflowMenu(m *Scene) { //gti:add
	if SettingsWindow != nil {
		NewButton(m).SetText("Settings").SetIcon(icons.Settings).SetKey(keyfun.Settings).
			OnClick(func(e events.Event) {
				SettingsWindow()
			})
	}
	if InspectorWindow != nil {
		NewButton(m).SetText("Inspect").SetIcon(icons.Edit).SetKey(keyfun.Inspector).
			OnClick(func(e events.Event) {
				InspectorWindow(tb.Scene)
			})
	}
	NewButton(m).SetText("Edit").SetMenu(func(m *Scene) {
		// todo: these need to actually do something -- currently just show keyboard shortcut
		NewButton(m).SetText("Copy").SetIcon(icons.Copy).SetKey(keyfun.Copy)
		NewButton(m).SetText("Cut").SetIcon(icons.Cut).SetKey(keyfun.Cut)
		NewButton(m).SetText("Paste").SetIcon(icons.Paste).SetKey(keyfun.Paste)
	})

	// no window menu on single-window platforms
	if TheApp.Platform().IsMobile() {
		return
	}
	NewButton(m).SetText("Window").SetMenu(func(m *Scene) {
		NewButton(m).SetText("Focus next").SetIcon(icons.CenterFocusStrong).
			SetKey(keyfun.WinFocusNext).OnClick(func(e events.Event) {
			AllRenderWins.FocusNext()
		})
		NewButton(m).SetText("Minimize").SetIcon(icons.Minimize).
			OnClick(func(e events.Event) {
				win := tb.Scene.RenderWin()
				if win != nil {
					win.Minimize()
				}
			})
		NewSeparator(m)
		NewButton(m).SetText("Close window").SetIcon(icons.Close).SetKey(keyfun.WinClose).
			OnClick(func(e events.Event) {
				win := tb.Scene.RenderWin()
				if win != nil {
					win.CloseReq()
				}
			})
		NewButton(m, "quit-app").SetText("Quit").SetIcon(icons.Close).SetShortcut("Command+Q").
			OnClick(func(e events.Event) {
				go TheApp.QuitReq()
			})
		NewSeparator(m)
		for _, w := range MainRenderWins {
			if w != nil {
				NewButton(m).SetText(w.Title).OnClick(func(e events.Event) {
					w.Raise()
				})
			}
		}
		if len(DialogRenderWins) > 0 {
			NewSeparator(m)
			for _, w := range DialogRenderWins {
				if w != nil {
					NewButton(m).SetText(w.Title).OnClick(func(e events.Event) {
						w.Raise()
					})
				}
			}
		}
	})
}

//////////////////////////////////////////////////////////////////////////////
//		AppChooser

// ConfigAppChooser configures the given [Chooser] to give access
// to all app resources, such as open scenes and buttons in the
// given toolbar. This chooser is typically placed at the start
// of the AppBar. You can extend the resources available for access
// in the app chooser using [Chooser.AddItemsFunc] and [ChooserItem.Func].
func ConfigAppChooser(ch *Chooser, tb *Toolbar) *Chooser {
	ch.SetEditable(true).SetType(ChooserOutlined).SetIcon(icons.Search)
	if TheApp.SystemPlatform().IsMobile() {
		ch.SetPlaceholder("Search")
	} else {
		ch.SetPlaceholder(fmt.Sprintf("Search (%s)", keyfun.Menu.Label()))
	}

	ch.OnWidgetAdded(func(w Widget) {
		if w.PathFrom(ch) == "parts/text" {
			tf := w.(*TextField)
			w.Style(func(s *styles.Style) {
				s.Background = colors.C(colors.Scheme.SurfaceContainerHighest)
				if !s.Is(states.Focused) && tf.Error == nil {
					s.Border = styles.Border{}
				}
				s.Border.Radius = styles.BorderRadiusFull
				s.Min.X.SetCustom(func(uc *units.Context) float32 {
					return min(mat32.Clamp(uc.Ch(40), uc.Vw(20), uc.Vw(80)), uc.Ch(80))
				})
			})
		}
	})

	ch.AddItemsFunc(func() {
		for _, rw := range AllRenderWins {
			for _, kv := range rw.MainStageMgr.Stack.Order {
				st := kv.Value
				// we do not include ourself
				if st == tb.Scene.Stage {
					continue
				}
				ch.Items = append(ch.Items, ChooserItem{
					Label:   st.Title,
					Icon:    icons.Toolbar,
					Tooltip: "Show " + st.Title,
					Func:    st.Raise,
				})
			}
		}
	})
	ch.AddItemsFunc(func() {
		AddButtonItems(&ch.Items, tb, "")
	})
	ch.OnFinal(events.Change, func(e events.Event) {
		// we must never have a chooser label so that it
		// always displays the search placeholder
		ch.CurrentIndex = -1
		ch.CurrentItem = ChooserItem{}
		ch.ShowCurrentItem()
	})
	ch.OnFirst(events.KeyChord, func(e events.Event) {
		kf := keyfun.Of(e.KeyChord())
		if kf == keyfun.Menu {
			e.SetHandled()
			ch.Send(events.Click, e)
		}
	})
	return ch
}

// AddButtonItems adds to the given items all of the buttons under
// the given parent. It navigates through button menus to find other
// buttons using a recursive approach that updates path with context
// about the original button menu. Consumers of this function should
// typically set path to "".
func AddButtonItems(items *[]ChooserItem, parent ki.Node, path string) {
	for _, kid := range *parent.Children() {
		bt := AsButton(kid)
		if bt == nil || bt.IsDisabled() {
			continue
		}
		lbl := bt.Text
		if lbl == "" {
			lbl = bt.Tooltip
		}
		if bt.HasMenu() {
			tmps := NewScene()
			bt.Menu(tmps)
			npath := path
			if npath != "" {
				npath += " > "
			}
			if bt.Name() != "overflow-menu" {
				npath += lbl
			}
			AddButtonItems(items, tmps, npath)
			continue
		}
		if path != "" {
			lbl = path + " > " + lbl
		}
		*items = append(*items, ChooserItem{
			Label:   lbl,
			Icon:    bt.Icon,
			Tooltip: bt.Tooltip,
			Func: func() {
				bt.Send(events.Click)
			},
		})
		// after the quit button, there are the render wins,
		// which we do not want to show here as we are already
		// showing the stages
		if bt.Name() == "quit-app" {
			break
		}
	}
}
