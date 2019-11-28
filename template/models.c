#include <stdint.h>
#include <msgpuck.h>
#include <stdio.h>

#include "models.h"

int decodeWithArr(WithArr *with_arr, const char **args, const char **args_end) {
    uint32_t field_count = mp_decode_array(args);
    if (field_count != 1) {
        goto wrong_field_count_error;
    }

    with_arr->arr_len = mp_decode_array(args);
    if (field_count > ARRAY_LEN) {
        goto too_big_array;
    }

    for (uint32_t i = 0; i < with_arr->arr_len; ++i) {
        with_arr->arr[i] = mp_decode_int(args);
    }

    return 0;

wrong_field_count_error:
    return 1;
too_big_array:
    return 2;
}

int decodeHelloRequest(HelloRequest *hello_request, const char **args, const char **args_end) {
    uint32_t field_count = mp_decode_array(args);

    int err = decodeWithArr(&hello_request->with_arr, args, args_end);
    if (err != 0) {
        return err;
    }

    hello_request->greeting = mp_decode_str(args, &hello_request->greeting_len);

    return 0;
}

int encodeHelloResponse(HelloResponse *hello_response, char **buf, char **buf_end) {
    char **body_begin = buf;

    char *body_end = mp_encode_str(*body_begin, hello_response->reply, strlen(hello_response->reply));

    for (char *i = *body_begin; i < body_end; i++) {
        unsigned char ch = *i;
        printf("%X", ch);
        if (i < body_end - 1) {
            printf("%s", " ");
        } else {
            printf("%s", "\n");
        }
    }

    return 0;
}
