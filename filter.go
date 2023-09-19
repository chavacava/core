// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"goki.dev/ki/v2/ki"
)

// Filter represents SVG filter* elements
type Filter struct {
	NodeBase
	FilterType string
}

// AddNewFilter adds a new filter to given parent node, with given name.
func AddNewFilter(parent ki.Ki, name string) *Filter {
	return parent.AddNewChild(FilterType, name).(*Filter)
}

func (g *Filter) SVGName() string { return "filter" }

func (g *Filter) CopyFieldsFrom(frm any) {
	fr := frm.(*Filter)
	g.NodeBase.CopyFieldsFrom(&fr.NodeBase)
	g.FilterType = fr.FilterType
}
