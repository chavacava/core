// Code generated by "enumgen.test.exe -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _FruitsMap = map[Fruits]string{
	0: `Apple`,
	1: `Orange`,
	2: `Peach`,
	3: `Strawberry`,
	4: `Blackberry`,
	5: `Blueberry`,
	6: `Apricot`,
}

// String returns the string representation
// of this Fruits value.
func (i Fruits) String() string {
	if str, ok := _FruitsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FruitsNoOp() {
	var x [1]struct{}
	_ = x[Apple-(0)]
	_ = x[Orange-(1)]
	_ = x[Peach-(2)]
	_ = x[Strawberry-(3)]
	_ = x[Blackberry-(4)]
	_ = x[Blueberry-(5)]
	_ = x[Apricot-(6)]
}

var _FruitsValues = []Fruits{Apple, Orange, Peach, Strawberry, Blackberry, Blueberry, Apricot}

// FruitsN is the highest valid value
// for type Fruits, plus one.
const FruitsN Fruits = 7

var _FruitsNameToValueMap = map[string]Fruits{
	`Apple`:      0,
	`apple`:      0,
	`Orange`:     1,
	`orange`:     1,
	`Peach`:      2,
	`peach`:      2,
	`Strawberry`: 3,
	`strawberry`: 3,
	`Blackberry`: 4,
	`blackberry`: 4,
	`Blueberry`:  5,
	`blueberry`:  5,
	`Apricot`:    6,
	`apricot`:    6,
}

var _FruitsDescMap = map[Fruits]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
	5: ``,
	6: ``,
}

// SetString sets the Fruits value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Fruits) SetString(s string) error {
	if val, ok := _FruitsNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _FruitsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Fruits")
}

// Int64 returns the Fruits value as an int64.
func (i Fruits) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Fruits value from an int64.
func (i *Fruits) SetInt64(in int64) {
	*i = Fruits(in)
}

// Desc returns the description of the Fruits value.
func (i Fruits) Desc() string {
	if str, ok := _FruitsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FruitsValues returns all possible values
// for the type Fruits.
func FruitsValues() []Fruits {
	return _FruitsValues
}

// Values returns all possible values
// for the type Fruits.
func (i Fruits) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FruitsValues))
	for i, d := range _FruitsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Fruits.
func (i Fruits) IsValid() bool {
	_, ok := _FruitsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Fruits) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Fruits) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Fruits) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Fruits) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("Fruits should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}

var _FoodsMap = map[Foods]string{
	7:  `Bread`,
	8:  `Lettuce`,
	9:  `Cheese`,
	10: `Meat`,
}

// String returns the string representation
// of this Foods value.
func (i Foods) String() string {
	if str, ok := _FoodsMap[i]; ok {
		return str
	}
	return Fruits(i).String()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FoodsNoOp() {
	var x [1]struct{}
	_ = x[Bread-(7)]
	_ = x[Lettuce-(8)]
	_ = x[Cheese-(9)]
	_ = x[Meat-(10)]
}

var _FoodsValues = []Foods{Bread, Lettuce, Cheese, Meat}

// FoodsN is the highest valid value
// for type Foods, plus one.
const FoodsN Foods = 11

var _FoodsNameToValueMap = map[string]Foods{
	`Bread`:   7,
	`bread`:   7,
	`Lettuce`: 8,
	`lettuce`: 8,
	`Cheese`:  9,
	`cheese`:  9,
	`Meat`:    10,
	`meat`:    10,
}

var _FoodsDescMap = map[Foods]string{
	7:  ``,
	8:  ``,
	9:  ``,
	10: ``,
}

// SetString sets the Foods value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Foods) SetString(s string) error {
	if val, ok := _FoodsNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _FoodsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return (*Fruits)(i).SetString(s)
}

// Int64 returns the Foods value as an int64.
func (i Foods) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Foods value from an int64.
func (i *Foods) SetInt64(in int64) {
	*i = Foods(in)
}

// Desc returns the description of the Foods value.
func (i Foods) Desc() string {
	if str, ok := _FoodsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FoodsValues returns all possible values
// for the type Foods.
func FoodsValues() []Foods {
	return _FoodsValues
}

// Values returns all possible values
// for the type Foods.
func (i Foods) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FoodsValues))
	for i, d := range _FoodsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Foods.
func (i Foods) IsValid() bool {
	_, ok := _FoodsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Foods) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Foods) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Foods) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Foods) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("Foods should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}

var _DaysMap = map[Days]string{
	1:  `DAY_SUNDAY`,
	3:  `DAY_MONDAY`,
	5:  `DAY_TUESDAY`,
	7:  `DAY_WEDNESDAY`,
	9:  `DAY_THURSDAY`,
	11: `DAY_FRIDAY`,
	13: `DAY_SATURDAY`,
}

// String returns the string representation
// of this Days value.
func (i Days) String() string {
	if str, ok := _DaysMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
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

// DaysN is the highest valid value
// for type Days, plus one.
const DaysN Days = 14

var _DaysNameToValueMap = map[string]Days{
	`DAY_SUNDAY`:    1,
	`day_sunday`:    1,
	`DAY_MONDAY`:    3,
	`day_monday`:    3,
	`DAY_TUESDAY`:   5,
	`day_tuesday`:   5,
	`DAY_WEDNESDAY`: 7,
	`day_wednesday`: 7,
	`DAY_THURSDAY`:  9,
	`day_thursday`:  9,
	`DAY_FRIDAY`:    11,
	`day_friday`:    11,
	`DAY_SATURDAY`:  13,
	`day_saturday`:  13,
}

var _DaysDescMap = map[Days]string{
	1:  `Sunday is the first day of the week`,
	3:  `Monday is the second day of the week`,
	5:  `Tuesday is the third day of the week`,
	7:  `Wednesday is the fourth day of the week`,
	9:  `Thursday is the fifth day of the week`,
	11: `Friday is the sixth day of the week`,
	13: `Saturday is the seventh day of the week`,
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
	return errors.New(s + " is not a valid value for type Days")
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

// DaysValues returns all possible values
// for the type Days.
func DaysValues() []Days {
	return _DaysValues
}

// Values returns all possible values
// for the type Days.
func (i Days) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DaysValues))
	for i, d := range _DaysValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Days.
func (i Days) IsValid() bool {
	_, ok := _DaysMap[i]
	return ok
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

var _StatesMap = map[States]string{
	1:  `enabled`,
	3:  `not-enabled`,
	5:  `focused`,
	7:  `vered`,
	9:  `currently-being-pressed-by-user`,
	11: `actively-focused`,
	13: `selected`,
}

// BitIndexString returns the string
// representation of this States value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i States) BitIndexString() string {
	if str, ok := _StatesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatesNoOp() {
	var x [1]struct{}
	_ = x[Enabled-(1)]
	_ = x[Disabled-(3)]
	_ = x[Focused-(5)]
	_ = x[Hovered-(7)]
	_ = x[Active-(9)]
	_ = x[ActivelyFocused-(11)]
	_ = x[Selected-(13)]
}

var _StatesValues = []States{Enabled, Disabled, Focused, Hovered, Active, ActivelyFocused, Selected}

// StatesN is the highest valid value
// for type States, plus one.
const StatesN States = 14

var _StatesNameToValueMap = map[string]States{
	`enabled`:                         1,
	`not-enabled`:                     3,
	`focused`:                         5,
	`vered`:                           7,
	`currently-being-pressed-by-user`: 9,
	`actively-focused`:                11,
	`selected`:                        13,
}

var _StatesDescMap = map[States]string{
	1:  `Enabled indicates the widget is enabled`,
	3:  `Disabled indicates the widget is disabled`,
	5:  `Focused indicates the widget has keyboard focus`,
	7:  `Hovered indicates the widget is being hovered over`,
	9:  `Active indicates the widget is being interacted with`,
	11: `ActivelyFocused indicates the widget has active keyboard focus`,
	13: `Selected indicates the widget is selected`,
}

// SetString sets the States value from its
// string representation, and returns an
// error if the string is invalid.
func (i *States) SetString(s string) error {
	*i = 0
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _StatesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _StatesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type States")
		}
	}
	return nil
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

// StatesValues returns all possible values
// for the type States.
func StatesValues() []States {
	return _StatesValues
}

// Values returns all possible values
// for the type States.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type States.
func (i States) IsValid() bool {
	_, ok := _StatesMap[i]
	return ok
}

// String returns the string representation
// of this States value.
func (i States) String() string {
	str := ""
	for _, ie := range _StatesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i States) HasFlag(f enums.BitFlag) bool {
	return i&(1<<uint32(f.Int64())) != 0
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

// MarshalJSON implements the [json.Marshaler] interface.
func (i States) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *States) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("States should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
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

var _LanguagesMap = map[Languages]string{
	6:  `Go`,
	10: `Python`,
	14: `JavaScript`,
	18: `Dart`,
	22: `Rust`,
	26: `Ruby`,
	30: `C`,
	34: `CPP`,
	38: `ObjectiveC`,
	42: `Java`,
	46: `TypeScript`,
	50: `Kotlin`,
	54: `Swift`,
}

// BitIndexString returns the string
// representation of this Languages value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Languages) BitIndexString() string {
	if str, ok := _LanguagesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _LanguagesNoOp() {
	var x [1]struct{}
	_ = x[Go-(6)]
	_ = x[Python-(10)]
	_ = x[JavaScript-(14)]
	_ = x[Dart-(18)]
	_ = x[Rust-(22)]
	_ = x[Ruby-(26)]
	_ = x[C-(30)]
	_ = x[CPP-(34)]
	_ = x[ObjectiveC-(38)]
	_ = x[Java-(42)]
	_ = x[TypeScript-(46)]
	_ = x[Kotlin-(50)]
	_ = x[Swift-(54)]
}

var _LanguagesValues = []Languages{Go, Python, JavaScript, Dart, Rust, Ruby, C, CPP, ObjectiveC, Java, TypeScript, Kotlin, Swift}

// LanguagesN is the highest valid value
// for type Languages, plus one.
const LanguagesN Languages = 55

var _LanguagesNameToValueMap = map[string]Languages{
	`Go`:         6,
	`go`:         6,
	`Python`:     10,
	`python`:     10,
	`JavaScript`: 14,
	`javascript`: 14,
	`Dart`:       18,
	`dart`:       18,
	`Rust`:       22,
	`rust`:       22,
	`Ruby`:       26,
	`ruby`:       26,
	`C`:          30,
	`c`:          30,
	`CPP`:        34,
	`cpp`:        34,
	`ObjectiveC`: 38,
	`objectivec`: 38,
	`Java`:       42,
	`java`:       42,
	`TypeScript`: 46,
	`typescript`: 46,
	`Kotlin`:     50,
	`kotlin`:     50,
	`Swift`:      54,
	`swift`:      54,
}

var _LanguagesDescMap = map[Languages]string{
	6:  `Go is the best programming language`,
	10: ``,
	14: `JavaScript is the worst programming language`,
	18: ``,
	22: ``,
	26: ``,
	30: ``,
	34: ``,
	38: ``,
	42: ``,
	46: ``,
	50: ``,
	54: ``,
}

// SetString sets the Languages value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Languages) SetString(s string) error {
	*i = 0
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _LanguagesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _LanguagesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Languages")
		}
	}
	return nil
}

// Int64 returns the Languages value as an int64.
func (i Languages) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Languages value from an int64.
func (i *Languages) SetInt64(in int64) {
	*i = Languages(in)
}

// Desc returns the description of the Languages value.
func (i Languages) Desc() string {
	if str, ok := _LanguagesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// LanguagesValues returns all possible values
// for the type Languages.
func LanguagesValues() []Languages {
	return _LanguagesValues
}

// Values returns all possible values
// for the type Languages.
func (i Languages) Values() []enums.Enum {
	res := make([]enums.Enum, len(_LanguagesValues))
	for i, d := range _LanguagesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Languages.
func (i Languages) IsValid() bool {
	_, ok := _LanguagesMap[i]
	return ok
}

// String returns the string representation
// of this Languages value.
func (i Languages) String() string {
	str := ""
	for _, ie := range _LanguagesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Languages) HasFlag(f enums.BitFlag) bool {
	return i&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Languages) SetFlag(on bool, f ...enums.BitFlag) {
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

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Languages) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Languages) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Languages) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Languages) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("Languages should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}

var _MoreLanguagesMap = map[MoreLanguages]string{
	55: `Perl`,
}

// BitIndexString returns the string
// representation of this MoreLanguages value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i MoreLanguages) BitIndexString() string {
	if str, ok := _MoreLanguagesMap[i]; ok {
		return str
	}
	return Languages(i).BitIndexString()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _MoreLanguagesNoOp() {
	var x [1]struct{}
	_ = x[Perl-(55)]
}

var _MoreLanguagesValues = []MoreLanguages{Perl}

// MoreLanguagesN is the highest valid value
// for type MoreLanguages, plus one.
const MoreLanguagesN MoreLanguages = 56

var _MoreLanguagesNameToValueMap = map[string]MoreLanguages{
	`Perl`: 55,
	`perl`: 55,
}

var _MoreLanguagesDescMap = map[MoreLanguages]string{
	55: ``,
}

// SetString sets the MoreLanguages value from its
// string representation, and returns an
// error if the string is invalid.
func (i *MoreLanguages) SetString(s string) error {
	*i = 0
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _MoreLanguagesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _MoreLanguagesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			err := (*Languages)(i).SetString(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the MoreLanguages value as an int64.
func (i MoreLanguages) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the MoreLanguages value from an int64.
func (i *MoreLanguages) SetInt64(in int64) {
	*i = MoreLanguages(in)
}

// Desc returns the description of the MoreLanguages value.
func (i MoreLanguages) Desc() string {
	if str, ok := _MoreLanguagesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// MoreLanguagesValues returns all possible values
// for the type MoreLanguages.
func MoreLanguagesValues() []MoreLanguages {
	return _MoreLanguagesValues
}

// Values returns all possible values
// for the type MoreLanguages.
func (i MoreLanguages) Values() []enums.Enum {
	res := make([]enums.Enum, len(_MoreLanguagesValues))
	for i, d := range _MoreLanguagesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type MoreLanguages.
func (i MoreLanguages) IsValid() bool {
	_, ok := _MoreLanguagesMap[i]
	return ok
}

// String returns the string representation
// of this MoreLanguages value.
func (i MoreLanguages) String() string {
	str := ""
	for _, ie := range _MoreLanguagesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i MoreLanguages) HasFlag(f enums.BitFlag) bool {
	return i&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *MoreLanguages) SetFlag(on bool, f ...enums.BitFlag) {
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

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i MoreLanguages) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *MoreLanguages) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i MoreLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *MoreLanguages) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("MoreLanguages should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}
