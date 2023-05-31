package cgo

import (
	"github.com/moontrade/unsafe/cgo/internal/trampoline"
	"time"
)

// BlockingSleep invokes C++ std::thread sleep utilizing the safe CGO trampoline.
// It is preferable to use time.Sleep in likely all circumstances.
func BlockingSleep(nanos time.Duration) {
	Blocking((*byte)(trampoline.Sleep), uintptr(nanos), 0)
}

// NonBlockingSleep invokes C++ std::thread sleep using Assembly Trampoline.
// It is only acceptable to do this  after runtime.LockOSThread() and only
// used when absolutely certain.
//
// WARNING!!! You should probably never use this
func NonBlockingSleep(nanos time.Duration) {
	NonBlocking((*byte)(trampoline.Sleep), uintptr(nanos), 0)
}

// NonBlockingNoop can be used for benchmarking Assembly Trampoline/CGO trampoline overhead.
// The C function is empty and returns void.
func NonBlockingNoop() {
	NonBlocking((*byte)(trampoline.Stub), 0, 0)
}

// BlockingNoop can be used for benchmarking CGO Trampoline overhead.
// The C function is empty and returns void.
func BlockingNoop() {
	Blocking((*byte)(trampoline.Stub), 0, 0)
}

// CGONoop can be used for benchmarking CGO overhead.
// The C function is empty and returns void.
func CGONoop() {
	trampoline.CGO()
}
