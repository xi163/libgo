package conv

import (
	"reflect"
	"strconv"
	"unsafe"
)

/*
*
struct转换成byte
*/
func StructToByte(v any, len int) []byte {
	var x reflect.SliceHeader
	x.Len = len
	x.Cap = len
	x.Data = reflect.ValueOf(v).Pointer()
	return *(*[]byte)(unsafe.Pointer(&x))
}

/*
*
byte转换成struct
*/
func ByteToStruct(b []byte) unsafe.Pointer {
	return unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	)
}

func ByteToStr(b []byte) string {
	return string(b)
}

/*
*
string 转换
*/
func StrToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StrToInt32(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return i
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StrToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func StrToByte(s string) []byte {
	return []byte(s)
}

/*
*
int 转换
*/
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func IntToInt32(i int) int32 {
	return int32(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

/*
*
int32 转换
*/
func Int32ToInt(i int32) int {
	return int(i)
}

func Int32ToInt64(i int32) int64 {
	return int64(i)
}

/*
*
int64 转换
*/
func Int64ToInt(i int64) int {
	return int(i)
}

func Int64ToInt32(i int64) int32 {
	return int32(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
