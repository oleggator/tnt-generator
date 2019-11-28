#include <stdio.h>

#include "models.h"

#define INPUT_BUFFER_LEN 2048
#define OUTPUT_BUFFER_LEN 2048

int main() {
    char input_buffer[INPUT_BUFFER_LEN];

    FILE *file = fopen("msgpack", "r");
    if (file == 0) {
        return 1;
    }
    fseek(file, 0, SEEK_END);
    long fsize = ftell(file);
    rewind(file);
    fread(input_buffer, 1, fsize, file);
    fclose(file);

    const char *input_buffer_begin = input_buffer;
    const char *input_buffer_end = input_buffer_begin + fsize;

    HelloRequest hello_request;
    decodeHelloRequest(&hello_request, &input_buffer_begin, &input_buffer_end);


    for (uint32_t i = 0; i < hello_request.with_arr.arr_len; ++i) {
        printf("%d\n", hello_request.with_arr.arr[i]);
    }
    printf("%.*s\n", hello_request.greeting_len, hello_request.greeting);


    HelloResponse hello_response = {
        .reply = "123467",
    };

    char output_buffer[OUTPUT_BUFFER_LEN];
    char *output_buffer_begin = output_buffer;
    char *output_buffer_end = output_buffer_begin + OUTPUT_BUFFER_LEN;

    encodeHelloResponse(&hello_response, &output_buffer_begin, &output_buffer_end);

    return 0;
}
