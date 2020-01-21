
#pragma once

#include <stdint.h>
#include <tarantool/module.h>

#define ARRAY_LEN 1024

typedef struct {
  int64_t arr[ARRAY_LEN];
  uint32_t arr_len;
} with_arr_t;

/* with_arr struct encoder
 *
 * @param	with_arr        struct to encode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int encode_with_arr(with_arr_t *with_arr, char *buf, char *buf_end);

/* with_arr struct decoder
 *
 * @param	with_arr        struct to decode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int decode_with_arr(with_arr_t *with_arr, const char **buf,
                    const char **buf_end);

typedef struct {
  with_arr_t with_arr;
  const char *greeting;
  uint32_t greeting_len;
} hello_request_t;

/* hello_request struct encoder
 *
 * @param	hello_request   struct to encode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int encode_hello_request(hello_request_t *hello_request, char *buf,
                         char *buf_end);

/* hello_request struct decoder
 *
 * @param	hello_request   struct to decode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int decode_hello_request(hello_request_t *hello_request, const char **buf,
                         const char **buf_end);

typedef struct {
  const char *reply;
  uint32_t reply_len;
} hello_response_t;

/* hello_response struct encoder
 *
 * @param	hello_response  struct to encode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int encode_hello_response(hello_response_t *hello_response, char *buf,
                          char *buf_end);

/* hello_response struct decoder
 *
 * @param	hello_response  struct to decode
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */

int decode_hello_response(hello_response_t *hello_response, const char **buf,
                          const char **buf_end);
