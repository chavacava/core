// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/histyle"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/laser"
)

////////////////////////////////////////////////////////////////////////////////////////
//  HiStyleValue

// HiStyleValue presents an action for displaying a mat32.Y and selecting
// from styles
type HiStyleValue struct {
	ValueBase
}

func (vv *HiStyleValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.ButtonType
	return vv.WidgetTyp
}

func (vv *HiStyleValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	ac := vv.Widget.(*gi.Button)
	txt := laser.ToString(vv.Value.Interface())
	ac.SetText(txt)
}

func (vv *HiStyleValue) ConfigWidget(widg gi.Widget) {
	vv.Widget = widg
	vv.StdConfigWidget(widg)
	ac := vv.Widget.(*gi.Button)
	ac.SetProp("border-radius", units.Dp(4))
	ac.OnClick(func(e events.Event) {
		vv.OpenDialog(ac, nil)
	})
	vv.UpdateWidget()
}

func (vv *HiStyleValue) HasDialog() bool {
	return true
}

func (vv *HiStyleValue) OpenDialog(ctx gi.Widget, fun func(dlg *gi.Dialog)) {
	if vv.IsInactive() {
		return
	}
	cur := laser.ToString(vv.Value.Interface())
	desc, _ := vv.Tag("desc")
	SliceViewSelectDialog(ctx, DlgOpts{Title: "Select a HiStyle Highlighting Style", Prompt: desc}, &histyle.StyleNames, cur, nil, func(dlg *gi.Dialog) {
		if dlg.Accepted {
			si := dlg.Data.(int)
			if si >= 0 {
				hs := histyle.StyleNames[si]
				vv.SetValue(hs)
				vv.UpdateWidget()
			}
		}
		if fun != nil {
			fun(dlg)
		}
	}).Run()
}

//////////////////////////////////////////////////////////////////////////////////////
//  HiStylesView

// HiStylesView opens a view of highlighting styles
func HiStylesView(st *histyle.Styles) {
	if gi.ActivateExistingMainWindow(st) {
		return
	}

	sc := gi.StageScene("hi-styles")
	sc.Title = "Hilighting Styles: use ViewStd to see builtin ones -- can add and customize -- save ones from standard and load into custom to modify standards."
	sc.Lay = gi.LayoutVert
	sc.Data = st

	title := gi.NewLabel(sc, "title").SetText(sc.Title)
	title.AddStyles(func(s *styles.Style) {
		s.SetMinPrefWidth(units.Ch(30)) // need for wrap
		s.SetStretchMaxWidth()
		s.Text.WhiteSpace = styles.WhiteSpaceNormal // wrap
	})

	tv := NewMapView(sc, "tv")
	tv.SetMap(st)
	tv.AddStyles(func(s *styles.Style) {
		tv.SetStretchMax()
	})

	histyle.StylesChanged = false
	tv.OnChange(func(e events.Event) {
		histyle.StylesChanged = true
	})

	// mmen := win.MainMenu
	// MainMenuView(st, win, mmen)

	// todo: close prompt
	/*
		inClosePrompt := false
		win.RenderWin.SetCloseReqFunc(func(w goosi.RenderWin) {
			if !histyle.StylesChanged || st != &histyle.CustomStyles { // only for main avail map..
				win.Close()
				return
			}
			if inClosePrompt {
				return
			}
			inClosePrompt = true
			gi.ChoiceDialog(sc, gi.DlgOpts{Title: "Save Styles Before Closing?",
				Prompt: "Do you want to save any changes to std preferences styles file before closing, or Cancel the close and do a Save to a different file?"},
				[]string{"Save and Close", "Discard and Close", "Cancel"}, func(dlg *gi.Dialog) {
					switch sig {
					case 0:
						st.SavePrefs()
						fmt.Printf("Preferences Saved to %v\n", histyle.PrefsStylesFileName)
						win.Close()
					case 1:
						st.OpenPrefs() // revert
						win.Close()
					case 2:
						inClosePrompt = false
						// default is to do nothing, i.e., cancel
					}
				})
		})

		win.MainMenuUpdated()
	*/

	gi.NewWindow(sc).Run() // todo: should be a dialog instead?
}
