package utils

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
