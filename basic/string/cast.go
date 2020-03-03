package string

import (
	"reflect"
	"unsafe"
)

/*
	string 与 []byte 结构十分相似,仅仅少了 cap 字段
	直接使用 unsafe 指针进行赋值
*/

// StringToBytes converts a string to bytes without copy.
func StringToBytes(s string) (b []byte) {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	h.Len = len(s)
	h.Cap = len(s)
	return
}

//
func StringToBytes2(s string) (b []byte) {
	return []byte(s)
}

// BytesToString converts a byte array to string without copy.
func BytesToString(b []byte) (s string) {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	h.Data = (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	h.Len = len(b)
	return
}

func BytesToString2(b []byte) (s string) {
	return string(b)
}
