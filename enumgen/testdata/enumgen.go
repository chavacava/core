// Code generated by "enumgen -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

const (
	_DaysName_0      = "DAY_SUNDAY"
	_DaysLowerName_0 = "day_sunday"
	_DaysName_1      = "DAY_MONDAY"
	_DaysLowerName_1 = "day_monday"
	_DaysName_2      = "DAY_TUESDAY"
	_DaysLowerName_2 = "day_tuesday"
	_DaysName_3      = "DAY_WEDNESDAY"
	_DaysLowerName_3 = "day_wednesday"
	_DaysName_4      = "DAY_THURSDAY"
	_DaysLowerName_4 = "day_thursday"
	_DaysName_5      = "DAY_FRIDAY"
	_DaysLowerName_5 = "day_friday"
	_DaysName_6      = "DAY_SATURDAY"
	_DaysLowerName_6 = "day_saturday"
)

var (
	_DaysIndex_0 = [...]uint8{0, 10}
	_DaysIndex_1 = [...]uint8{0, 10}
	_DaysIndex_2 = [...]uint8{0, 11}
	_DaysIndex_3 = [...]uint8{0, 13}
	_DaysIndex_4 = [...]uint8{0, 12}
	_DaysIndex_5 = [...]uint8{0, 10}
	_DaysIndex_6 = [...]uint8{0, 12}
)

// String returns the string representation
// of this Days value.
func (i Days) String() string {
	switch {
	case i == 1:
		return _DaysName_0
	case i == 3:
		return _DaysName_1
	case i == 5:
		return _DaysName_2
	case i == 7:
		return _DaysName_3
	case i == 9:
		return _DaysName_4
	case i == 11:
		return _DaysName_5
	case i == 13:
		return _DaysName_6
	default:
		return "Days(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _DaysNoOp() {
	var x [1]struct{}
	_ = x[Sunday-(1)]
	_ = x[Monday-(3)]
	_ = x[Tuesday-(5)]
	_ = x[Wednesday-(7)]
	_ = x[Thursday-(9)]
	_ = x[Friday-(11)]
	_ = x[Saturday-(13)]
}

var _DaysValues = []Days{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// DaysN is the total number of
// enum values for type Days.
const DaysN Days = 7

var _DaysNameToValueMap = map[string]Days{
	_DaysName_0[0:10]:      Sunday,
	_DaysLowerName_0[0:10]: Sunday,
	_DaysName_1[0:10]:      Monday,
	_DaysLowerName_1[0:10]: Monday,
	_DaysName_2[0:11]:      Tuesday,
	_DaysLowerName_2[0:11]: Tuesday,
	_DaysName_3[0:13]:      Wednesday,
	_DaysLowerName_3[0:13]: Wednesday,
	_DaysName_4[0:12]:      Thursday,
	_DaysLowerName_4[0:12]: Thursday,
	_DaysName_5[0:10]:      Friday,
	_DaysLowerName_5[0:10]: Friday,
	_DaysName_6[0:12]:      Saturday,
	_DaysLowerName_6[0:12]: Saturday,
}

var _DaysNames = []string{
	_DaysName_0[0:10],
	_DaysName_1[0:10],
	_DaysName_2[0:11],
	_DaysName_3[0:13],
	_DaysName_4[0:12],
	_DaysName_5[0:10],
	_DaysName_6[0:12],
}

var _DaysDescMap = map[Days]string{
	1:  _DaysDescs[0],
	3:  _DaysDescs[1],
	5:  _DaysDescs[2],
	7:  _DaysDescs[3],
	9:  _DaysDescs[4],
	11: _DaysDescs[5],
	13: _DaysDescs[6],
}

var _DaysDescs = []string{
	`Sunday is the first day of the week`,
	`Monday is the second day of the week`,
	`Tuesday is the third day of the week`,
	`Wednesday is the fourth day of the week`,
	`Thursday is the fifth day of the week`,
	`Friday is the sixth day of the week`,
	`Saturday is the seventh day of the week`,
}

// SetString sets the Days value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Days) SetString(s string) error {
	if val, ok := _DaysNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _DaysNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " does not belong to Days values")
}

// Int64 returns the Days value as an int64.
func (i Days) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Days value from an int64.
func (i *Days) SetInt64(in int64) {
	*i = Days(in)
}

// Desc returns the description of the Days value.
func (i Days) Desc() string {
	if str, ok := _DaysDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DaysValues returns all possible values of
// the type Days. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on Days.
func DaysValues() []Days {
	return _DaysValues
}

// Values returns all possible values of
// type Days. This slice will be in the
// same order as those returned by Strings and Descs.
func (i Days) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DaysValues))
	for i, d := range _DaysValues {
		res[i] = &d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i Days) Strings() []string {
	return _DaysNames
}

// Descs returns the descriptions of all
// possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i Days) Descs() []string {
	return _DaysDescs
}

// IsValid returns whether the value is a
// valid option for type Days.
func (i Days) IsValid() bool {
	for _, v := range _DaysValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Days) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Days) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalGQL implements the [graphql.Marshaler] interface.
func (i Days) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(i.String())))
}

// UnmarshalGQL implements the [graphql.Unmarshaler] interface.
func (i *Days) UnmarshalGQL(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Days should be a string, but got a value of type %T instead", value)
	}
	return i.SetString(str)
}

const _StatesName = "enablednot-enabledfocusedveredcurrently-being-pressed-by-userselected"

var _StatesIndex = [...]uint8{0, 7, 18, 25, 30, 61, 69}

const _StatesLowerName = "enablednot-enabledfocusedveredcurrently-being-pressed-by-userselected"

// String returns the string representation
// of this States value.
func (i States) String() string {
	if i < 0 || i >= States(len(_StatesIndex)-1) {
		return "States(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatesName[_StatesIndex[i]:_StatesIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatesNoOp() {
	var x [1]struct{}
	_ = x[Enabled-(0)]
	_ = x[Disabled-(1)]
	_ = x[Focused-(2)]
	_ = x[Hovered-(3)]
	_ = x[Active-(4)]
	_ = x[Selected-(5)]
}

var _StatesValues = []States{Enabled, Disabled, Focused, Hovered, Active, Selected}

// StatesN is the total number of
// enum values for type States.
const StatesN States = 6

var _StatesNameToValueMap = map[string]States{
	_StatesName[0:7]:        Enabled,
	_StatesLowerName[0:7]:   Enabled,
	_StatesName[7:18]:       Disabled,
	_StatesLowerName[7:18]:  Disabled,
	_StatesName[18:25]:      Focused,
	_StatesLowerName[18:25]: Focused,
	_StatesName[25:30]:      Hovered,
	_StatesLowerName[25:30]: Hovered,
	_StatesName[30:61]:      Active,
	_StatesLowerName[30:61]: Active,
	_StatesName[61:69]:      Selected,
	_StatesLowerName[61:69]: Selected,
}

var _StatesNames = []string{
	_StatesName[0:7],
	_StatesName[7:18],
	_StatesName[18:25],
	_StatesName[25:30],
	_StatesName[30:61],
	_StatesName[61:69],
}

var _StatesDescMap = map[States]string{
	0: _StatesDescs[0],
	1: _StatesDescs[1],
	2: _StatesDescs[2],
	3: _StatesDescs[3],
	4: _StatesDescs[4],
	5: _StatesDescs[5],
}

var _StatesDescs = []string{
	`Enabled indicates the widget is enabled`,
	`Disabled indicates the widget is disabled`,
	`Focused indicates the widget has keyboard focus`,
	`Hovered indicates the widget is being hovered over`,
	`Active indicates the widget is being interacted with`,
	`Selected indicates the widget is selected`,
}

// SetString sets the States value from its
// string representation, and returns an
// error if the string is invalid.
func (i *States) SetString(s string) error {
	if val, ok := _StatesNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _StatesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " does not belong to States values")
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) {
	*i = States(in)
}

// Desc returns the description of the States value.
func (i States) Desc() string {
	if str, ok := _StatesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatesValues returns all possible values of
// the type States. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on States.
func StatesValues() []States {
	return _StatesValues
}

// Values returns all possible values of
// type States. This slice will be in the
// same order as those returned by Strings and Descs.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = &d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type States.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i States) Strings() []string {
	return _StatesNames
}

// Descs returns the descriptions of all
// possible values of type States.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i States) Descs() []string {
	return _StatesDescs
}

// IsValid returns whether the value is a
// valid option for type States.
func (i States) IsValid() bool {
	for _, v := range _StatesValues {
		if i == v {
			return true
		}
	}
	return false
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i *States) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *States) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// Scan implements the [driver.Valuer] interface.
func (i States) Value() (driver.Value, error) {
	return i.String(), nil
}

// Value implements the [sql.Scanner] interface.
func (i *States) Scan(value any) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value for type States: %[1]T(%[1]v)", value)
	}

	return i.SetString(str)
}
