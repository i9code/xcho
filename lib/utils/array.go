package utils

import (
	"reflect"
)

// IsInArray 通用的判断数据是否在数组中
func IsInArray(val interface{}, array interface{}) (exist bool, index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exist = true

				return
			}
		}
	}

	return
}
