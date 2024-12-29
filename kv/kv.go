/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package kv

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

const (
	_True = 1
)

type (
	Kind int
)

const (
	KindAny Kind = iota // Do not change
	KindBool
	KindDuration
	KindFloat
	KindInt
	KindInt64
	KindString
	KindTime
	KindUint
)

type (
	Value struct {
		_   [0]func() // disallow ==
		num uint64
		any any
	}

	KeyValue struct {
		Key   string
		Value Value
	}

	stringPtr *byte

	tLocation *time.Location
	tTime     time.Time
)

///////////////////////////////
/// Constructeurs: Value //////
///////////////////////////////

func AnyValue(value any) Value {
	switch value := value.(type) {
	case bool:
		return BoolValue(value)

	case float32:
		return FloatValue(float64(value))
	case float64:
		return FloatValue(value)

	case int:
		return IntValue(value)
	case int8:
		return IntValue(int(value))
	case int16:
		return IntValue(int(value))
	case int32:
		return IntValue(int(value))
	case int64:
		return Int64Value(value)

	case string:
		return StringValue(value)

	case time.Duration:
		return DurationValue(value)

	case time.Time:
		return TimeValue(value)

	case uint:
		return UintValue(uint64(value))
	case uint8:
		return UintValue(uint64(value))
	case uint16:
		return UintValue(uint64(value))
	case uint32:
		return UintValue(uint64(value))
	case uint64:
		return UintValue(value)

	default:
		return Value{any: value}
	}
}

func BoolValue(value bool) Value {
	v := uint64(0)
	if value {
		v = _True
	}
	return Value{num: v, any: KindBool}
}

func DurationValue(value time.Duration) Value {
	return Value{num: uint64(value.Nanoseconds()), any: KindDuration}
}

func FloatValue(value float64) Value {
	return Value{num: math.Float64bits(value), any: KindFloat}
}

func IntValue(value int) Value {
	return Value{num: uint64(value), any: KindInt}
}

func Int64Value(value int64) Value {
	return Value{num: uint64(value), any: KindInt64}
}

func StringValue(value string) Value {
	return Value{num: uint64(len(value)), any: stringPtr(unsafe.StringData(value))}
}

func TimeValue(value time.Time) Value {
	if value.IsZero() {
		return Value{any: tLocation(nil)}
	}

	nsec := value.UnixNano()

	if value.Equal(time.Unix(0, nsec)) {
		return Value{num: uint64(nsec), any: tLocation(value.Location())}
	}

	return Value{any: tTime(value.Round(0))}
}

func UintValue(value uint64) Value {
	return Value{num: value, any: KindUint}
}

///////////////////////////////
/// Constructeurs: KeyValue ///
///////////////////////////////

func Any(key string, value any) KeyValue {
	return KeyValue{key, AnyValue(value)}
}

func Bool(key string, value bool) KeyValue {
	return KeyValue{key, BoolValue(value)}
}

func Duration(key string, value time.Duration) KeyValue {
	return KeyValue{key, DurationValue(value)}
}

func Float(key string, value float64) KeyValue {
	return KeyValue{key, FloatValue(value)}
}

func Int(key string, value int) KeyValue {
	return KeyValue{key, IntValue(value)}
}

func Int64(key string, value int64) KeyValue {
	return KeyValue{key, Int64Value(value)}
}

func String(key string, value string) KeyValue {
	return KeyValue{key, StringValue(value)}
}

func Time(key string, value time.Time) KeyValue {
	return KeyValue{key, TimeValue(value)}
}

func Uint(key string, value uint64) KeyValue {
	return KeyValue{key, UintValue(value)}
}

func Error(err error) KeyValue {
	return KeyValue{"error", StringValue(err.Error())}
}

///////////////////////////////
/// Accesseurs: Value /////////
///////////////////////////////

func (v Value) Kind() Kind {
	switch t := v.any.(type) {
	case Kind:
		return t

	case stringPtr:
		return KindString

	case tLocation, tTime:
		return KindTime

	default:
		return KindAny
	}
}

func (v Value) Bool() bool {
	return v.num == _True
}

func (v Value) Duration() time.Duration {
	return time.Duration(int64(v.num))
}

func (v Value) Float() float64 {
	return math.Float64frombits(v.num)
}

func (v Value) Int() int {
	return int(v.num)
}

func (v Value) Int64() int64 {
	return int64(v.num)
}

func (v Value) String() string {
	return unsafe.String(v.any.(stringPtr), v.num)
}

func (v Value) Time() time.Time {
	switch a := v.any.(type) {
	case tLocation:
		if a == nil {
			return time.Time{}
		}

		return time.Unix(0, int64(v.num)).In(a)

	case tTime:
		return time.Time(a)

	default:
		panic(fmt.Sprintf("bad time type: %T", v.any)) /////////////////////////////////////////////////////////////////
	}
}

func (v Value) Uint() uint64 {
	return v.num
}

/*
####### END ############################################################################################################
*/
