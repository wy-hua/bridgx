package cmp

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/galaxy-future/BridgX/pkg/utils"
	jsoniter "github.com/json-iterator/go"
)

const (
	columnTag           = "diff"
	ignoreFieldTagValue = "-"
)

var (
	errDifferentTypeCompare = errors.New("different types cannot be compared")
	errInvalidType          = errors.New("only support struct")
)

type DiffResult struct {
	Fields []FieldInfo
}

type FieldInfo struct {
	Key      string
	ValueOld string `json:"old"`
	ValueNew string `json:"new"`
}

func toString(v reflect.Value) string {
	if isZeroValue(v) {
		return ""
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if isZeroValue(v) {
		return ""
	}
	if v.Type().Name() == "Time" {
		t, ok := v.Interface().(time.Time)
		if ok {
			return utils.FormatTime(t)
		}
	}
	return fmt.Sprintf("%v", v)
}

// Diff will compare the `old` and `new` one then return diff result.Call DiffResult.Display function
// will return display map with new filed json value.
func Diff(o, n interface{}) (res DiffResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
			return
		}
	}()

	t1, t2 := reflect.TypeOf(o), reflect.TypeOf(n)
	v1, v2 := reflect.ValueOf(o), reflect.ValueOf(n)
	if t1 == nil && t2 == nil {
		return DiffResult{}, nil
	} else if t1 == nil {
		return diffWithOneNil(t2, v2, true)
	} else if t2 == nil {
		return diffWithOneNil(t1, v1, false)
	}

	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
		v1 = v1.Elem()
	}
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
		v2 = v2.Elem()
	}
	if t1 != t2 {
		return DiffResult{}, errDifferentTypeCompare
	}
	return diff(t1, v1, v2)
}

func diff(t reflect.Type, v1, v2 reflect.Value) (res DiffResult, err error) {
	if v1.Kind() == reflect.Struct {
		for i := 0; i < v1.NumField(); i++ {
			tag := getDiffTag(t.Field(i))
			if tag == "" {
				continue
			}
			f1Value := v1.Field(i)
			f2Value := v2.Field(i)
			if compare(f1Value, f2Value) {
				continue
			}

			res.Fields = append(res.Fields, FieldInfo{
				Key:      tag,
				ValueOld: toString(f1Value),
				ValueNew: toString(f2Value),
			})
		}
	} else if v1.Kind() == reflect.Map {

	}
	return res, nil
}

func diffWithOneNil(t reflect.Type, v reflect.Value, oldIsNil bool) (res DiffResult, err error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			tag := getDiffTag(t.Field(i))
			if tag == "" {
				continue
			}
			fValue := v.Field(i)
			if fValue.IsZero() {
				continue
			}

			fieldInfo := FieldInfo{
				Key: tag,
			}
			if oldIsNil {
				fieldInfo.ValueNew = toString(fValue)
			} else {
				fieldInfo.ValueOld = toString(fValue)
			}
			res.Fields = append(res.Fields, fieldInfo)
		}
	} else if v.Kind() == reflect.Map {

	}
	return res, nil
}

func compare(v1, v2 reflect.Value) bool {
	k1, k2 := v1.Kind(), v2.Kind()
	if k1 != k2 {
		return false
	}
	switch k1 {
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Slice, reflect.Array, reflect.Map:
		b1, _ := jsoniter.Marshal(v1.Interface())
		b2, _ := jsoniter.Marshal(v2.Interface())
		return reflect.DeepEqual(b1, b2)
	case reflect.Ptr, reflect.Interface:
		return reflect.DeepEqual(v1.Elem(), v2.Elem())
	case reflect.Func, reflect.Struct, reflect.Chan, reflect.UnsafePointer:
		return true
	default:
		return true
	}
}

func isZeroValue(v reflect.Value) bool {
	return v == reflect.Value{}
}

func getDiffTag(field reflect.StructField) string {
	tag := field.Tag.Get(columnTag)
	if tag == ignoreFieldTagValue || !field.IsExported() {
		return ""
	}
	if tag == "" {
		tag = field.Name
	}
	return tag
}
