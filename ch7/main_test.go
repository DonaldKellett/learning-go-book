package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func NilLogger(_ string) {}

func TestHelloGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode, "HTTP status code should be 405 Method Not Allowed")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"GET method not supported"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "GET /hello should return the message \"GET method not supported\"")
}

func TestHelloPostEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/hello", nil)
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "HTTP status code should be 400 Bad Request")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"POST request body invalid"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "POST /hello with empty body should return the message \"POST request body invalid\"")
}

func TestHelloPostInvalidUserID(t *testing.T) {
	reqBody := GreetingRequestBody{
		UserID: 4,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/hello", strings.NewReader(string(jsonReqBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "HTTP status code should be 400 Bad Request")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"unknown user"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "POST /hello with invalid user_id should return the message \"unknown user\"")
}

func TestHelloPostValidUserID1(t *testing.T) {
	reqBody := GreetingRequestBody{
		UserID: 1,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/hello", strings.NewReader(string(jsonReqBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "HTTP status code should be 200 OK")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"Hello, Fred"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "POST /hello with user_id=1 should return the message \"Hello, Fred\"")
}

func TestHelloPostValidUserID2(t *testing.T) {
	reqBody := GreetingRequestBody{
		UserID: 2,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/hello", strings.NewReader(string(jsonReqBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "HTTP status code should be 200 OK")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"Hello, Mary"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "POST /hello with user_id=2 should return the message \"Hello, Mary\"")
}

func TestHelloPostValidUserID3(t *testing.T) {
	reqBody := GreetingRequestBody{
		UserID: 3,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/hello", strings.NewReader(string(jsonReqBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	l := LoggerAdapter(NilLogger)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	c.Hello(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "HTTP status code should be 200 OK")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var actual Message
	expected := Message{"Hello, Pat"}
	e := json.Unmarshal(data, &actual)
	assert.Nil(t, e)
	assert.Equal(t, expected, actual, "POST /hello with user_id=3 should return the message \"Hello, Pat\"")
}
