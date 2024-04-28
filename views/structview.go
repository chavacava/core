// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"cogentcore.org/core/base/reflectx"
	"cogentcore.org/core/base/strcase"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// NoSentenceCaseFor indicates to not transform field names in
// [StructView]s into "Sentence case" for types whose full,
// package-path-qualified name contains any of these strings.
// For example, this can be used to disable sentence casing
// for types with scientific abbreviations in field names,
// which are more readable when not sentence cased. However,
// this should not be needed in most circumstances.
var NoSentenceCaseFor []string

// NoSentenceCaseForType returns whether the given fully
// package-path-qualified name contains anything in the
// [NoSentenceCaseFor] list.
func NoSentenceCaseForType(tnm string) bool {
	return slices.ContainsFunc(NoSentenceCaseFor, func(s string) bool {
		return strings.Contains(tnm, s)
	})
}

// StructView represents a struct with rows of field names and editable values.
type StructView struct {
	core.Frame

	// Struct is the pointer to the struct that we are viewing.
	Struct any `set:"-"`

	// StructValue is the Value for the struct itself
	// if this was created within the Value framework.
	// Otherwise, it is nil.
	StructValue Value `set:"-"`

	// Values are [Value] representations of the struct field values.
	Values []Value `set:"-" json:"-" xml:"-"`

	// ViewPath is a record of parent view names that have led up to this view.
	// It is displayed as extra contextual information in view dialogs.
	ViewPath string

	// isShouldShower is whether the struct implements [core.ShouldShower], which results
	// in additional updating being done at certain points.
	isShouldShower bool
}

func (sv *StructView) OnInit() {
	sv.Frame.OnInit()
	sv.SetStyles()
}

func (sv *StructView) SetStyles() {
	sv.Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(0, 0)
	})
	sv.OnWidgetAdded(func(w core.Widget) {
		pfrom := w.PathFrom(sv)
		switch {
		case pfrom == "struct-grid":
			w.Style(func(s *styles.Style) {
				s.Display = styles.Grid
				s.Grow.Set(0, 0)
				if sv.SizeClass() == core.SizeCompact {
					s.Columns = 1
				} else {
					s.Columns = 2
				}
			})
		}
	})
}

// SetStruct sets the source struct that we are viewing -- rebuilds the
// children to represent this struct
func (sv *StructView) SetStruct(st any) *StructView {
	sv.Struct = st
	sv.Update()
	return sv
}

// UpdateFields updates each of the value-view widgets for the fields --
// called by the ViewSig update
func (sv *StructView) UpdateFields() {
	for _, vv := range sv.Values {
		// we do not update focused elements to prevent panics
		if wb := vv.AsWidgetBase(); wb != nil {
			if wb.StateIs(states.Focused) {
				continue
			}
		}
		vv.Update()
	}
	sv.NeedsRender()
}

// UpdateField updates the value-view widget for the named field
func (sv *StructView) UpdateField(field string) {
	for _, vv := range sv.Values {
		if vv.Name() == field {
			vv.Update()
			break
		}
	}
	sv.NeedsRender()
}

// Config configures the view
func (sv *StructView) Config() {
	if ks, ok := sv.Struct.(tree.Node); ok {
		if ks == nil || ks.This() == nil {
			return
		}
	}
	if sv.HasChildren() {
		sv.ConfigStructGrid()
		return
	}
	core.NewFrame(sv, "struct-grid")
	sv.ConfigStructGrid()
	sv.NeedsLayout()
}

// IsConfiged returns true if the widget is fully configured
func (sv *StructView) IsConfiged() bool {
	return len(sv.Kids) != 0
}

// StructGrid returns the grid layout widget, which contains all the fields and values
func (sv *StructView) StructGrid() *core.Frame {
	return sv.ChildByName("struct-grid", 2).(*core.Frame)
}

// ConfigStructGrid configures the StructGrid for the current struct.
// returns true if any fields changed.
func (sv *StructView) ConfigStructGrid() bool {
	if reflectx.AnyIsNil(sv.Struct) {
		return false
	}
	sc := true
	if len(NoSentenceCaseFor) > 0 {
		sc = !NoSentenceCaseForType(types.TypeNameObj(sv.Struct))
	}
	sg := sv.StructGrid()
	// note: widget re-use does not work due to all the closures
	sg.DeleteChildren()
	config := tree.Config{}
	dupeFields := map[string]bool{}
	sv.Values = make([]Value, 0)

	shouldShow := func(field reflect.StructField, stru any) bool {
		ftags := field.Tag
		vwtag := ftags.Get("view")
		if vwtag == "-" {
			return false
		}
		if ss, ok := stru.(core.ShouldShower); ok {
			sv.isShouldShower = true
			if !ss.ShouldShow(field.Name) {
				return false
			}
		}
		return true
	}

	reflectx.WalkValueFlatFieldsIf(sv.Struct,
		func(stru any, typ reflect.Type, field reflect.StructField, fieldVal reflect.Value) bool {
			return shouldShow(field, sv.Struct)
		},
		func(fval any, typ reflect.Type, field reflect.StructField, fieldVal reflect.Value) bool {
			// todo: check tags, skip various etc
			ftags := field.Tag
			vwtag := ftags.Get("view")
			if !shouldShow(field, sv.Struct) {
				return true
			}
			if vwtag == "add-fields" && field.Type.Kind() == reflect.Struct {
				fvalp := fieldVal.Addr().Interface()
				reflectx.WalkValueFlatFieldsIf(fvalp,
					func(stru any, typ reflect.Type, sfield reflect.StructField, fieldVal reflect.Value) bool {
						return shouldShow(sfield, fvalp)
					},
					func(sfval any, styp reflect.Type, sfield reflect.StructField, sfieldVal reflect.Value) bool {
						if !shouldShow(sfield, fvalp) {
							return true
						}
						svv := FieldToValue(fvalp, sfield.Name, sfval)
						if svv == nil { // shouldn't happen
							return true
						}
						svvp := sfieldVal.Addr()
						svv.SetStructValue(svvp, fvalp, &sfield, sv.ViewPath)

						svtyp := svv.WidgetType()
						// todo: other things with view tag..
						fnm := field.Name + " • " + sfield.Name
						if _, exists := dupeFields[fnm]; exists {
							slog.Error("StructView: duplicate field name:", "name:", fnm)
						} else {
							dupeFields[fnm] = true
						}
						if sc {
							svv.SetLabel(strcase.ToSentence(fnm))
						} else {
							svv.SetLabel(fnm)
						}
						labnm := fmt.Sprintf("label-%v", fnm)
						valnm := fmt.Sprintf("value-%v", fnm)
						config.Add(core.TextType, labnm)
						config.Add(svtyp, valnm) // todo: extend to diff types using interface..
						sv.Values = append(sv.Values, svv)
						return true
					})
				return true
			}
			vv := FieldToValue(sv.Struct, field.Name, fval)
			if vv == nil { // shouldn't happen
				return true
			}
			if _, exists := dupeFields[field.Name]; exists {
				slog.Error("StructView: duplicate field name:", "name:", field.Name)
			} else {
				dupeFields[field.Name] = true
			}
			vvp := fieldVal.Addr()
			vv.SetStructValue(vvp, sv.Struct, &field, sv.ViewPath)
			vtyp := vv.WidgetType()
			// todo: other things with view tag..
			labnm := fmt.Sprintf("label-%v", field.Name)
			valnm := fmt.Sprintf("value-%v", field.Name)
			config.Add(core.TextType, labnm)
			config.Add(vtyp, valnm) // todo: extend to diff types using interface..
			sv.Values = append(sv.Values, vv)
			return true
		})
	sg.ConfigChildren(config) // fields could be non-unique with labels..
	for i, vv := range sv.Values {
		lbl := sg.Child(i * 2).(*core.Text)
		lbl.Style(func(s *styles.Style) {
			s.SetTextWrap(false)
		})
		lbl.Tooltip = vv.Doc()
		vv.AsValueData().ViewPath = sv.ViewPath
		w, wb := core.AsWidget(sg.Child((i * 2) + 1))
		hasDef, readOnlyTag := StructViewFieldTags(vv, lbl, w, sv.IsReadOnly())
		if hasDef {
			lbl.Style(func(s *styles.Style) {
				dtag, _ := vv.Tag("default")
				isDef, _ := StructFieldIsDef(dtag, vv.Val().Interface(), reflectx.NonPointerValue(vv.Val()).Kind())
				dcr := "(Double click to reset to default) "
				if !isDef {
					s.Color = colors.C(colors.Scheme.Primary.Base)
					s.Cursor = cursors.Poof
					if !strings.HasPrefix(lbl.Tooltip, dcr) {
						lbl.Tooltip = dcr + lbl.Tooltip
					}
				} else {
					lbl.Tooltip = strings.TrimPrefix(lbl.Tooltip, dcr)
				}
			})
			lbl.OnDoubleClick(func(e events.Event) {
				dtag, _ := vv.Tag("default")
				isDef, _ := StructFieldIsDef(dtag, vv.Val().Interface(), reflectx.NonPointerValue(vv.Val()).Kind())
				if isDef {
					return
				}
				e.SetHandled()
				err := reflectx.SetFromDefaultTag(vv.Val(), dtag)
				if err != nil {
					core.ErrorSnackbar(lbl, err, "Error setting default value")
				} else {
					vv.Update()
					vv.SendChange(e)
				}
			})
		}
		if w.NodeType() != vv.WidgetType() {
			slog.Error("StructView: Widget Type is not the proper type.  This usually means there are duplicate field names (including across embedded types", "field:", lbl.Text, "is:", w.NodeType().Name, "should be:", vv.WidgetType().Name)
			break
		}
		Config(vv, w)
		vv.AsWidgetBase().OnInput(func(e events.Event) {
			if tag, _ := vv.Tag("immediate"); tag == "+" {
				wb.SendChange(e)
				sv.SendChange(e)
			}
			sv.Send(events.Input, e)
		})
		if !sv.IsReadOnly() && !readOnlyTag {
			vv.OnChange(func(e events.Event) {
				sv.UpdateFieldAction()
				// note: updating vv here is redundant -- relevant field will have already updated
				if !reflectx.KindIsBasic(reflectx.NonPointerValue(vv.Val()).Kind()) {
					if updtr, ok := sv.Struct.(core.Updater); ok {
						updtr.Update()
					}
				}
				if hasDef {
					lbl.Update()
				}
				sv.SendChange(e)
			})
		}
	}
	return true
}

func (sv *StructView) UpdateFieldAction() {
	if !sv.IsConfiged() {
		return
	}
	if sv.isShouldShower {
		sv.Update()
	}
}

/////////////////////////////////////////////////////////////////////////
//  Tag parsing

// StructViewFieldTags processes the tags for a field in a struct view, setting
// the properties on the label or widget appropriately
// returns true if there were any "default" default tags -- if so, needs updating
func StructViewFieldTags(vv Value, lbl *core.Text, w core.Widget, isReadOnly bool) (hasDef, readOnlyTag bool) {
	lbl.Text = vv.Label()
	if et, has := vv.Tag("edit"); has && et == "-" {
		readOnlyTag = true
		w.AsWidget().SetReadOnly(true)
	} else {
		if isReadOnly {
			w.AsWidget().SetReadOnly(true)
			vv.SetTag("edit", "-")
		}
	}
	defStr, hasDef := vv.Tag("default")
	if hasDef {
		lbl.Tooltip = "(Default: " + defStr + ") " + vv.Doc()
	} else {
		lbl.Tooltip = vv.Doc()
	}
	return
}

// StructFieldIsDef processses "default" tag for default value(s) of field
// defs = default values as strings as either comma-separated list of valid values
// or low:high value range (only for int or float numeric types)
// valPtr = pointer to value
// returns true if value is default, and string to add to tooltip for default values.
// Uses JSON format for composite types (struct, slice, map), replacing " with '
// so it is easier to use in def tag.
func StructFieldIsDef(defs string, valPtr any, kind reflect.Kind) (bool, string) {
	defStr := "(Default: " + defs + ")"
	if kind >= reflect.Int && kind <= reflect.Complex128 && strings.Contains(defs, ":") {
		dtags := strings.Split(defs, ":")
		lo, _ := strconv.ParseFloat(dtags[0], 64)
		hi, _ := strconv.ParseFloat(dtags[1], 64)
		vf, err := reflectx.ToFloat(valPtr)
		if err != nil {
			slog.Error("views.StructFieldIsDef: error parsing struct field numerical range def tag", "type", reflectx.NonPointerType(reflect.TypeOf(valPtr)), "def", defs, "err", err)
			return true, defStr
		}
		return lo <= vf && vf <= hi, defStr
	}
	v := reflectx.NonPointerValue(reflect.ValueOf(valPtr))
	dtags := strings.Split(defs, ",")
	if strings.ContainsAny(defs, "{[") { // complex type, so don't split on commas
		dtags = []string{defs}
	}
	for _, def := range dtags {
		def = reflectx.FormatDefault(def)
		if def == "" {
			if v.IsZero() {
				return true, defStr
			}
			continue
		}
		dv := reflect.New(v.Type())
		err := reflectx.SetRobust(dv.Interface(), def)
		if err != nil {
			slog.Error("views.StructFieldIsDef: error parsing struct field def tag", "type", v.Type(), "def", def, "err", err)
			return true, defStr
		}
		if reflect.DeepEqual(v.Interface(), dv.Elem().Interface()) {
			return true, defStr
		}
	}
	return false, defStr
}

// StructFieldVals represents field values in a struct, at multiple
// levels of depth potentially (represented by the Path field)
// used for StructNonDefFields for example.
type StructFieldVals struct {

	// path of field.field parent fields to this field
	Path string

	// type information for field
	Field reflect.StructField

	// value of field (as a pointer)
	Val reflect.Value

	// def tag information for default values
	Defs string
}

// StructNonDefFields processses "default" tag for default value(s)
// of fields in given struct and starting path, and returns all
// fields not at their default values.
// See also StructNoDefFieldsStr for a string representation of this information.
// Uses reflectx.FlatFieldsValueFunc to get all embedded fields.
// Uses a recursive strategy -- any fields that are themselves structs are
// expanded, and the field name represented by dots path separators.
func StructNonDefFields(structPtr any, path string) []StructFieldVals {
	var flds []StructFieldVals
	reflectx.WalkValueFlatFields(structPtr, func(fval any, typ reflect.Type, field reflect.StructField, fieldVal reflect.Value) bool {
		vvp := fieldVal.Addr()
		dtag, got := field.Tag.Lookup("default")
		if field.Type.Kind() == reflect.Struct && (!got || dtag == "") {
			spath := path
			if path != "" {
				spath += "."
			}
			spath += field.Name
			subs := StructNonDefFields(vvp.Interface(), spath)
			if len(subs) > 0 {
				flds = append(flds, subs...)
			}
			return true
		}
		if !got {
			return true
		}
		def, defStr := StructFieldIsDef(dtag, vvp.Interface(), field.Type.Kind())
		if def {
			return true
		}
		flds = append(flds, StructFieldVals{Path: path, Field: field, Val: vvp, Defs: defStr})
		return true
	})
	return flds
}

// StructNonDefFieldsStr processses "default" tag for default value(s) of fields in
// given struct, and returns a string of all those not at their default values,
// in format: Path.Field: val // default value(s)
// Uses a recursive strategy -- any fields that are themselves structs are
// expanded, and the field name represented by dots path separators.
func StructNonDefFieldsStr(structPtr any, path string) string {
	flds := StructNonDefFields(structPtr, path)
	if len(flds) == 0 {
		return ""
	}
	str := ""
	for _, fld := range flds {
		pth := fld.Path
		fnm := fld.Field.Name
		val := reflectx.ToStringPrec(fld.Val.Interface(), 6)
		dfs := fld.Defs
		if len(pth) > 0 {
			fnm = pth + "." + fnm
		}
		str += fmt.Sprintf("%s: %s // %s<br>\n", fnm, val, dfs)
	}
	return str
}

// StructViewDialog opens a dialog (optionally in a new, separate window)
// for viewing / editing the given struct object, in the context of given ctx widget
func StructViewDialog(ctx core.Widget, stru any, title string, newWindow bool) {
	d := core.NewBody().AddTitle(title)
	NewStructView(d).SetStruct(stru)
	if tb, ok := stru.(core.Toolbarer); ok {
		d.AddAppBar(tb.ConfigToolbar)
	}
	ds := d.NewFullDialog(ctx)
	if newWindow {
		ds.SetNewWindow(true)
	}
	ds.Run()
}
