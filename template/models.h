#pragma once

#include <stdint.h>

#define ARRAY_LEN 1024

typedef struct {
    int arr[ARRAY_LEN];
    uint32_t arr_len;
} WithArr;
int decodeWithArr(WithArr *with_arr, const char **args, const char **args_end);


typedef struct {
    WithArr with_arr;
    const char *greeting;
    uint32_t greeting_len;
} HelloRequest;
int decodeHelloRequest(HelloRequest *hello_request, const char **args, const char **args_end);


typedef struct {
    char *reply;
} HelloResponse;
int encodeHelloResponse(HelloResponse *, char **buf, char **buf_end);
