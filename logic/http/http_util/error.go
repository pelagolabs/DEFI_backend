package http_util

import "errors"

type HTTPError struct {
	error

	errorCode uint32
}

func WarpHTTPError(err error) *HTTPError {
	return WarpHTTPErrorWithCode(500, err)
}

func WarpHTTPErrorWithCode(errorCode uint32, err error) *HTTPError {
	return &HTTPError{
		error:     err,
		errorCode: errorCode,
	}
}

func NewHttpError(errorCode uint32, err string) *HTTPError {
	return &HTTPError{
		error:     errors.New(err),
		errorCode: errorCode,
	}
}

func NewHttpRedirect(path string) *HTTPError {
	return &HTTPError{
		error:     errors.New(path),
		errorCode: 302,
	}
}
