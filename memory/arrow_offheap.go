package memory

import (
	"reflect"
	"unsafe"
)

var ArrowAllocator arrowAllocator

type arrowAllocator struct{}

func (arrowAllocator) Allocate(size int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(Alloc(uintptr(size))),
		Len:  size,
		Cap:  size,
	}))
}

func (arrowAllocator) Reallocate(size int, b []byte) []byte {
	if len(b) < 1 {
		if size < 1 {
			return nil
		}
		return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
			Data: uintptr(Alloc(uintptr(size))),
			Len:  size,
			Cap:  size,
		}))
	}
	newAlloc := Realloc(Pointer(unsafe.Pointer(&b[0])), uintptr(size))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(newAlloc),
		Len:  size,
		Cap:  size,
	}))
}

func (arrowAllocator) Free(b []byte) {
	if cap(b) == 0 {
		return
	}
	Free(Pointer(unsafe.Pointer(&b[0])))
}
