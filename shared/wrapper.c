#include "HsFFI.h"
#include <stdio.h>
#include "Magic_stub.h"

HsBool magic_h_init(void) {
    int argc = 1;
    char *argv[] = { "magic", NULL };
    hs_init(&argc, (char***)&argv);
    return HS_BOOL_TRUE;
}

void magic_h_exit(void) {
    hs_exit();
}
