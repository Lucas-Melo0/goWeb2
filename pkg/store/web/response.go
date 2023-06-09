package web

import (
	"strconv"
)

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {

	if code < 300 {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	if code == 401 {
		return Response{strconv.FormatInt(int64(code), 10), data, "unauthorized"}
	}
	if code < 600 {
		return Response{strconv.FormatInt(int64(code), 10), data, "server error"}
	}
	return Response{strconv.FormatInt(int64(code), 10), data, ""}
}
