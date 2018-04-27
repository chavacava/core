// Copyright (c) 2018, Randall C. O'Reilly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kit

// github.com/rcoreilly/goki/ki/kit

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/rcoreilly/goki/ki/bitflag"
)

// design notes: for methods that return string, not passing error b/c you can
// easily check for null string, and registering errors in log for setter
// methods, returning error and also logging so it is safe to ignore err if
// you don't care

// Bit flags are setup just using the ordinal count iota, and the only diff is
// the methods which do 1 << flag when operating on them
// see bitflag package

// EnumRegistry is a map from an enum-style const int type name to a
// corresponding reflect.Type and conversion methods generated by (modified)
// stringer that convert to / from strings -- need to explicitly register each
// new type by calling AddEnum in the process of creating a new global
// variable, as in:
//
// var KiT_MyEnum = kit.Enums.AddEnum(MyEnumN, bitFlag true/false,
//    TypeNameProps (or nil))
//
// where MyEnum is the name of the type, MyEnumN is the enum value
// representing the number of defined enums (always good practice to define
// this value, for ease of extension by others), and TypeNameProps is nil or a
// map[string]interface{} of properties, OR:
//
// var KiT_MyEnum = kit.Enums.AddEnumAltLower(MyEnumN, bitFlag true/false,
//    TypeNameProps, "Prefix")
//
// which automatically registers alternative names as lower-case versions of
// const names with given prefix removed -- often what is used in e.g., json
// or xml kinds of formats
//
// special properties:
//
// * "N": max value of enum defined -- number of enum entries (assuming
// ordinal, which is all that is currently supported here)
//
// * "BitFlag": true -- each value represents a bit in a set of bit flags, so
// the string rep of a value contains an or-list of names for each bit set,
// separated by |
//
// * "AltStrings": map[int64]string -- provides an alternative string mapping for
// the enum values
//
type EnumRegistry struct {
	Enums map[string]reflect.Type
	// Props contains properties that can be associated with each enum type -- e.g., "BitFlag": true  --  "AltStrings" : map[int64]string, or other custom settings
	Props map[string]map[string]interface{}
	// Vals contains cached EnumValue representations of the enum values -- used by EnumValues method
	Vals map[string][]EnumValue
}

// Enums is master registry of enum types -- can also create your own package-specific ones
var Enums EnumRegistry

// AddEnum adds a given type to the registry -- requires the N value to set N
// from and grab type info from -- if bitFlag then sets BitFlag property, and
// each value represents a bit in a set of bit flags, so the string rep of a
// value contains an or-list of names for each bit set, separated by | -- can
// also add additional properties -- they are copied so can be re-used across enums
func (tr *EnumRegistry) AddEnum(en interface{}, bitFlag bool, props map[string]interface{}) reflect.Type {
	if tr.Enums == nil {
		tr.Enums = make(map[string]reflect.Type)
		tr.Props = make(map[string]map[string]interface{})
		tr.Vals = make(map[string][]EnumValue)
	}

	// get the pointer-to version and elem so it is a settable type!
	typ := PtrType(reflect.TypeOf(en)).Elem()
	n := EnumToInt64(en)
	tn := FullTypeName(typ)
	tr.Enums[tn] = typ
	if props != nil {
		// make a copy of props for enums -- often shared
		nwprops := make(map[string]interface{}, len(props))
		for key, val := range props {
			nwprops[key] = val
		}
		tr.Props[tn] = nwprops
	}
	tp := tr.Properties(tn)
	tp["N"] = n
	if bitFlag {
		tp := tr.Properties(tn)
		tp["BitFlag"] = true
	}
	// fmt.Printf("added enum: %v with n: %v\n", tn, n)
	return typ
}

// AddEnumAltLower adds a given type to the registry -- requires the N value
// to set N from and grab type info from -- automatically initializes
// AltStrings alternative string map based on the name with given prefix
// removed (e.g., a type name-based prefix) and lower-cased -- also requires
// the number of enums -- assumes starts at 0
func (tr *EnumRegistry) AddEnumAltLower(en interface{}, bitFlag bool, props map[string]interface{}, prefix string) reflect.Type {
	typ := tr.AddEnum(en, bitFlag, props)
	n := EnumToInt64(en)
	tn := FullTypeName(typ)
	alts := make(map[int64]string)
	tp := tr.Properties(tn)
	for i := int64(0); i < n; i++ {
		str := EnumInt64ToString(i, typ)
		str = strings.ToLower(strings.TrimPrefix(str, prefix))
		// fmt.Printf("adding enum: %v\n", str)
		alts[i] = str
	}
	tp["AltStrings"] = alts
	return typ
}

// Enum finds an enum type based on its type name -- returns nil if not found
func (tr *EnumRegistry) Enum(name string) reflect.Type {
	return tr.Enums[name]
}

// TypeRegistered returns true if the given type is registered as an enum type
func (tr *EnumRegistry) TypeRegistered(typ reflect.Type) bool {
	enumName := FullTypeName(typ)
	_, ok := tr.Enums[enumName]
	// if ok {
	// 	fmt.Printf("enum type: %v registered\n", enumName)
	// }
	return ok
}

// Props returns properties for this type -- makes props map if not already made
func (tr *EnumRegistry) Properties(enumName string) map[string]interface{} {
	tp, ok := tr.Props[enumName]
	if !ok {
		tp = make(map[string]interface{})
		tr.Props[enumName] = tp
	}
	return tp
}

// Prop safely finds an enum type property from enum type name and property
// key -- nil if not found
func (tr *EnumRegistry) Prop(enumName, propKey string) interface{} {
	tp, ok := tr.Props[enumName]
	if !ok {
		// fmt.Printf("no props for enum type: %v\n", enumName)
		return nil
	}
	p, ok := tp[propKey]
	if !ok {
		// fmt.Printf("no props for key: %v\n", propKey)
		return nil
	}
	return p
}

// AltStrings returns optional alternative string map for enums -- e.g.,
// lower-case, without prefixes etc -- can put multiple such alt strings in
// the one string with your own separator, in a predefined order, if
// necessary, and just call strings.Split on those and get the one you want --
// nil if not set
func (tr *EnumRegistry) AltStrings(enumName string) map[int64]string {
	ps := tr.Prop(enumName, "AltStrings")
	if ps == nil {
		return nil
	}
	m, ok := ps.(map[int64]string)
	if !ok {
		log.Printf("kit.EnumRegistry AltStrings error: AltStrings property must be a map[int64]string type, is not -- is instead: %T\n", m)
		return nil
	}
	return m
}

// NVals returns the number of defined enum values
func (tr *EnumRegistry) NVals(eval interface{}) int64 {
	typ := reflect.TypeOf(eval)
	n, _ := ToInt(tr.Prop(FullTypeName(typ), "N"))
	return n
}

// IsBitFlag checks if this enum is for bit flags instead of mutually-exclusive int
// values -- checks BitFlag property -- if true string rep of a value contains
// an or-list of names for each bit set, separated by |
func (tr *EnumRegistry) IsBitFlag(eval interface{}) bool {
	tn := FullTypeName(reflect.TypeOf(eval))
	b, _ := ToBool(tr.Prop(tn, "BitFlag"))
	return b
}

// EnumToInt64 converts an enum into an int64 using reflect -- just use int64(eval) when you
// have the enum in hand -- this is when you just have a generic item
func EnumToInt64(eval interface{}) int64 {
	ev := NonPtrValue(reflect.ValueOf(eval))
	var ival int64
	reflect.ValueOf(&ival).Elem().Set(ev.Convert(reflect.TypeOf(ival)))
	return ival
}

// SetEnumFromInt64 sets enum value from int64 value -- must pass a pointer to
// the enum and also needs raw type of the enum as well -- can't get it from
// the interface{} reliably
func SetEnumFromInt64(eval interface{}, ival int64, typ reflect.Type) error {
	if reflect.TypeOf(eval).Kind() != reflect.Ptr {
		err := fmt.Errorf("kit.SetEnumFromInt64: must pass a pointer to the enum: Type: %v, Kind: %v\n", reflect.TypeOf(eval).Name(), reflect.TypeOf(eval).Kind())
		log.Printf("%v", err)
		return err
	}
	reflect.ValueOf(eval).Elem().Set(reflect.ValueOf(ival).Convert(typ))
	return nil
}

// EnumIfaceFromInt64 returns an interface{} value which is an enum value of
// given type, set to given integer value
func EnumIfaceFromInt64(ival int64, typ reflect.Type) interface{} {
	evnp := reflect.New(PtrType(typ))
	evpi := evnp.Interface()
	evn := reflect.New(typ)
	evi := evn.Interface()
	evpi = &evi
	reflect.ValueOf(evpi).Elem().Set(reflect.ValueOf(ival).Convert(typ))
	return evi
}

// EnumInt64ToString first converts an int64 to enum of given type, and then
// converts that to a string value
func EnumInt64ToString(ival int64, typ reflect.Type) string {
	evnp := reflect.New(PtrType(typ))
	evpi := evnp.Interface()
	evn := reflect.New(typ)
	evi := evn.Interface()
	evpi = &evi
	SetEnumFromInt64(evpi, ival, typ)
	return EnumToString(evi)
}

// EnumToString converts an enum value to its corresponding string value --
// you could just call fmt.Sprintf("%v") too but this is slightly faster, and
// it also works for bitflags which regular stringer does not
func EnumToString(eval interface{}) string {
	strer, ok := eval.(fmt.Stringer) // will fail if not impl
	if !ok {
		log.Printf("kit.EnumToString: fmt.Stringer interface not supported by type %v\n", reflect.TypeOf(eval).Name())
		return ""
	}
	return strer.String()
}

// BitFlagsToString converts an int64 of bit flags into a string
// representation of the bits that are set -- en is number of defined bits,
// and also provides the type name for looking up strings
func BitFlagsToString(bflg int64, en interface{}) string {
	et := PtrType(reflect.TypeOf(en)).Elem()
	n := int(EnumToInt64(en))
	str := ""
	for i := 0; i < n; i++ {
		if bitflag.Has(bflg, i) {
			evs := EnumInt64ToString(int64(i), et)
			if str == "" {
				str = evs
			} else {
				str += "|" + evs
			}
		}
	}
	return str
}

// note: convenience methods b/c it is easier to find on registry type

// EnumToString converts an enum value to its corresponding string value --
// you could just call fmt.Sprintf("%v") too but this is slightly faster
func (tr *EnumRegistry) EnumToString(eval interface{}) string {
	return EnumToString(eval)
}

// EnumToAltString converts an enum value to its corresponding alternative string value
func (tr *EnumRegistry) EnumToAltString(eval interface{}) string {
	if reflect.TypeOf(eval).Kind() == reflect.Ptr {
		eval = reflect.ValueOf(eval).Elem() // deref the pointer
	}
	et := reflect.TypeOf(eval)
	tn := FullTypeName(et)
	alts := tr.AltStrings(tn)
	if alts == nil {
		log.Printf("kit.EnumToAltString: no alternative string map for type %v\n", tn)
		return ""
	}
	// convert to int64 for lookup
	ival := EnumToInt64(eval)
	return alts[ival]
}

// EnumInt64ToAltString converts an int64 value to the enum of given type, and
// then into corresponding string value
func (tr *EnumRegistry) EnumInt64ToString(ival int64, typ reflect.Type) string {
	return EnumInt64ToString(ival, typ)
}

// EnumInt64ToAltString converts an int64 value to the enum of given type, and
// then into corresponding alternative string value
func (tr *EnumRegistry) EnumInt64ToAltString(ival int64, typnm string) string {
	alts := tr.AltStrings(typnm)
	if alts == nil {
		log.Printf("kit.EnumInt64ToAltString: no alternative string map for type %v\n", typnm)
		return ""
	}
	return alts[ival]
}

// SetEnumValueFromString sets enum value from string using reflect.Value
// IMPORTANT: requires the modified stringer go generate utility
// that generates a StringToTypeName method
func SetEnumValueFromString(eval reflect.Value, str string) error {
	etp := eval.Type()
	if etp.Kind() != reflect.Ptr {
		err := fmt.Errorf("kit.SetEnumValueFromString -- you must pass a pointer enum, not type: %v kind %v\n", etp, etp.Kind())
		// log.Printf("%v", err)
		return err
	}
	et := etp.Elem()
	methnm := "FromString"
	meth := eval.MethodByName(methnm)
	if ValueIsZero(meth) || meth.IsNil() {
		err := fmt.Errorf("kit.SetEnumValueFromString: stringer-generated FromString() method not found: %v for type: %v %T\n", methnm, et.Name(), eval.Interface())
		log.Printf("%v", err)
		return err
	}
	sv := reflect.ValueOf(str)
	args := make([]reflect.Value, 1)
	args[0] = sv
	meth.Call(args)
	// fmt.Printf("return from FromString method: %v\n", rv[0].Interface())
	return nil
}

// SetEnumValueFromString sets enum value from string, into a reflect.Value
// IMPORTANT: requires the modified stringer go generate utility
// that generates a StringToTypeName method
func (tr *EnumRegistry) SetEnumValueFromString(eval reflect.Value, str string) error {
	return SetEnumValueFromString(eval, str)
}

// SetEnumFromString sets enum value from string -- must pass a *pointer* to
// the enum item. IMPORTANT: requires the modified stringer go generate utility
// that generates a StringToTypeName method
func SetEnumFromString(eptr interface{}, str string) error {
	return SetEnumValueFromString(reflect.ValueOf(eptr), str)
}

// SetEnumFromString sets enum value from string -- must pass a *pointer* to
// the enum item. IMPORTANT: requires the modified stringer go generate utility
// that generates a StringToTypeName method
func (tr *EnumRegistry) SetEnumFromString(eptr interface{}, str string) error {
	return SetEnumFromString(eptr, str)
}

// SetEnumFromAltString sets from alternative string list using an interface{}
// to the enum -- must pass a *pointer* to the enum item.
func (tr *EnumRegistry) SetEnumFromAltString(eptr interface{}, str string) error {
	return tr.SetEnumValueFromAltString(reflect.ValueOf(eptr), str)
}

// SetEnumValueFromAltString sets value from alternative string using a
// reflect.Value -- must pass a *pointer* value to the enum item.
func (tr *EnumRegistry) SetEnumValueFromAltString(eval reflect.Value, str string) error {
	etp := eval.Type()
	if etp.Kind() != reflect.Ptr {
		err := fmt.Errorf("kit.SetEnumValueFromString -- you must pass a pointer enum, not type: %v kind %v\n", etp, etp.Kind())
		log.Printf("%v", err)
		return err
	}
	et := etp.Elem()
	tn := FullTypeName(et)
	alts := tr.AltStrings(tn)
	if alts == nil {
		err := fmt.Errorf("kit.SetEnumValueFromAltString: no alternative string map for type %v\n", tn)
		// log.Printf("%v", err)
		return err
	}
	for i, v := range alts {
		if v == str {
			return tr.SetEnumValueFromInt64(eval, int64(i))
		}
	}
	err := fmt.Errorf("kit.SetEnumValueFromAltString: string: %v not found in alt list of strings for type%v\n", str, tn)
	// log.Printf("%v", err)
	return err
}

// SetEnumValueFromStringAltFirst first attempts to set an enum from an
// alternative string, and if that fails, then it tries to set from the
// regular string representation func (tr *EnumRegistry)
func (tr *EnumRegistry) SetEnumValueFromStringAltFirst(eval reflect.Value, str string) error {
	err := tr.SetEnumValueFromAltString(eval, str)
	if err != nil {
		return tr.SetEnumValueFromString(eval, str)
	}
	return err
}

// SetEnumFromStringAltFirst first attempts to set an enum from an
// alternative string, and if that fails, then it tries to set from the
// regular string representation func (tr *EnumRegistry)
func (tr *EnumRegistry) SetEnumFromStringAltFirst(eptr interface{}, str string) error {
	err := tr.SetEnumFromAltString(eptr, str)
	if err != nil {
		return tr.SetEnumFromString(eptr, str)
	}
	return err
}

// SetEnumValueFromInt64 sets the enum value using reflect.Value
// representation from a generic int64 value
func (tr *EnumRegistry) SetEnumValueFromInt64(eval reflect.Value, ival int64) error {
	evi := eval.Interface()
	et := eval.Type().Elem()
	return SetEnumFromInt64(evi, ival, et)
}

///////////////////////////////////////////////////////////////////////////////
//  EnumValue

// EnumValue represents enum values, in common int64 terms, e.g., for GUI
type EnumValue struct {
	Name  string       `desc:"name for this value"`
	Value int64        `desc:"integer value"`
	Type  reflect.Type `desc:"the enum type that this value belongs to"`
}

// Set sets the values of the EnumValue struct
func (ev *EnumValue) Set(name string, val int64, typ reflect.Type) {
	ev.Name = name
	ev.Value = val
	ev.Type = typ
}

// String satisfies fmt.Stringer and provides a string representation of enum: just the name
func (ev EnumValue) String() string {
	return ev.Name
}

// Values returns an EnumValue slice for all the values of an enum type -- if
// alt is true and alt names exist, then those are used
func (tr *EnumRegistry) Values(enumName string, alt bool) []EnumValue {
	vals, ok := tr.Vals[enumName]
	if ok {
		return vals
	}
	alts := tr.AltStrings(enumName)
	et := tr.Enums[enumName]
	n := tr.Prop(enumName, "N").(int64)
	vals = make([]EnumValue, n)
	for i := int64(0); i < n; i++ {
		str := EnumInt64ToString(i, et) // todo: what happens when no string for given values?
		if alt && alts != nil {
			str = alts[i]
		}
		vals[i].Set(str, i, et)
	}
	tr.Vals[enumName] = vals
	return vals
}

// TypeValues returns an EnumValue slice for all the values of an enum type --
// if alt is true and alt names exist, then those are used
func (tr *EnumRegistry) TypeValues(et reflect.Type, alt bool) []EnumValue {
	return tr.Values(FullTypeName(et), alt)
}

// AllTagged returns a list of all registered enum types that include a given
// property key value -- does not check for the value of that value -- just
// its existence
func (tr *EnumRegistry) AllTagged(key string) []reflect.Type {
	tl := make([]reflect.Type, 0)
	for _, typ := range tr.Enums {
		tp := tr.Prop(FullTypeName(typ), key)
		if tp == nil {
			continue
		}
		tl = append(tl, typ)
	}
	return tl
}

/////////////////////////////////////////////////////////////
// Following is for testing..

// testing
type TestFlags int32

const (
	TestFlagsNil TestFlags = iota
	TestFlag1
	TestFlag2
	TestFlagsN
)

//go:generate stringer -type=TestFlags

var KiT_TestFlags = Enums.AddEnumAltLower(TestFlagsN, false, nil, "Test")
