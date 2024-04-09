// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tree

import (
	"fmt"
	"slices"

	"cogentcore.org/core/gti"
)

// Slice is just a slice of tree nodes: []Node, providing methods for accessing
// elements in the slice, and JSON marshal / unmarshal with encoding of
// underlying types
type Slice []Node

// NOTE: we have to define [*Slice] functions operating on a generic *[]Node
// element as the first (not receiver) argument, to be able to use these
// functions in any other types that are based on [Slice] or are other forms
// of []Node.

// SliceIsValidIndex checks whether the given index is a valid index into slice,
// within range of 0..len-1.  Returns error if not.
func SliceIsValidIndex(sl *[]Node, idx int) error {
	if idx >= 0 && idx < len(*sl) {
		return nil
	}
	return fmt.Errorf("tree.Slice: invalid index: %v with len = %v", idx, len(*sl))
}

// IsValidIndex checks whether the given index is a valid index into slice,
// within range of 0..len-1.  Returns error if not.
func (sl *Slice) IsValidIndex(idx int) error {
	if idx >= 0 && idx < len(*sl) {
		return nil
	}
	return fmt.Errorf("tree.Slice: invalid index: %v with len = %v", idx, len(*sl))
}

// Elem returns element at index; panics if index is invalid.
func (sl *Slice) Elem(idx int) Node {
	return (*sl)[idx]
}

// ElemTry returns element at index; returns error if index is invalid.
func (sl *Slice) ElemTry(idx int) (Node, error) {
	if err := sl.IsValidIndex(idx); err != nil {
		return nil, err
	}
	return (*sl)[idx], nil
}

// ElemFromEnd returns element at index from end of slice (0 = last element,
// 1 = 2nd to last, etc).  Panics if invalid index.
func (sl *Slice) ElemFromEnd(idx int) Node {
	return (*sl)[len(*sl)-1-idx]
}

// ElemFromEndTry returns element at index from end of slice (0 = last element,
// 1 = 2nd to last, etc). Try version returns error on invalid index.
func (sl *Slice) ElemFromEndTry(idx int) (Node, error) {
	return sl.ElemTry(len(*sl) - 1 - idx)
}

// SliceIndexByFunc finds index of item based on match function (which must
// return true for a find match, false for not).  Returns false if not found.
// startIndex arg allows for optimized bidirectional find if you have an idea
// where it might be, which can be key speedup for large lists. If no value
// is specified for startIndex, it starts in the middle, which is a good default.
func SliceIndexByFunc(sl *[]Node, match func(k Node) bool, startIndex ...int) (int, bool) {
	sz := len(*sl)
	if sz == 0 {
		return -1, false
	}
	si := -1
	if len(startIndex) > 0 {
		si = startIndex[0]
	}
	if si < 0 {
		si = sz / 2
	}
	if si == 0 {
		for idx, child := range *sl {
			if match(child) {
				return idx, true
			}
		}
	} else {
		if si >= sz {
			si = sz - 1
		}
		upi := si + 1
		dni := si
		upo := false
		for {
			if !upo && upi < sz {
				if match((*sl)[upi]) {
					return upi, true
				}
				upi++
			} else {
				upo = true
			}
			if dni >= 0 {
				if match((*sl)[dni]) {
					return dni, true
				}
				dni--
			} else if upo {
				break
			}
		}
	}
	return -1, false
}

// IndexByFunc finds index of item based on match function (which must return
// true for a find match, false for not).  Returns false if not found.
// startIndex arg allows for optimized bidirectional find if you have an idea
// where it might be, which can be key speedup for large lists. If no value
// is specified for startIndex, it starts in the middle, which is a good default.
func (sl *Slice) IndexByFunc(match func(k Node) bool, startIndex ...int) (int, bool) {
	return SliceIndexByFunc((*[]Node)(sl), match, startIndex...)
}

// SliceIndexOf returns index of element in list, false if not there.  startIndex arg
// allows for optimized bidirectional find if you have an idea where it might
// be, which can be key speedup for large lists. If no value is specified for startIndex,
// it starts in the middle, which is a good default.
func SliceIndexOf(sl *[]Node, kid Node, startIndex ...int) (int, bool) {
	return SliceIndexByFunc(sl, func(ch Node) bool { return ch == kid }, startIndex...)
}

// IndexOf returns index of element in list, false if not there.  startIndex arg
// allows for optimized bidirectional find if you have an idea where it might
// be, which can be key speedup for large lists. If no value is specified for
// startIndex, it starts in the middle, which is a good default.
func (sl *Slice) IndexOf(kid Node, startIndex ...int) (int, bool) {
	return sl.IndexByFunc(func(ch Node) bool { return ch == kid }, startIndex...)
}

// SliceIndexByName returns index of first element that has given name, false if
// not found. See [Slice.IndexOf] for info on startIndex.
func SliceIndexByName(sl *[]Node, name string, startIndex ...int) (int, bool) {
	return SliceIndexByFunc(sl, func(ch Node) bool { return ch.Name() == name }, startIndex...)
}

// IndexByName returns index of first element that has given name, false if
// not found. See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) IndexByName(name string, startIndex ...int) (int, bool) {
	return sl.IndexByFunc(func(ch Node) bool { return ch.Name() == name }, startIndex...)
}

// IndexByType returns index of element that either is that type or embeds
// that type, false if not found. See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) IndexByType(t *gti.Type, embeds bool, startIndex ...int) (int, bool) {
	if embeds {
		return sl.IndexByFunc(func(ch Node) bool { return ch.NodeType().HasEmbed(t) }, startIndex...)
	}
	return sl.IndexByFunc(func(ch Node) bool { return ch.NodeType() == t }, startIndex...)
}

// ElemByName returns first element that has given name, nil if not found.
// See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) ElemByName(name string, startIndex ...int) Node {
	idx, ok := sl.IndexByName(name, startIndex...)
	if !ok {
		return nil
	}
	return (*sl)[idx]
}

// ElemByNameTry returns first element that has given name, error if not found.
// See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) ElemByNameTry(name string, startIndex ...int) (Node, error) {
	idx, ok := sl.IndexByName(name, startIndex...)
	if !ok {
		return nil, fmt.Errorf("tree.Slice: element named: %v not found", name)
	}
	return (*sl)[idx], nil
}

// ElemByType returns index of element that either is that type or embeds
// that type, nil if not found. See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) ElemByType(t *gti.Type, embeds bool, startIndex ...int) Node {
	idx, ok := sl.IndexByType(t, embeds, startIndex...)
	if !ok {
		return nil
	}
	return (*sl)[idx]
}

// ElemByTypeTry returns index of element that either is that type or embeds
// that type, error if not found. See [Slice.IndexOf] for info on startIndex.
func (sl *Slice) ElemByTypeTry(t *gti.Type, embeds bool, startIndex ...int) (Node, error) {
	idx, ok := sl.IndexByType(t, embeds, startIndex...)
	if !ok {
		return nil, fmt.Errorf("tree.Slice: element of type: %v not found", t)
	}
	return (*sl)[idx], nil
}

// Insert item at index; does not do any parent updating etc; use
// the [Node] or [NodeBase] method unless you know what you are doing.
func (sl *Slice) Insert(k Node, i int) {
	*sl = slices.Insert(*sl, i, k)
}

// SliceDeleteAtIndex deletes item at index; does not do any further management of
// deleted item. It is an optimized version for avoiding memory leaks. It returns
// an error if the index is invalid.
func SliceDeleteAtIndex(sl *[]Node, i int) error {
	if err := SliceIsValidIndex(sl, i); err != nil {
		return err
	}
	*sl = slices.Delete(*sl, i, i+1)
	return nil
}

// DeleteAtIndex deletes item at index; does not do any further management of
// deleted item. It is an optimized version for avoiding memory leaks. It returns
// an error if the index is invalid.
func (sl *Slice) DeleteAtIndex(idx int) error {
	return SliceDeleteAtIndex((*[]Node)(sl), idx)
}

// SliceMove moves element from one position to another.  Returns error if
// either index is invalid.
func SliceMove(sl *[]Node, frm, to int) error {
	if err := SliceIsValidIndex(sl, frm); err != nil {
		return err
	}
	if err := SliceIsValidIndex(sl, to); err != nil {
		return err
	}
	if frm == to {
		return nil
	}
	tmp := (*sl)[frm]
	SliceDeleteAtIndex(sl, frm)
	*sl = slices.Insert(*sl, to, tmp)
	return nil
}

// Move element from one position to another.  Returns error if either index
// is invalid.
func (sl *Slice) Move(frm, to int) error {
	return SliceMove((*[]Node)(sl), frm, to)
}

// SliceSwap swaps elements between positions.  Returns error if either index is invalid
func SliceSwap(sl *[]Node, i, j int) error {
	if err := SliceIsValidIndex(sl, i); err != nil {
		return err
	}
	if err := SliceIsValidIndex(sl, j); err != nil {
		return err
	}
	if i == j {
		return nil
	}
	(*sl)[j], (*sl)[i] = (*sl)[i], (*sl)[j]
	return nil
}

// Swap elements between positions.  Returns error if either index is invalid
func (sl *Slice) Swap(i, j int) error {
	return SliceSwap((*[]Node)(sl), i, j)
}

///////////////////////////////////////////////////////////////////////////
// Config

// Config is a major work-horse routine for minimally destructive reshaping of
// a tree structure to fit a target configuration, specified in terms of a
// type-and-name list. It returns whether any changes were made to the slice.
func (sl *Slice) Config(n Node, config Config) bool {
	mods := false
	// first make a map for looking up the indexes of the names
	nm := make(map[string]int)
	for i, tn := range config {
		nm[tn.Name] = i
	}
	// first remove any children not in the config
	sz := len(*sl)
	for i := sz - 1; i >= 0; i-- {
		kid := (*sl)[i]
		knm := kid.Name()
		ti, ok := nm[knm]
		if !ok {
			sl.configDeleteKid(kid, i, &mods)
		} else if kid.NodeType() != config[ti].Type {
			sl.configDeleteKid(kid, i, &mods)
		}
	}
	// next add and move items as needed -- in order so guaranteed
	for i, tn := range config {
		kidx, ok := sl.IndexByName(tn.Name, i)
		if !ok {
			mods = true
			nkid := NewOfType(tn.Type)
			nkid.SetName(tn.Name)
			initNode(nkid)
			sl.Insert(nkid, i)
			if n != nil {
				SetParent(nkid, n)
			}
		} else {
			if kidx != i {
				mods = true
				sl.Move(kidx, i)
			}
		}
	}
	return mods
}

func (sl *Slice) configDeleteKid(kid Node, i int, mods *bool) {
	*mods = true
	kid.Destroy()
	sl.DeleteAtIndex(i)
}

// CopyFrom another Slice.  It is efficient by using the Config method
// which attempts to preserve any existing nodes in the destination
// if they have the same name and type -- so a copy from a source to
// a target that only differ minimally will be minimally destructive.
// it is essential that child names are unique.
func (sl *Slice) CopyFrom(frm Slice) {
	sl.ConfigCopy(nil, frm)
	for i, kid := range *sl {
		fmk := frm[i]
		kid.CopyFrom(fmk)
	}
}

// ConfigCopy uses Config method to copy name / type config of Slice from source
// If n is != nil then Update etc is called properly.
// it is essential that child names are unique.
func (sl *Slice) ConfigCopy(n Node, frm Slice) {
	sz := len(frm)
	if sz > 0 || n == nil {
		cfg := make(Config, sz)
		for i, kid := range frm {
			cfg[i].Type = kid.NodeType()
			cfg[i].Name = kid.Name()
		}
		sl.Config(n, cfg)
	} else {
		n.DeleteChildren()
	}
}
