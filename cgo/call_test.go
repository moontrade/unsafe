package cgo

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkCall(b *testing.B) {
	b.Run("Assembly Trampoline Call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NonBlockingNoop()
		}
	})
	b.Run("CGO Trampoline Call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BlockingNoop()
		}
	})
	b.Run("CGO Standard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CGONoop()
		}
	})
}

func TestSleep(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for i := 0; i < 10000; i++ {
		NonBlockingSleep(time.Second)
		println(time.Now().UnixNano())
	}
}
