package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	rootRequest, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Root request error: %s", err)
	}

	cases := []struct {
		w                    *httptest.ResponseRecorder
		r                    *http.Request
		expectedResponseCode int
		expectedResponseBody []byte
	}{
		{
			w:                    httptest.NewRecorder(),
			r:                    rootRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte("Welcome!\n"),
		},
	}

	for _, c := range cases {
		//c.r.Header.Set("X-Access-Token", c.accessTokenHeader)

		Index(c.w, c.r)

		if c.expectedResponseCode != c.w.Code {
			t.Errorf("Status Code didn't match:\n\t%q\n\t%q", c.expectedResponseCode, c.w.Code)
		}

		if !bytes.Equal(c.expectedResponseBody, c.w.Body.Bytes()) {
			t.Errorf("Body didn't match:\n\t%q\n\t%q", string(c.expectedResponseBody), c.w.Body.String())
		}
	}
}
