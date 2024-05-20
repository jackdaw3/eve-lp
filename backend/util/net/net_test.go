package net

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hello string = "hello"
	mu    sync.Mutex
)

func setUpServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		defer mu.Unlock()
		mu.Lock()
		fmt.Fprintln(w, hello)
	})

	server := httptest.NewServer(handler)
	return server
}

func TestGetWithRetries(t *testing.T) {
	server := setUpServer()
	defer server.Close()

	request := server.URL + "/hello"
	content, _, err := Get(request)
	assert.NoError(t, err)
	assert.NotZero(t, len(content))
}
