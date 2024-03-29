package trampoline

/*
#cgo CXXFLAGS: -std=c++11 -I./
#cgo LDFLAGS: -lstdc++
#include "trampoline.h"
*/
import "C"
import "unsafe"

var (
	Stub  = C.moontrade_stub
	Sleep = C.moontrade_sleep // moontrade_sleep(u64 nanoseconds)
)

func CGO() {
	C.moontrade_stub()
}

func NonBlocking(fn *byte, arg0, arg1 uintptr) {
	Blocking(fn, arg0, arg1)
}

func Blocking(fn *byte, arg0, arg1 uintptr) {
	C.moontrade_trampoline((C.size_t)(uintptr(unsafe.Pointer(fn))), (C.size_t)(arg0), (C.size_t)(arg1))
}
