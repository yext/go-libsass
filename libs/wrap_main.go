// +build !dev

package libs

// #cgo CFLAGS: -O2 -fPIC -Iinclude
// #cgo CPPFLAGS: -Iinclude
// #cgo CXXFLAGS: -g -std=c++0x -O2 -fPIC
// #cgo LDFLAGS: -lstdc++ -lm
// #cgo darwin linux LDFLAGS: -ldl
//
import "C"
