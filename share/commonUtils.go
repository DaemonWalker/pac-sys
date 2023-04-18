package share

import (
	"encoding/json"
	"io"
	"pac-sys/models"
)

func CreatePanic(code int, message string) {
	info := models.PanicInfo{Code: code, Message: message}
	panic(info)
}

func ErrorPanic(err error) {
	CreatePanic(500, err.Error())
}

func StatusPanic(code int) {
	info := models.PanicInfo{Code: code}
	panic(info)
}

func JSONDecode(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return nil
}

func ConvertArray[TS any, TT any](ts []TS, foo func(ts TS) TT) []TT {
	array := make([]TT, len(ts))
	for i, v := range ts {
		array[i] = foo(v)
	}
	return array
}

func ArrayContains[T comparable](ts []T, t T) bool {
	for _, v := range ts {
		if t == v {
			return true
		}
	}
	return false
}
