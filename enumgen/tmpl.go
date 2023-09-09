// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on http://github.com/dmarkham/enumer and
// golang.org/x/tools/cmd/stringer:

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"fmt"
	"log"
	"text/template"
)

// TmplData contains the data passed to a generation template
type TmplData struct {
	TypeName      string // the name of the enum type
	IsBitFlag     bool   // whether the type is a bit flag
	MaxValueP1    string // the highest defined value for the type, plus one, as a string
	MethodName    string // method name (String or BitIndexString)
	MethodComment string // doc comment for the method
	IfInvalid     string // the code for what to do if the value is invalid (used in String and SetString)
	Slice         string // the code for the slice to use/return (used in Values and bit flag String)
	Extends       string // the type this type extends, if any
}

// ExecTmpl executes the given template with the given data and
// writes the result to [Generator.Buf]. It fatally logs any error.
func (g *Generator) ExecTmpl(t *template.Template, data *TmplData) {
	err := t.Execute(&g.Buf, data)
	if err != nil {
		log.Fatalf("programmer error: internal error: error executing template: %v", err)
	}
}

// SetMethod sets [TmplData.MethodName] and [TmplData.MethodComment]
// based on whether the type is a bit flag type. It assumes that
// [TmplData.TypeName] and [TmplData.IsBitFlag] are already set.
func (td *TmplData) SetMethod() {
	if td.IsBitFlag {
		td.MethodName = "BitIndexString"
		td.MethodComment = fmt.Sprintf(`// BitIndexString returns the string
		// representation of this %s value
		// if it is a bit index value
		// (typically an enum constant), and
		// not an actual bit flag value.`, td.TypeName)
	} else {
		td.MethodName = "String"
		td.MethodComment = fmt.Sprintf(`// String returns the string representation
		// of this %s value.`, td.TypeName)
	}
}
