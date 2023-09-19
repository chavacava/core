// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"fmt"
	"log"
	"strings"

	"github.com/srwiley/rasterx"
	"goki.dev/girl/gist"
	"goki.dev/ki/v2/ki"
	"goki.dev/laser"
	"goki.dev/mat32/v2"
)

// FindDefByName finds Defs item by name, using cached indexes for speed
func (sv *SVG) FindDefByName(defnm string) Node {
	if sv.DefIdxs == nil {
		sv.DefIdxs = make(map[string]int)
	}
	idx, has := sv.DefIdxs[defnm]
	if !has {
		idx = len(sv.Defs.Kids) / 2
	}
	idx, has = sv.Defs.Kids.IndexByName(defnm, idx)
	if has {
		sv.DefIdxs[defnm] = idx
		return sv.Defs.Kids[idx].(Node)
	}
	delete(sv.DefIdxs, defnm) // not found -- delete from map
	return nil
}

func (sv *SVG) FindNamedElement(name string) Node {
	name = strings.TrimPrefix(name, "#")
	def := sv.FindDefByName(name)
	log.Printf("SVG FindNamedElement: could not find name: %v\n", name)
	return nil
}

// NameFromURL returns just the name referred to in a url(#name)
// if it is not a url(#) format then returns empty string.
func NameFromURL(url string) string {
	if len(url) < 7 {
		return ""
	}
	if url[:5] != "url(#" {
		return ""
	}
	ref := url[5:]
	sz := len(ref)
	if ref[sz-1] == ')' {
		ref = ref[:sz-1]
	}
	return ref
}

// NameToURL returns url as: url(#name)
func NameToURL(nm string) string {
	return "url(#" + nm + ")"
}

// NodeFindURL finds a url element in the parent SVG of given node.
// Returns nil if not found.
// Works with full 'url(#Name)' string or plain name or "none"
func NodeFindURL(gi Node, url string) Node {
	if url == "none" {
		return nil
	}
	ref := NameFromURL(url)
	if ref == "" {
		ref = url
	}
	if ref == "" {
		return nil
	}
	psvg := ParentSVG(g)
	var rv Node
	if psvg != nil {
		rv = psvg.FindNamedElement(ref)
	} else {
		rv = g.FindNamedElement(ref)
	}
	if rv == nil {
		log.Printf("svg.NodeFindURL could not find element named: %v in parents of svg el: %v\n", url, gii.Path())
	}
	return rv
}

// NodePropURL returns a url(#name) url from given prop name on node,
// or empty string if none.  Returned value is just the 'name' part
// of the url, not the full string.
func NodePropURL(kn ki.Ki, prop string) string {
	fp, err := kn.PropTry(prop)
	if err != nil {
		return ""
	}
	fs, iss := fp.(string)
	if !iss {
		return ""
	}
	return NameFromURL(fs)
}

// MarkerByName finds marker property of given name, or generic "marker"
// type, and if set, attempts to find that marker and return it
func MarkerByName(gii Node, marker string) *Marker {
	url := NodePropURL(gii, marker)
	if url == "" {
		url = NodePropURL(gii, "marker")
	}
	if url == "" {
		return nil
	}
	mrkn := NodeFindURL(gii, url)
	if mrkn == nil {
		return nil
	}
	mrk, ok := mrkn.(*Marker)
	if !ok {
		log.Printf("gi.svg Found element named: %v but isn't a Marker type, instead is: %T", url, mrkn)
		return nil
	}
	return mrk
}

//////////////////////////////////////////////////////////////////////////////
//  Gradient management utilities for updating geometry

// GradientByName returns the gradient of given name, stored on SVG node
func GradientByName(gi Node, grnm string) *Gradient {
	gri := NodeFindURL(gi, grnm)
	if gri == nil {
		return nil
	}
	gr, ok := gri.(*Gradient)
	if !ok {
		log.Printf("SVG Found element named: %v but isn't a Gradient type, instead is: %T", grnm, gri)
		return nil
	}
	return gr
}

// GradientApplyXForm applies the given transform to any gradients for this node,
// that are using specific coordinates (not bounding box which is automatic)
func (g *NodeBase) GradientApplyXForm(xf mat32.Mat2) {
	gi := g.This().(Node)
	gnm := NodePropURL(gi, "fill")
	if gnm != "" {
		gr := GradientByName(gi, gnm)
		if gr != nil {
			gr.Grad.ApplyXForm(xf)
		}
	}
	gnm = NodePropURL(gi, "stroke")
	if gnm != "" {
		gr := GradientByName(gi, gnm)
		if gr != nil {
			gr.Grad.ApplyXForm(xf)
		}
	}
}

// GradientApplyXFormPt applies the given transform with ctr point
// to any gradients for this node, that are using specific coordinates
// (not bounding box which is automatic)
func (g *NodeBase) GradientApplyXFormPt(xf mat32.Mat2, pt mat32.Vec2) {
	gi := g.This().(Node)
	gnm := NodePropURL(gi, "fill")
	if gnm != "" {
		gr := GradientByName(gi, gnm)
		if gr != nil {
			gr.Grad.ApplyXFormPt(xf, pt)
		}
	}
	gnm = NodePropURL(gi, "stroke")
	if gnm != "" {
		gr := GradientByName(gi, gnm)
		if gr != nil {
			gr.Grad.ApplyXFormPt(xf, pt)
		}
	}
}

// GradientWritePoints writes the UserSpaceOnUse gradient points to
// a slice of floating point numbers, appending to end of slice.
func GradientWritePts(gr *gist.ColorSpec, dat *[]float32) {
	if gr.Gradient == nil {
		return
	}
	if gr.Gradient.Units == rasterx.ObjectBoundingBox {
		return
	}
	*dat = append(*dat, float32(gr.Gradient.Matrix.A))
	*dat = append(*dat, float32(gr.Gradient.Matrix.B))
	*dat = append(*dat, float32(gr.Gradient.Matrix.C))
	*dat = append(*dat, float32(gr.Gradient.Matrix.D))
	*dat = append(*dat, float32(gr.Gradient.Matrix.E))
	*dat = append(*dat, float32(gr.Gradient.Matrix.F))
	if !gr.Gradient.IsRadial {
		for i := 0; i < 4; i++ {
			*dat = append(*dat, float32(gr.Gradient.Points[i]))
		}
	}
}

// GradientWritePts writes the geometry of the gradients for this node
// to a slice of floating point numbers, appending to end of slice.
func (g *NodeBase) GradientWritePts(dat *[]float32) {
	gnm := NodePropURL(g, "fill")
	if gnm != "" {
		gr := GradientByName(g, gnm)
		if gr != nil {
			GradientWritePts(&gr.Grad, dat)
		}
	}
	gnm = NodePropURL(g, "stroke")
	if gnm != "" {
		gr := GradientByName(g, gnm)
		if gr != nil {
			GradientWritePts(&gr.Grad, dat)
		}
	}
}

// GradientReadPoints reads the UserSpaceOnUse gradient points from
// a slice of floating point numbers, reading from the end.
func GradientReadPts(gr *gist.ColorSpec, dat []float32) {
	if gr.Gradient == nil {
		return
	}
	if gr.Gradient.Units == rasterx.ObjectBoundingBox {
		return
	}
	sz := len(dat)
	n := 6
	if !gr.Gradient.IsRadial {
		n = 10
		for i := 0; i < 4; i++ {
			gr.Gradient.Points[i] = float64(dat[(sz-4)+i])
		}
	}
	gr.Gradient.Matrix.A = float64(dat[(sz-n)+0])
	gr.Gradient.Matrix.B = float64(dat[(sz-n)+1])
	gr.Gradient.Matrix.C = float64(dat[(sz-n)+2])
	gr.Gradient.Matrix.D = float64(dat[(sz-n)+3])
	gr.Gradient.Matrix.E = float64(dat[(sz-n)+4])
	gr.Gradient.Matrix.F = float64(dat[(sz-n)+5])
}

// GradientReadPts reads the geometry of the gradients for this node
// from a slice of floating point numbers, reading from the end.
func (g *NodeBase) GradientReadPts(dat []float32) {
	gnm := NodePropURL(g, "fill")
	if gnm != "" {
		gr := GradientByName(g, gnm)
		if gr != nil {
			GradientReadPts(&gr.Grad, dat)
		}
	}
	gnm = NodePropURL(g, "stroke")
	if gnm != "" {
		gr := GradientByName(g, gnm)
		if gr != nil {
			GradientReadPts(&gr.Grad, dat)
		}
	}
}

//////////////////////////////////////////////////////////////////////////////
//  Gradient management utilities for creating element-specific grads

// UpdateGradientStops copies stops from StopsName gradient if it is set
func UpdateGradientStops(gr *Gradient) {
	if gr.StopsName == "" {
		return
	}
	sgr := GradientByName(gr, gr.StopsName)
	if sgr != nil {
		gr.Grad.CopyStopsFrom(&sgr.Grad)
	}
}

// DeleteNodeGradient deletes the node-specific gradient on given node
// of given name, which can be a full url(# name or just the bare name.
// Returns true if deleted.
func DeleteNodeGradient(gi Node, grnm string) bool {
	gr := GradientByName(gi, grnm)
	if gr == nil || gr.StopsName == "" {
		return false
	}
	psvg := ParentSVG(gi)
	if psvg == nil {
		return false
	}
	unm := NameFromURL(grnm)
	psvg.Defs.DeleteChildByName(unm, ki.DestroyKids)
	return true
}

// AddNewNodeGradient adds a new gradient specific to given node
// that points to given stops name.  returns the new gradient
// and the url that points to it (nil if parent svg cannot be found).
// Initializes gradient to use bounding box of object, but using userSpaceOnUse setting
func AddNewNodeGradient(gi Node, radial bool, stops string) (*Gradient, string) {
	psvg := ParentSVG(gi)
	if psvg == nil {
		return nil, ""
	}
	gr, url := psvg.AddNewGradient(radial)
	gr.StopsName = stops
	bbox := gi.(Node).LocalBBox()
	gr.Grad.SetGradientPoints(bbox)
	UpdateGradientStops(gr)
	return gr, url
}

// AddNewGradient adds a new gradient, either linear or radial,
// with a new unique id
func (sv *SVG) AddNewGradient(radial bool) (*Gradient, string) {
	gnm := ""
	if radial {
		gnm = "radialGradient"
	} else {
		gnm = "linearGradient"
	}
	updt := sv.UpdateStart()
	id := sv.NewUniqueId()
	gnm = NameId(gnm, id)
	sv.SetChildAdded()
	gr := sv.Defs.AddNewChild(TypeGradient, gnm).(*Gradient)
	url := NameToURL(gnm)
	if radial {
		gr.Grad.NewRadialGradient()
	} else {
		gr.Grad.NewLinearGradient()
	}
	sv.UpdateEnd(updt)
	return gr, url
}

// UpdateNodeGradientProp ensures that node has a gradient property of given type
func UpdateNodeGradientProp(gi Node, prop string, radial bool, stops string) (*Gradient, string) {
	ps := gi.Prop(prop)
	if ps == nil {
		gr, url := AddNewNodeGradient(gi, radial, stops)
		gi.SetProp(prop, url)
		return gr, url
	}
	pstr := ps.(string)
	trgst := ""
	if radial {
		trgst = "radialGradient"
	} else {
		trgst = "linearGradient"
	}
	url := "url(#" + trgst
	if strings.HasPrefix(pstr, url) {
		gr := GradientByName(gi, pstr)
		gr.StopsName = stops
		UpdateGradientStops(gr)
		return gr, NameToURL(gr.Nm)
	}
	if strings.HasPrefix(pstr, "url(#") { // wrong kind
		DeleteNodeGradient(gi, pstr)
	}
	gr, url := AddNewNodeGradient(gi, radial, stops)
	gi.SetProp(prop, url)
	return gr, url
}

// UpdateNodeGradientPoints updates the points for node based on current bbox
func UpdateNodeGradientPoints(gi Node, prop string) {
	ps := gi.Prop(prop)
	if ps == nil {
		return
	}
	pstr := ps.(string)
	url := "url(#"
	if !strings.HasPrefix(pstr, url) {
		return
	}
	gr := GradientByName(gi, pstr)
	if gr == nil {
		return
	}
	bbox := gi.(Node).LocalBBox()
	gr.Grad.SetGradientPoints(bbox)
	gr.Grad.Gradient.Matrix = rasterx.Identity
}

// CloneNodeGradientProp creates a new clone of the existing gradient for node
// if set for given property key ("fill" or "stroke").
// returns new gradient.
func CloneNodeGradientProp(gi Node, prop string) *Gradient {
	ps := gi.Prop(prop)
	if ps == nil {
		return nil
	}
	pstr := ps.(string)
	radial := false
	if strings.HasPrefix(pstr, "url(#radialGradient") {
		radial = true
	} else if !strings.HasPrefix(pstr, "url(#linearGradient") {
		return nil
	}
	gr := GradientByName(gi, pstr)
	if gr == nil {
		return nil
	}
	ngr, url := AddNewNodeGradient(gi, radial, gr.StopsName)
	gi.SetProp(prop, url)
	ngr.Grad.CopyFrom(&gr.Grad)
	return gr
}

// DeleteNodeGradientProp deletes any existing gradient for node
// if set for given property key ("fill" or "stroke").
// Returns true if deleted.
func DeleteNodeGradientProp(gi Node, prop string) bool {
	ps := gi.Prop(prop)
	if ps == nil {
		return false
	}
	pstr := ps.(string)
	if !strings.HasPrefix(pstr, "url(#radialGradient") && !strings.HasPrefix(pstr, "url(#linearGradient") {
		return false
	}
	return DeleteNodeGradient(gi, pstr)
}

// UpdateAllGradientStops removes any items from Defs that are not actually referred to
// by anything in the current SVG tree.  Returns true if items were removed.
// Does not remove gradients with StopsName = "" with extant stops -- these
// should be removed manually, as they are not automatically generated.
func (sv *SVG) UpdateAllGradientStops() {
	for _, k := range sv.Defs.Kids {
		gr, ok := k.(*Gradient)
		if ok {
			UpdateGradientStops(gr)
		}
	}
}

const SVGRefCountKey = "SVGRefCount"

func IncRefCount(k ki.Ki) {
	rc := k.Prop(SVGRefCountKey).(int)
	rc++
	k.SetProp(SVGRefCountKey, rc)
}

// RemoveOrphanedDefs removes any items from Defs that are not actually referred to
// by anything in the current SVG tree.  Returns true if items were removed.
// Does not remove gradients with StopsName = "" with extant stops -- these
// should be removed manually, as they are not automatically generated.
func (sv *SVG) RemoveOrphanedDefs() bool {
	updt := sv.UpdateStart()
	sv.SetFullReRender()
	refkey := SVGRefCountKey
	for _, k := range sv.Defs.Kids {
		k.SetProp(refkey, 0)
	}
	sv.FuncDownMeFirst(0, nil, func(k ki.Ki, level int, d any) bool {
		pr := k.Properties()
		for _, v := range *pr {
			ps := laser.ToString(v)
			if !strings.HasPrefix(ps, "url(#") {
				continue
			}
			nm := NameFromURL(ps)
			el := sv.FindDefByName(nm)
			if el != nil {
				IncRefCount(el)
			}
		}
		if gr, isgr := k.(*Gradient); isgr {
			if gr.StopsName != "" {
				el := sv.FindDefByName(gr.StopsName)
				if el != nil {
					IncRefCount(el)
				}
			} else {
				if gr.Grad.Gradient != nil && len(gr.Grad.Gradient.Stops) > 0 {
					IncRefCount(k) // keep us around
				}
			}
		}
		return ki.Continue
	})
	sz := len(sv.Defs.Kids)
	del := false
	for i := sz - 1; i >= 0; i-- {
		k := sv.Defs.Kids[i]
		rc := k.Prop(refkey).(int)
		if rc == 0 {
			fmt.Printf("Deleting unused item: %s\n", k.Name())
			sv.Defs.Kids.DeleteAtIndex(i)
			del = true
		} else {
			k.DeleteProp(refkey)
		}
	}
	sv.UpdateEnd(updt)
	return del
}
