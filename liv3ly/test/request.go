package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func DoRequest(handler http.Handler, method, path string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	return res, nil
}
