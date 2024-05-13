// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// ValueWidgeter is an interface that types can implement to specify the
// [ValueWidget] that should be used to represent them in the GUI.
type ValueWidgeter interface {

	// ValueWidget returns the [ValueWidget] that should be used to represent
	// the value in the GUI. If it returns nil, then [ToValueWidget] will
	// fall back onto the next step. This function does NOT need to call [core.Bind].
	ValueWidget() ValueWidget
}

// ValueWidgetTypes is a map of functions that return a [ValueWidget]
// for a value of a certain fully package path qualified type name.
// It is used by [ToValueWidget]. If a function returns nil, it falls
// back onto the next step. You can add to this using the [AddValueWidgetType]
// helper function.
var ValueWidgetTypes = map[string]func(value any) ValueWidget{}

// ValueWidgetConverters is a slice of functions that return a [ValueWidget]
// for a value. It is used by [ToValueWidget]. If a function returns nil,
// it falls back on the next function in the slice, and if all functions return nil,
// it falls back on the default bindings. These functions do NOT need to call
// [core.Bind].
var ValueWidgetConverters []func(value any) ValueWidget

// AddValueWidgetType binds the given value type to the given [ValueWidget] type,
// meaning that [ToValueWidget] will return a new [ValueWidget] of the given type
// when it receives values of the given value type. It uses [ValueWidgetTypes].
// This function is called with various standard types automatically.
func AddValueWidgetType[T any, W ValueWidget]() {
	var v T
	name := types.TypeNameValue(v)
	ValueWidgetTypes[name] = func(value any) ValueWidget {
		return tree.New[W]()
	}
}

func init() {
	AddValueWidgetType[string, *Text]()
	AddValueWidgetType[bool, *Switch]()
}

// ToValueWidget converts the given value into an appropriate [ValueWidget]
// whose associated value is bound to the given value. It first checks the
// [ValueWidgeter] interface, then the [ValueWidgetTypes], then the
// [ValueWidgetConverters], and finally it falls back on a set of default
// bindings. If any step results in nil, it falls back on the next step.
func ToValueWidget(value any) ValueWidget {
	if vwr, ok := value.(ValueWidgeter); ok {
		if vw := vwr.ValueWidget(); vw != nil {
			return Bind(value, vw)
		}
	}
	if vwt, ok := ValueWidgetTypes[types.TypeNameValue(value)]; ok {
		if vw := vwt(value); vw != nil {
			return Bind(value, vw)
		}
	}
	for _, converter := range ValueWidgetConverters {
		if vw := converter(value); vw != nil {
			return Bind(value, vw)
		}
	}
	return nil
}
