#include <msgpuck.h>

#include "wrappers.h"
#include "models.h"
#include "funcs.h"

const int BUFFER_LEN = 1024;

//void send_response(const HelloResponse *hello_response) {
//    int err;
//
//    /*
//        Encode response
//    */
//    char buffer[BUFFER_LEN];
//    char *buffer_end = buffer + BUFFER_LEN;
//
//    err = encodeHelloResponse(&hello_response, buffer, buffer_end);
//    if (err != 0)
//    {
//        return err;
//    }
//
//    box_tuple_format_t *tuple_format = box_tuple_format_default();
//    box_tuple_t *tuple = box_tuple_new(tuple_format, buffer, buffer_end);
//    box_return_tuple(ctx, tuple);
//}

int SayHelloWrapper(box_function_ctx_t *ctx, const char *args, const char *args_end) {
    int err;

    /*
        Decode request
    */
    HelloRequest hello_request;
    err = decodeHelloRequest(&hello_request, &args, &args_end);
    if (err != 0) {
        return err;
    }

    /*
        Execute function
    */
    HelloResponse hello_response;
    err = SayHello(&hello_request, &hello_response);
    if (err != 0) {
        return err;
    }

    return 0;
}

