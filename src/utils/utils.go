package utils

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

// LogFormat LogFormat
type LogFormat struct{}

func (logFormat *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer

	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	buffer.WriteString("[" + entry.Level.String()[0:4] + "] ")
	buffer.WriteString(entry.Time.Format("2006/01/02 15:04:05.000 "))

	for key, value := range entry.Data {
		buffer.WriteByte('[')
		buffer.WriteString(key)
		buffer.WriteByte(':')
		fmt.Fprint(buffer, value)
		buffer.WriteString("] ")
	}

	buffer.WriteString(entry.Message)
	buffer.WriteByte('\n')

	return buffer.Bytes(), nil
}

func HasData[T any](slice []T) bool {
	LenOfSlice := len(slice)
	if LenOfSlice > 0 {
		return true
	} else {
		return false
	}
}
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

type HandleDiff[T comparable] func(item1 T, item2 T) bool

func HandleDiffDefault[T comparable](val1 T, val2 T) bool {
	return val1 == val2
}

func Diff[T comparable](items1 []T, items2 []T, callback HandleDiff[T]) []T {
	acc := []T{}
	for _, item1 := range items1 {
		find := false
		for _, item2 := range items2 {
			if callback(item1, item2) {
				find = true
				break
			}
		}
		if !find {
			acc = append(acc, item1)
		}
	}
	return acc
}
