package ut

import (
	"testing"

	"github.com/pkg/errors"
)

func TestLogger(t *testing.T) {
	Logger("this is sample error log")

	LoggerError(errors.New("this is error log"))

	Loggerln(
		"message1",
		"message2",
		33333,
		map[string]interface{}{
			"name": "zhanagsan",
		},
		struct {
		}{})

	LoggerDebug(errors.New("this is debug log"))

}

func TestLoggerln(t *testing.T) {
	type args struct {
		str []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Loggerln(tt.args.str...)
		})
	}
}

func TestLoggerError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerError(tt.args.err)
		})
	}
}

func TestLoggerDebug(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerDebug(tt.args.err)
		})
	}
}
