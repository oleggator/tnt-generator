
#include <stdint.h>
#include <msgpuck.h>
#include <stdio.h>

#include "models.h"



int encode_with_arr(with_arr_t * with_arr,
	char ** buf, char ** buf_end)
{
	// field arr
}


int decode_with_arr(with_arr_t * with_arr,
	const char ** buf, const char ** buf_end)
{
	// field arr
}



int encode_hello_request(hello_request_t * hello_request,
	char ** buf, char ** buf_end)
{
	// field with_arr
	// field greeting
}


int decode_hello_request(hello_request_t * hello_request,
	const char ** buf, const char ** buf_end)
{
	// field with_arr
	// field greeting
}



int encode_hello_response(hello_response_t * hello_response,
	char ** buf, char ** buf_end)
{
	// field reply
}


int decode_hello_response(hello_response_t * hello_response,
	const char ** buf, const char ** buf_end)
{
	// field reply
}

