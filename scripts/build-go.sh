#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "Building Haskell foreign library (magic-ffi)..."
cabal build magic-ffi

echo "Syncing generated artifacts into csrc/ for cgo..."
LIB_PATH=$(find "$ROOT_DIR/dist-newstyle" -path '*f/magic-ffi/*' -name 'libmagic-ffi.dylib' | head -n1)
STUB_PATH=$(find "$ROOT_DIR/dist-newstyle" -path '*f/magic-ffi/*' -name 'Magic_stub.h' | head -n1)

if [[ -z "${LIB_PATH}" || -z "${STUB_PATH}" ]]; then
  echo "Failed to locate libmagic-ffi.dylib or Magic_stub.h under dist-newstyle. Did cabal build succeed?" >&2
  exit 1
fi

cp "${LIB_PATH}" "$ROOT_DIR/csrc/libmagic-ffi.dylib"
cp "${STUB_PATH}" "$ROOT_DIR/csrc/Magic_stub.h"

echo "Locating GHC runtime libraries for rpath..."
GHC_LIBDIR="$(ghc --print-libdir)"
RPATH_DIR=$(find "${GHC_LIBDIR}" -name 'libHSrts-ghc*.dylib' -print -quit)
if [[ -z "${RPATH_DIR}" ]]; then
  echo "Could not find libHSrts under ${GHC_LIBDIR}" >&2
  exit 1
fi
RPATH_DIR="$(dirname "${RPATH_DIR}")"

export CGO_CFLAGS="-I${ROOT_DIR}/csrc"
export CGO_LDFLAGS="-L${ROOT_DIR}/csrc -lmagic-ffi -Wl,-rpath,${ROOT_DIR}/csrc -Wl,-rpath,${RPATH_DIR}"

if ! command -v go >/dev/null 2>&1; then
  echo "Go toolchain not found on PATH; install Go (>=1.21) to build the sample." >&2
  exit 1
fi

echo "Building Go sample..."
(
  cd "$ROOT_DIR/csrc"
  go build -o hs-from-go .
)

cat <<EOF
Build complete.

Artifacts placed in csrc/:
  - libmagic-ffi.dylib (Haskell foreign library)
  - Magic_stub.h (header for exported functions)
  - hs-from-go (Go binary that calls add_hs)

To run:
  DYLD_LIBRARY_PATH="${ROOT_DIR}/csrc:${RPATH_DIR}" ./csrc/hs-from-go
EOF
