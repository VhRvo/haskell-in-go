# Haskell from Go

Small example of calling a `foreign export` Haskell function from Go via cgo.

## Layout
- `src/Magic.hs`: Haskell code exporting `add_hs` (adds two `CInt` values).
- `csrc/wrapper.c` / `csrc/wrapper.h`: helpers that start/stop the Haskell RTS.
- `csrc/main.go`: Go program that calls `add_hs`.
- `scripts/build-go.sh`: builds the shared library with Cabal, copies the artefacts into `csrc/`, and builds the Go binary.

## Prerequisites
- GHC + Cabal (same toolchain used to build the Haskell library)
- Go toolchain (tested with Go 1.21)

## Build and run
```sh
# From repo root
./scripts/build-go.sh

# Then run with the library search path printed by the script, e.g.
DYLD_LIBRARY_PATH="/path/to/for_go/csrc:/path/to/ghc/lib/ghc-*/lib/aarch64-osx-ghc-*" ./csrc/hs-from-go
```

The build script locates `libmagic-ffi.dylib` and `Magic_stub.h` under `dist-newstyle`, copies them into `csrc/`, sets the required `CGO_*` flags (including `-rpath` to the GHC libs), and builds `./csrc/hs-from-go`.
