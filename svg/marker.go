// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"log"

	"cogentcore.org/core/math32"
)

// Marker represents marker elements that can be drawn along paths (arrow heads, etc)
type Marker struct {
	NodeBase

	// reference position to align the vertex position with, specified in ViewBox coordinates
	RefPos math32.Vector2 `xml:"{refX,refY}"`

	// size of marker to render, in Units units
	Size math32.Vector2 `xml:"{markerWidth,markerHeight}"`

	// units to use
	Units MarkerUnits `xml:"markerUnits"`

	// viewbox defines the internal coordinate system for the drawing elements within the marker
	ViewBox ViewBox

	// orientation of the marker -- either 'auto' or an angle
	Orient string `xml:"orient"`

	// current vertex position
	VertexPos math32.Vector2

	// current vertex angle in radians
	VertexAngle float32

	// current stroke width
	StrokeWidth float32

	// net transform computed from settings and current values -- applied prior to rendering
	Transform math32.Matrix2

	// effective size for actual rendering
	EffSize math32.Vector2
}

func (mkr *Marker) SVGName() string { return "marker" }

func (mkr *Marker) EnforceSVGName() bool { return false }

// MarkerUnits specifies units to use for svg marker elements
type MarkerUnits int32 //enum: enum

const (
	StrokeWidth MarkerUnits = iota
	UserSpaceOnUse
	MarkerUnitsN
)

// RenderMarker renders the marker using given vertex position, angle (in
// radians), and stroke width
func (mkr *Marker) RenderMarker(sv *SVG, vertexPos math32.Vector2, vertexAng, strokeWidth float32) {
	mkr.VertexPos = vertexPos
	mkr.VertexAngle = vertexAng
	mkr.StrokeWidth = strokeWidth
	if mkr.Units == StrokeWidth {
		mkr.EffSize = mkr.Size.MulScalar(strokeWidth)
	} else {
		mkr.EffSize = mkr.Size
	}

	ang := vertexAng
	if mkr.Orient != "auto" {
		ang, _ = math32.ParseAngle32(mkr.Orient)
	}
	if mkr.ViewBox.Size == (math32.Vector2{}) {
		mkr.ViewBox.Size = math32.Vec2(3, 3)
	}
	mkr.Transform = math32.Rotate2D(ang).Scale(mkr.EffSize.X/mkr.ViewBox.Size.X, mkr.EffSize.Y/mkr.ViewBox.Size.Y).Translate(-mkr.RefPos.X, -mkr.RefPos.Y)
	mkr.Transform.X0 += vertexPos.X
	mkr.Transform.Y0 += vertexPos.Y

	mkr.Paint.Transform = mkr.Transform

	// fmt.Println("render marker:", mrk.Name, strokeWidth)
	mkr.Render(sv)
}

func (mkr *Marker) Render(sv *SVG) {
	pc := &mkr.Paint
	rs := &sv.RenderState
	rs.PushTransform(pc.Transform)

	mkr.RenderChildren(sv)
	mkr.BBoxes(sv) // must come after render

	rs.PopTransform()
}

func (mkr *Marker) BBoxes(sv *SVG) {
	if mkr.This == nil {
		return
	}
	ni := mkr.This.(Node)
	mkr.BBox = ni.NodeBBox(sv)
	mkr.BBox.Canon()
	mkr.VisBBox = sv.Geom.SizeRect().Intersect(mkr.BBox)
}

//////////////////////////////////////////////////////////
// 	SVG marker management

// MarkerByName finds marker property of given name, or generic "marker"
// type, and if set, attempts to find that marker and return it
func (sv *SVG) MarkerByName(n Node, marker string) *Marker {
	url := NodePropURL(n, marker)
	if url == "" {
		url = NodePropURL(n, "marker")
	}
	if url == "" {
		return nil
	}
	mrkn := sv.NodeFindURL(n, url)
	if mrkn == nil {
		return nil
	}
	mrk, ok := mrkn.(*Marker)
	if !ok {
		log.Printf("SVG Found element named: %v but isn't a Marker type, instead is: %T", url, mrkn)
		return nil
	}
	return mrk
}
