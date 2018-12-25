// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syms

import "github.com/goki/ki"

// Type contains all the information about types.  Types can be builtin
// or composed of builtin types.  Each type can have one or more elements,
// e.g., fields for a struct or class, multiple values for a go function,
// or the two types for a map (key, value), etc..
type Type struct {
	Name  string   `desc:"name of the type -- can be the name of a field or the role for a type element"`
	Kind  Kinds    `desc:"kind of type -- overall nature of the type"`
	Desc  string   `desc:"documentation about this type, extracted from code"`
	Els   TypeEls  `desc:"elements of this type -- ordering and meaning varies depending on the Kind of type -- for Primitive types this is the parent type, for Composite types it describes the key elements of the type: Tuple = each element's type; Array = type of elements; Struct = each field, etc (see docs for each in Kinds)"`
	Size  []int    `desc:"for primitive types, this is the number of bits, for composite types, it is the number of elements, which can be multi-dimensional in some cases"`
	Props ki.Props `desc:"additional type properties, such as const, virtual, static -- these are just recorded textually and not systematized to keep things open-ended -- many of the most important properties can be inferred from the Kind property"`
	Ast   ki.Ki    `json:"-" xml:"-" desc:"Ast node that corresponds to this type -- only valid during parsing"`
}

// NewType returns a new Type struct initialized with given name and kind
func NewType(name string, kind Kinds) *Type {
	ty := &Type{Name: name, Kind: kind}
	return ty
}

// String() satisfies the fmt.Stringer interface
func (ty *Type) String() string {
	return ty.Name + ": " + ty.Kind.String()
}

//////////////////////////////////////////////////////////////////////////////////
// TypeEls

// TypeEl is a type element -- has a name (local to the type, e.g., field name)
// and a type name that can be looked up in a master list of types
type TypeEl struct {
	Name string `desc:"element name -- e.g., field name for struct, or functional name for other types"`
	Type string `desc:"type name -- looked up on relevant lists -- includes scoping / package / namespace name as appropriate"`
}

// TypeEls are the type elements for types
type TypeEls []TypeEl

// Add adds a new type element
func (te *TypeEls) Add(nm, typ string) {
	(*te) = append(*te, TypeEl{Name: nm, Type: typ})
}
