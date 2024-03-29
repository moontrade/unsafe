//go:build !libfuzzer && !tinygo && (amd64 || arm64) && (linux || darwin)

package cgo

import "github.com/moontrade/unsafe/cgo/internal/trampoline"

// NonBlocking C function fn without going all the way through cgo
// Be very careful using it. If the C code blocks it can/will
// lock up your app.
// Example: NonBlocking((*byte)(C.my_c_func), 0, 0)
//
//	void my_c_func(size_t arg0, size_t arg1) {
//	}
//
//go:nosplit
//go:noescape
func NonBlocking(fn *byte, arg0, arg1 uintptr)

func Blocking(fn *byte, arg0, arg1 uintptr) {
	trampoline.Blocking(fn, arg0, arg1)
}
