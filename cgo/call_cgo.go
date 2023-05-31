//go:build !amd64 && !arm64 && !tinygo

package cgo

import "github.com/moontrade/unsafe/cgo/internal/trampoline"

func NonBlocking(fn *byte, arg0, arg1 uintptr) {
	trampoline.Blocking(fn, arg0, arg1)
}

func Blocking(fn *byte, arg0, arg1 uintptr) {
	trampoline.Blocking(fn, arg0, arg1)
}
