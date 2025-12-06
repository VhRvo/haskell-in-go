package main

/*
// #include "Magic_stub.h"

#cgo CFLAGS: -I/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/rts-1.0.2/include

// static linking?
#cgo LDFLAGS: wrapper.o -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/rts-1.0.2 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/base-4.17.2.1 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/ghc-prim-0.9.1 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/ghc-bignum-1.3 -lHSrts-1.0.2 -lHSbase-4.17.2.1 -lHSghc-prim-0.9.1 -lHSghc-bignum-1.3 -lffi -liconv -lmagic-ffi
#include "wrapper.c"
*/

import "C"

import "fmt"

func init() {
    // 初始化 Haskell RTS（整个程序只需要一次）
    C.magic_h_init()
}

func main() {
    defer C.magic_h_exit()

    a := C.HsInt32(10)
    b := C.HsInt32(32)

    // 这里直接调用 Haskell 导出的 add_hs
    // res := C.hsAdd(a, b)

    // fmt.Println("10 + 32 =", int32(res))
	fmt.Println(a, b)
}
