package libsass

import "github.com/yext/go-libsass/libs"

// Error takes a Go error and returns a libsass Error
func Error(err error) SassValue {
	return SassValue{value: libs.MakeError(err.Error())}
}

// Warn takes a string and causes a warning in libsass
func Warn(s string) SassValue {
	return SassValue{value: libs.MakeWarning(s)}
}
