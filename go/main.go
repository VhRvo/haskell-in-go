package main

/*
export DYLD_LIBRARY_PATH=/Users/zhangyongzhuo/.ghcup/ghc/9.6.7/lib/ghc-9.6.7/lib/aarch64-osx-ghc-9.6.7:/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/shared:$DYLD_LIBRARY_PATH
*/

/*
#include "Magic_stub.h"
#include <stdio.h>

// use dynamic linking
#cgo CFLAGS: -I/Users/zhangyongzhuo/.ghcup/ghc/9.6.7/lib/ghc-9.6.7/lib/aarch64-osx-ghc-9.6.7/rts-1.0.2/include -I/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/csrc
#cgo LDFLAGS: -L/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/shared -L/Users/zhangyongzhuo/.ghcup/ghc/9.6.7/lib/ghc-9.6.7/lib/aarch64-osx-ghc-9.6.7/ -lHSrts-1.0.2-ghc9.6.7 -lHSbase-4.18.3.0-ghc9.6.7 -lHSghc-prim-0.10.0-ghc9.6.7 -lHSghc-bignum-1.3-ghc9.6.7 -lmagic-ffi -lffi -liconv -lm -ldl

// 9.2.8
// #cgo CFLAGS: -I/Users/zhangyongzhuo/.ghcup/ghc/9.2.8/include -I/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/csrc
// #cgo LDFLAGS: -L/Users/zhangyongzhuo/.ghcup/ghc/9.2.8/lib/ghc-9.2.8/lib/aarch64-osx-ghc-9.2.8/rts-1.0.2 -L/Users/zhangyongzhuo/.ghcup/ghc/9.2.8/lib/ghc-9.2.8/lib/aarch64-osx-ghc-9.2.8/base-4.16.4.0 -L/Users/zhangyongzhuo/.ghcup/ghc/9.2.8/lib/ghc-9.2.8/lib/aarch64-osx-ghc-9.2.8/ghc-prim-0.8.0 -L/Users/zhangyongzhuo/.ghcup/ghc/9.2.8/lib/ghc-9.2.8/lib/aarch64-osx-ghc-9.2.8/ghc-bignum-1.2 -L/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/shared -lmagic-ffi -lHSrts-1.0.2_thr -lHSbase-4.16.4.0 -lHSghc-prim-0.8.0 -lHSghc-bignum-1.2 -lffi -liconv

// 9.4.8
// #cgo CFLAGS: -I/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/rts-1.0.2/include -I/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/csrc
// #cgo LDFLAGS: -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/rts-1.0.2 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/base-4.17.2.1 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/ghc-prim-0.9.1 -L/Users/zhangyongzhuo/.ghcup/ghc/9.4.8/lib/ghc-9.4.8/lib/aarch64-osx-ghc-9.4.8/ghc-bignum-1.3 -L/Users/zhangyongzhuo/Desktop/Workplace/Languages/Haskell/haskell-in-go/shared -lHSrts-1.0.2 -lHSbase-4.17.2.1 -lHSghc-prim-0.9.1 -lHSghc-bignum-1.3 -lffi -liconv -lm -ldl -lmagic-ffi

static void my_init_hs() {
	int argc = 1;
	char *argv[] = { "program_name", "+RTS", "-A128m", "-H512m", "-RTS", NULL };
	// char *argv[] = { "magic", NULL };
	char **p_argv = argv;
	char **argv_array[] = { p_argv, NULL };
	printf("Initializing Haskell runtime...\n");
	hs_init(&argc, argv_array);
	// hs_init(&argc, NULL);
	// hs_init_with_rtsopts(NULL, NULL);
	// hs_init_ghc(NULL, NULL, (RtsConfig)NULL);
	printf("End Haskell runtime...\n");
}
*/
import "C"

import "fmt"

func init() {
	// 初始化 Haskell RTS（整个程序只需要一次）
}

func main() {
	C.my_init_hs()

	// defer C.magic_h_exit()
	defer C.hs_exit()

	a, b := 2, 3
	fmt.Println(a, b)

	a1 := C.HsInt32(10)
	b1 := C.HsInt32(32)

	// 这里直接调用 Haskell 导出的 add_hs
	res := C.add_hs(a1, b1)
	fmt.Println("10 + 32 =", int32(res))
}
