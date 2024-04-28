// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"bytes"
	"html"
	"os"
	"strings"
	"text/template"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/strcase"
	"cogentcore.org/core/units"
)

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString(
		`// Code generated by "go run gen.go"; DO NOT EDIT.

// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package units
	`)
	for _, v := range units.UnitsValues() {
		// we ignore dots because they also set the dots field
		// as a special case
		if v == units.UnitDot {
			continue
		}
		s := v.String()
		d := data{
			Lower: s,
			Camel: strcase.ToCamel(s),
		}
		// actual desc after "represents"
		_, d.Desc, _ = strings.Cut(v.Desc(), " represents ")
		d.Desc = html.UnescapeString(d.Desc)
		errors.Must(funcs.Execute(buf, d))
	}
	errors.Must(os.WriteFile("unitgen.go", buf.Bytes(), 0666))
}

type data struct {
	Lower string
	Camel string
	Desc  string
}

var funcs = template.Must(template.New("funcs").Parse(
	`
// {{.Camel}} returns a new {{.Lower}} value.
// {{.Camel}} is {{.Desc}}
func {{.Camel}}(value float32) Value {
	return Value{Value: value, Unit: Unit{{.Camel}}}
}

// {{.Camel}} sets the value in terms of {{.Lower}}.
// {{.Camel}} is {{.Desc}}
func (v *Value) {{.Camel}}(value float32) {
	v.Value = value
	v.Unit = Unit{{.Camel}}
}

// {{.Camel}} converts the given {{.Lower}} value to dots.
// {{.Camel}} is {{.Desc}}
func (uc *Context) {{.Camel}}(value float32) float32 {
	return uc.ToDots(value, Unit{{.Camel}})
}
`))
