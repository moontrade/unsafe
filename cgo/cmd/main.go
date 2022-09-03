package main

import (
	"github.com/moontrade/unsafe/cgo"
)

func main() {
	//cgo.CGO()
	cgo.NonBlocking((*byte)(nil), 0, 0)
}
