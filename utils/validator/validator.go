package validator

import "reflect"

func Empty(val any) bool {
	switch val {
	case nil:
		return true
	default:
		v := reflect.ValueOf(val)
		switch v.Kind() {
		case reflect.String, reflect.Array:
			return v.Len() == 0
		case reflect.Map, reflect.Slice:
			return v.Len() == 0 || v.IsNil()
		case reflect.Bool:
			return !v.Bool()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return v.Int() == 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return v.Uint() == 0
		case reflect.Float32, reflect.Float64:
			return v.Float() == 0
		case reflect.Interface, reflect.Ptr:
			return v.IsNil()
		}
		return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
	}
}
