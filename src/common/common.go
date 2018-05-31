package common

import "reflect"

func SliceEqual(a, b interface{}) bool {
	voa, vob := reflect.ValueOf(a), reflect.ValueOf(b)
	if voa.Kind() != reflect.Slice || vob.Kind() != reflect.Slice {
		panic("Param $1 and $2 must be a slice")
	}

	if voa.Len() != vob.Len() {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < voa.Len(); i++ {
		if voa.Index(i).Interface() != vob.Index(i).Interface() {
			return false
		}
	}
	return true
}
