// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"fmt"
	"image"
	"maps"
	"reflect"
	"strings"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/slicesx"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
)

// Node is the interface for all SVG nodes.
type Node interface {
	tree.Node

	// AsNodeBase returns the [NodeBase] for our node, which gives
	// access to all the base-level data structures and methods
	// without requiring interface methods.
	AsNodeBase() *NodeBase

	// Render draws the node to the svg image.
	Render(sv *SVG)

	// BBoxes computes BBox and VisBBox during Render.
	BBoxes(sv *SVG)

	// LocalBBox returns the bounding box of node in local dimensions.
	LocalBBox() math32.Box2

	// NodeBBox returns the bounding box in image coordinates for this node.
	NodeBBox(sv *SVG) image.Rectangle

	// SetNodePos sets the upper left effective position of this element, in local dimensions.
	SetNodePos(pos math32.Vector2)

	// SetNodeSize sets the overall effective size of this element, in local dimensions.
	SetNodeSize(sz math32.Vector2)

	// ApplyTransform applies the given 2D transform to the geometry of this node
	// this just does a direct transform multiplication on coordinates.
	ApplyTransform(sv *SVG, xf math32.Matrix2)

	// ApplyDeltaTransform applies the given 2D delta transforms to the geometry of this node
	// relative to given point.  Trans translation and point are in top-level coordinates,
	// so must be transformed into local coords first.
	// Point is upper left corner of selection box that anchors the translation and scaling,
	// and for rotation it is the center point around which to rotate.
	ApplyDeltaTransform(sv *SVG, trans math32.Vector2, scale math32.Vector2, rot float32, pt math32.Vector2)

	// WriteGeom writes the geometry of the node to a slice of floating point numbers
	// the length and ordering of which is specific to each node type.
	// Slice must be passed and will be resized if not the correct length.
	WriteGeom(sv *SVG, dat *[]float32)

	// ReadGeom reads the geometry of the node from a slice of floating point numbers
	// the length and ordering of which is specific to each node type.
	ReadGeom(sv *SVG, dat []float32)

	// SVGName returns the SVG element name (e.g., "rect", "path" etc).
	SVGName() string

	// EnforceSVGName returns true if in general this element should
	// be named with its SVGName plus a unique id.
	// Groups and Markers are false.
	EnforceSVGName() bool
}

// NodeBase is the base type for all elements within an SVG tree.
// It implements the [Node] interface and contains the core functionality.
type NodeBase struct {
	tree.NodeBase

	// Class contains user-defined class name(s) used primarily for attaching
	// CSS styles to different display elements.
	// Multiple class names can be used to combine properties;
	// use spaces to separate per css standard.
	Class string

	// CSS is the cascading style sheet at this level.
	// These styles apply here and to everything below, until superceded.
	// Use .class and #name Properties elements to apply entire styles
	// to given elements, and type for element type.
	CSS map[string]any `xml:"css" set:"-"`

	// CSSAgg is the aggregated css properties from all higher nodes down to this node.
	CSSAgg map[string]any `copier:"-" json:"-" xml:"-" set:"-" display:"no-inline"`

	// BBox is the bounding box for the node within the SVG Pixels image.
	// This one can be outside the visible range of the SVG image.
	// VisBBox is intersected and only shows visible portion.
	BBox image.Rectangle `copier:"-" json:"-" xml:"-" set:"-"`

	// VisBBox is the visible bounding box for the node intersected with the SVG image geometry.
	VisBBox image.Rectangle `copier:"-" json:"-" xml:"-" set:"-"`

	// Paint is the paint style information for this node.
	Paint styles.Paint `json:"-" xml:"-" set:"-"`

	// isDef is whether this is in [SVG.Defs].
	isDef bool
}

func (n *NodeBase) AsNodeBase() *NodeBase {
	return n
}

func (n *NodeBase) SVGName() string {
	return "base"
}

func (n *NodeBase) EnforceSVGName() bool {
	return true
}

func (n *NodeBase) SetPos(pos math32.Vector2) {
}

func (n *NodeBase) SetSize(sz math32.Vector2) {
}

func (n *NodeBase) LocalBBox() math32.Box2 {
	bb := math32.Box2{}
	return bb
}

func (n *NodeBase) BaseInterface() reflect.Type {
	return reflect.TypeOf((*NodeBase)(nil)).Elem()
}

func (n *NodeBase) PaintStyle() *styles.Paint {
	return &n.Paint
}

func (n *NodeBase) Init() {
	n.Paint.Defaults()
}

// SetColorProperties sets color property from a string representation.
// It breaks color alpha out as opacity.  prop is either "stroke" or "fill"
func (n *NodeBase) SetColorProperties(prop, color string) {
	clr := errors.Log1(colors.FromString(color))
	n.SetProperty(prop+"-opacity", fmt.Sprintf("%g", float32(clr.A)/255))
	// we have consumed the A via opacity, so we reset it to 255
	clr.A = 255
	n.SetProperty(prop, colors.AsHex(clr))
}

// ParTransform returns the full compounded 2D transform matrix for all
// of the parents of this node.  If self is true, then include our
// own transform too.
func (n *NodeBase) ParTransform(self bool) math32.Matrix2 {
	pars := []Node{}
	xf := math32.Identity2()
	node := n.This.(Node)
	for {
		if node.AsTree().Parent == nil {
			break
		}
		node = node.AsTree().Parent.(Node)
		pars = append(pars, node)
	}
	np := len(pars)
	if np > 0 {
		xf = pars[np-1].AsNodeBase().PaintStyle().Transform
	}
	for i := np - 2; i >= 0; i-- {
		n := pars[i]
		xf.SetMul(n.AsNodeBase().PaintStyle().Transform)
	}
	if self {
		xf.SetMul(n.Paint.Transform)
	}
	return xf
}

// ApplyTransform applies the given 2D transform to the geometry of this node
// this just does a direct transform multiplication on coordinates.
func (n *NodeBase) ApplyTransform(sv *SVG, xf math32.Matrix2) {
}

// DeltaTransform computes the net transform matrix for given delta transform parameters
// and the transformed version of the reference point.  If self is true, then
// include the current node self transform, otherwise don't.  Groups do not
// but regular rendering nodes do.
func (n *NodeBase) DeltaTransform(trans math32.Vector2, scale math32.Vector2, rot float32, pt math32.Vector2, self bool) (math32.Matrix2, math32.Vector2) {
	mxi := n.ParTransform(self)
	mxi = mxi.Inverse()
	lpt := mxi.MulVector2AsPoint(pt)
	ldel := mxi.MulVector2AsVector(trans)
	xf := math32.Scale2D(scale.X, scale.Y).Rotate(rot)
	xf.X0 = ldel.X
	xf.Y0 = ldel.Y
	return xf, lpt
}

// ApplyDeltaTransform applies the given 2D delta transforms to the geometry of this node
// relative to given point.  Trans translation and point are in top-level coordinates,
// so must be transformed into local coords first.
// Point is upper left corner of selection box that anchors the translation and scaling,
// and for rotation it is the center point around which to rotate
func (n *NodeBase) ApplyDeltaTransform(sv *SVG, trans math32.Vector2, scale math32.Vector2, rot float32, pt math32.Vector2) {
}

// WriteTransform writes the node transform to slice at starting index.
// slice must already be allocated sufficiently.
func (n *NodeBase) WriteTransform(dat []float32, idx int) {
	dat[idx+0] = n.Paint.Transform.XX
	dat[idx+1] = n.Paint.Transform.YX
	dat[idx+2] = n.Paint.Transform.XY
	dat[idx+3] = n.Paint.Transform.YY
	dat[idx+4] = n.Paint.Transform.X0
	dat[idx+5] = n.Paint.Transform.Y0
}

// ReadTransform reads the node transform from slice at starting index.
func (n *NodeBase) ReadTransform(dat []float32, idx int) {
	n.Paint.Transform.XX = dat[idx+0]
	n.Paint.Transform.YX = dat[idx+1]
	n.Paint.Transform.XY = dat[idx+2]
	n.Paint.Transform.YY = dat[idx+3]
	n.Paint.Transform.X0 = dat[idx+4]
	n.Paint.Transform.Y0 = dat[idx+5]
}

// WriteGeom writes the geometry of the node to a slice of floating point numbers
// the length and ordering of which is specific to each node type.
// Slice must be passed and will be resized if not the correct length.
func (n *NodeBase) WriteGeom(sv *SVG, dat *[]float32) {
	*dat = slicesx.SetLength(*dat, 6)
	n.WriteTransform(*dat, 0)
}

// ReadGeom reads the geometry of the node from a slice of floating point numbers
// the length and ordering of which is specific to each node type.
func (n *NodeBase) ReadGeom(sv *SVG, dat []float32) {
	n.ReadTransform(dat, 0)
}

// SVGWalkDown does [tree.NodeBase.WalkDown] on given node using given walk function
// with SVG Node parameters.
func SVGWalkDown(n Node, fun func(sn Node, snb *NodeBase) bool) {
	n.AsTree().WalkDown(func(n tree.Node) bool {
		sn := n.(Node)
		return fun(sn, sn.AsNodeBase())
	})
}

// SVGWalkDownNoDefs does [tree.Node.WalkDown] on given node using given walk function
// with SVG Node parameters.  Automatically filters Defs nodes (IsDef) and MetaData,
// i.e., it only processes concrete graphical nodes.
func SVGWalkDownNoDefs(n Node, fun func(sn Node, snb *NodeBase) bool) {
	n.AsTree().WalkDown(func(cn tree.Node) bool {
		sn := cn.(Node)
		snb := sn.AsNodeBase()
		_, md := sn.(*MetaData)
		if snb.isDef || md {
			return tree.Break
		}
		return fun(sn, snb)
	})
}

// FirstNonGroupNode returns the first item that is not a group
// recursing into groups until a non-group item is found.
func FirstNonGroupNode(n Node) Node {
	var ngn Node
	SVGWalkDownNoDefs(n, func(sn Node, snb *NodeBase) bool {
		if _, isgp := sn.(*Group); isgp {
			return tree.Continue
		}
		ngn = sn
		return tree.Break
	})
	return ngn
}

// NodesContainingPoint returns all Nodes with Bounding Box that contains
// given point, optionally only those that are terminal nodes (no leaves).
// Excludes the starting node.
func NodesContainingPoint(n Node, pt image.Point, leavesOnly bool) []Node {
	var cn []Node
	SVGWalkDown(n, func(sn Node, snb *NodeBase) bool {
		if sn == n {
			return tree.Continue
		}
		if leavesOnly && snb.HasChildren() {
			return tree.Continue
		}
		if snb.Paint.Off {
			return tree.Break
		}
		if pt.In(snb.BBox) {
			cn = append(cn, sn)
		}
		return tree.Continue
	})
	return cn
}

//////////////////////////////////////////////////////////////////
// Standard Node infrastructure

// Style styles the Paint values directly from node properties
func (n *NodeBase) Style(sv *SVG) {
	pc := &n.Paint
	pc.Defaults()
	ctxt := colors.Context(sv)
	pc.StyleSet = false // this is always first call, restart

	var parCSSAgg map[string]any
	if n.Parent != nil { // && g.Par != sv.Root.This
		pn := n.Parent.(Node)
		parCSSAgg = pn.AsNodeBase().CSSAgg
		pp := pn.AsNodeBase().PaintStyle()
		pc.CopyStyleFrom(pp)
		pc.SetStyleProperties(pp, n.Properties, ctxt)
	} else {
		pc.SetStyleProperties(nil, n.Properties, ctxt)
	}
	pc.ToDotsImpl(&pc.UnitContext) // we always inherit parent's unit context -- SVG sets it once-and-for-all

	if parCSSAgg != nil {
		AggCSS(&n.CSSAgg, parCSSAgg)
	} else {
		n.CSSAgg = nil
	}
	AggCSS(&n.CSSAgg, n.CSS)
	n.StyleCSS(sv, n.CSSAgg)

	pc.StrokeStyle.Opacity *= pc.FontStyle.Opacity // applies to all
	pc.FillStyle.Opacity *= pc.FontStyle.Opacity

	pc.Off = !pc.Display || (pc.StrokeStyle.Color == nil && pc.FillStyle.Color == nil)
}

// AggCSS aggregates css properties
func AggCSS(agg *map[string]any, css map[string]any) {
	if *agg == nil {
		*agg = make(map[string]any)
	}
	maps.Copy(*agg, css)
}

// ApplyCSS applies css styles to given node,
// using key to select sub-properties from overall properties list
func (n *NodeBase) ApplyCSS(sv *SVG, key string, css map[string]any) bool {
	pp, got := css[key]
	if !got {
		return false
	}
	pmap, ok := pp.(map[string]any) // must be a properties map
	if !ok {
		return false
	}
	pc := &n.Paint
	ctxt := colors.Context(sv)
	if n.Parent != sv.Root.This {
		pp := n.Parent.(Node).AsNodeBase().PaintStyle()
		pc.SetStyleProperties(pp, pmap, ctxt)
	} else {
		pc.SetStyleProperties(nil, pmap, ctxt)
	}
	return true
}

// StyleCSS applies css style properties to given SVG node
// parsing out type, .class, and #name selectors
func (n *NodeBase) StyleCSS(sv *SVG, css map[string]any) {
	tyn := strings.ToLower(n.NodeType().Name) // type is most general, first
	n.ApplyCSS(sv, tyn, css)
	cln := "." + strings.ToLower(n.Class) // then class
	n.ApplyCSS(sv, cln, css)
	idnm := "#" + strings.ToLower(n.Name) // then name
	n.ApplyCSS(sv, idnm, css)
}

// LocalBBoxToWin converts a local bounding box to SVG coordinates
func (n *NodeBase) LocalBBoxToWin(bb math32.Box2) image.Rectangle {
	mxi := n.ParTransform(true) // include self
	return bb.MulMatrix2(mxi).ToRect()
}

func (n *NodeBase) NodeBBox(sv *SVG) image.Rectangle {
	rs := &sv.RenderState
	return rs.LastRenderBBox
}

func (n *NodeBase) SetNodePos(pos math32.Vector2) {
	// no-op by default
}

func (n *NodeBase) SetNodeSize(sz math32.Vector2) {
	// no-op by default
}

// LocalLineWidth returns the line width in local coordinates
func (n *NodeBase) LocalLineWidth() float32 {
	pc := &n.Paint
	if pc.StrokeStyle.Color == nil {
		return 0
	}
	return pc.StrokeStyle.Width.Dots
}

// ComputeBBox is called by default in render to compute bounding boxes for
// gui interaction -- can only be done in rendering because that is when all
// the proper transforms are all in place -- VpBBox is intersected with parent SVG
func (n *NodeBase) BBoxes(sv *SVG) {
	if n.This == nil {
		return
	}
	ni := n.This.(Node)
	n.BBox = ni.NodeBBox(sv)
	n.BBox.Canon()
	n.VisBBox = sv.Geom.SizeRect().Intersect(n.BBox)
}

// PushTransform checks our bounding box and visibility, returning false if
// out of bounds.  If visible, pushes our transform.
// Must be called as first step in Render.
func (n *NodeBase) PushTransform(sv *SVG) (bool, *paint.Context) {
	n.BBox = image.Rectangle{}
	if n.Paint.Off || n == nil || n.This == nil {
		return false, nil
	}
	ni := n.This.(Node)
	// if g.IsInvisible() { // just the Invisible flag
	// 	return false, nil
	// }
	lbb := ni.LocalBBox()
	n.BBox = n.LocalBBoxToWin(lbb)
	n.VisBBox = sv.Geom.SizeRect().Intersect(n.BBox)
	nvis := n.VisBBox == image.Rectangle{}
	// g.SetInvisibleState(nvis) // don't set

	if nvis && !n.isDef {
		return false, nil
	}

	rs := &sv.RenderState
	rs.PushTransform(n.Paint.Transform)

	pc := &paint.Context{rs, &n.Paint}
	return true, pc
}

func (n *NodeBase) RenderChildren(sv *SVG) {
	for _, kid := range n.Children {
		ni := kid.(Node)
		ni.Render(sv)
	}
}

func (n *NodeBase) Render(sv *SVG) {
	vis, rs := n.PushTransform(sv)
	if !vis {
		return
	}
	// pc := &g.Paint
	// render path elements, then compute bbox, then fill / stroke
	n.BBoxes(sv)
	n.RenderChildren(sv)
	rs.PopTransform()
}
