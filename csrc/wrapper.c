#include <HsFFI.h>
#include <stdio.h>
#include "wrapper.h"

// Initialize the Haskell RTS so exported functions can run.
HsBool magic_h_init(void) {
    int argc = 1;
    char *argv[] = { "magic", NULL };
    hs_init(&argc, (char ***)&argv);
    return HS_BOOL_TRUE;
}

// Shutdown the RTS when no more Haskell calls are needed.
void magic_h_exit(void) {
    hs_exit();
}
