package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lmagic-ffi -Wl,-rpath,${SRCDIR}
#include "Magic_stub.h"
#include "wrapper.h"

*/

import "C"

import "fmt"

func init() {
	// Initialize the Haskell RTS once before calling into Haskell.
	if C.magic_h_init() == 0 {
		panic("failed to start Haskell RTS")
	}
}

func main() {
	defer C.magic_h_exit()

	a := C.HsInt32(10)
	b := C.HsInt32(32)

	// Invoke the exported Haskell function.
	res := C.add_hs(a, b)
	fmt.Printf("add_hs(%d, %d) = %d\n", int32(a), int32(b), int32(res))
}
