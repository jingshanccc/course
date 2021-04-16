package util

import (
	"fmt"
	"reflect"
)

// Contains: 检查数组是否包含指定元素 是返回元素下标 否返回-1
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	if array == nil {
		return
	}
	fmt.Printf("type:%v\n", reflect.TypeOf(array).Kind())
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}

		}
	case reflect.Array:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}

		}
	}

	return
}
