#include <tarantool/module.h>

#pragma once

/*
    1. Decode arguments to structs
    2. Pass structs to funcs
    3. Encode func return values
*/
int SayHelloWrapper(box_function_ctx_t *ctx, const char *args, const char *args_end);
