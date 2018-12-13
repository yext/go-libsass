package libsass

import "github.com/yext/go-libsass/libs"

// Version reports libsass version information
func Version() string {
	return libs.Version()
}
