package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Giully314/snippetbox/internal/assert"
)

func TestCommonHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	commonHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	expected := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expected)

	expected = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), expected)

	// Check that the middleware has correctly set the X-Content-Type-Options
	// header on the response.
	expected = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expected)

	// Check that the middleware has correctly set the X-Frame-Options header
	// on the response.
	expected = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expected)

	// Check that the middleware has correctly set the X-XSS-Protection header
	// on the response
	expected = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expected)

	// Check that the middleware has correctly set the Server header on the
	// response.
	expected = "Go"
	assert.Equal(t, rs.Header.Get("Server"), expected)

	// Check that the middleware has correctly called the next handler in line
	// and the response status code and body are as expected.
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
