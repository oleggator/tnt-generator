
#include <msgpuck.h>
#include <stdint.h>
#include <stdio.h>

#include "models.h"

int encode_with_arr(with_arr_t *with_arr, char *buf, char *buf_end) {
  int err = 0;
  char *end;
  end = mp_encode_array(buf, 1);
  // field arr

  end = mp_encode_array(buf, with_arr->arr_len);
  for (uint32_t i = 0; i < with_arr->arr_len; ++i) {
    end = mp_encode_int(end, with_arr->arr[i]);
  }
}

int decode_with_arr(with_arr_t *with_arr, const char **buf,
                    const char **buf_end) {
  int err = 0;

  uint32_t field_count = mp_decode_array(buf);
  if (field_count != 1) {
    goto wrong_field_count_error;
  }
  // field arr

  with_arr->arr_len = mp_decode_array(buf);
  for (uint32_t i = 0; i < with_arr->arr_len; ++i) {
    with_arr->arr[i] = mp_decode_int(buf);
  }

  return 0;

wrong_field_count_error:
  /*	say_error("wrong '%s' fields count - %d, must be %d", ".Name",
   * field_count, 1);*/
  return 1;
too_big_array:
  return 2;
}

int encode_hello_request(hello_request_t *hello_request, char *buf,
                         char *buf_end) {
  int err = 0;
  char *end;
  end = mp_encode_array(buf, 2);
  // field with_arr

  err = encode_with_arr(&hello_request->with_arr, buf, buf_end);
  if (err != 0) {
    return err;
  };

  // field greeting
  end =
      mp_encode_str(end, hello_request->greeting, hello_request->greeting_len);
}

int decode_hello_request(hello_request_t *hello_request, const char **buf,
                         const char **buf_end) {
  int err = 0;

  uint32_t field_count = mp_decode_array(buf);
  if (field_count != 2) {
    goto wrong_field_count_error;
  }
  // field with_arr

  err = decode_with_arr(&hello_request->with_arr, buf, buf_end);
  if (err != 0) {
    return err;
  };

  // field greeting
  hello_request->greeting = mp_decode_str(buf, &hello_request->greeting_len);

  return 0;

wrong_field_count_error:
  /*	say_error("wrong '%s' fields count - %d, must be %d", ".Name",
   * field_count, 2);*/
  return 1;
too_big_array:
  return 2;
}

int encode_hello_response(hello_response_t *hello_response, char *buf,
                          char *buf_end) {
  int err = 0;
  char *end;
  end = mp_encode_array(buf, 1);
  // field reply
  end = mp_encode_str(end, hello_response->reply, hello_response->reply_len);
}

int decode_hello_response(hello_response_t *hello_response, const char **buf,
                          const char **buf_end) {
  int err = 0;

  uint32_t field_count = mp_decode_array(buf);
  if (field_count != 1) {
    goto wrong_field_count_error;
  }
  // field reply
  hello_response->reply = mp_decode_str(buf, &hello_response->reply_len);

  return 0;

wrong_field_count_error:
  /*	say_error("wrong '%s' fields count - %d, must be %d", ".Name",
   * field_count, 1);*/
  return 1;
too_big_array:
  return 2;
}
