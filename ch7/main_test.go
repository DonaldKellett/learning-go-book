package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	ping(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "HTTP status code should be 200 OK")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"pong"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "GET /ping should return the message \"pong\"")
}
