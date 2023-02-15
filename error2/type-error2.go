package main

import "errors"

// func (r *RequestError) Error() string {
// 	return r.Err.Error()
// }

func (r *RequestError) doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}
