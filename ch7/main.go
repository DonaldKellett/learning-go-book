package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	port uint64
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

type GreetingRequestBody struct {
	UserID int `json:"user_id"`
}

func (c Controller) Hello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	w.Header().Set("Content-Type", "application/json")
	wEnc := json.NewEncoder(w)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		s := fmt.Sprintf("%s method not supported", r.Method)
		wEnc.Encode(Message{s})
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		wEnc.Encode(Message{"failed to read body of POST request"})
		return
	}
	var grb GreetingRequestBody
	err = json.Unmarshal(b, &grb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		wEnc.Encode(Message{"POST request body invalid"})
		return
	}
	userID := strconv.Itoa(grb.UserID)
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		wEnc.Encode(Message{"unknown user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	wEnc.Encode(Message{message})
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

type Message struct {
	Message string `json:"message"`
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{"pong"})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Uint64Var(&port, "port", 0, "A port number in the range 1-65535, both inclusive")
	flag.Parse()
	if port < 1 || port > 65535 {
		s := fmt.Sprintf("Expected input port number %d to be within the closed range 1-65535", port)
		panic(s)
	}
	http.HandleFunc("/ping", ping)

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.Hello)

	addr := fmt.Sprintf(":%d", port)
	e := http.ListenAndServe(addr, nil)
	check(e)
}
