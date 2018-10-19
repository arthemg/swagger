package dataparser

import (
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	CorrectURL   string = "https://api.github.com/repositories"
	IncorrectURL string = "https://api.github.com/repositorie"
)

func TestCorrectcheckURL(t *testing.T) {
	result := checkURL(CorrectURL)
	if result != true {
		t.Error("Testing Correct URlExpected true, got ", result)
	}
}

func TestIncorrectURL(t *testing.T) {
	result := checkURL(IncorrectURL)
	if result == true {
		t.Error("Testing Incorrect URL Expected false, got ", result)
	}
}

type Message struct {
	// Name, Text string
	mess []interface{}
}

func TestGetJSON(t *testing.T) {
	url := "https://api.github.com/repositories"
	var messages []Message
	err := getJSON(url, &messages)
	if err != nil {
		t.Error("TestGetJSON failed SHOULD NOT HAVE error, got", err)
	}

	brokenURL := "THIS DOES NOT EXISTS AS URL"
	messages = []Message{}
	err = getJSON(brokenURL, &messages)

	if err == nil {
		t.Error("TestGetJSON failed should HAVE error, got", err)
	}
}

func TestPingCheck(t *testing.T) {
	status := pingCheck("google.com")
	if !status {
		t.Error("Testing Incorrect Server should be true, got ", status)
	}
	wrongStatus := pingCheck("https://google.com")
	if wrongStatus {
		t.Error("Testing Incorrect Server should be false, got ", wrongStatus)
	}
}

func TestJSONGet(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	testData := &DataURLs{
		DataLocation: "https://api.github.com/repositories",
		URLToPing:    "api.github.com",
	}
	params := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositories", "api.github.com"}}
	r := JSONGet(testData)
	w := httptest.NewRecorder()
	r(params).WriteResponse(w, runtime.JSONProducer())
	if w.Code != 200 {
		t.Error("Should receive Status Code 200, got", w.Code)
	}
	IncorrectParams := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositor", "api.github.com"}}
	r = JSONGet(testData)
	w = httptest.NewRecorder()
	r(IncorrectParams).WriteResponse(w, runtime.JSONProducer())
	if w.Code != 404 {
		t.Error("Should receive Status Code 404, got", w.Code)
	}

	IncorrectURLPing := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositories", "https://httpstat.us/404"}}
	r = JSONGet(testData)
	w = httptest.NewRecorder()
	r(IncorrectURLPing).WriteResponse(w, runtime.JSONProducer())
	if w.Code != 404 {
		t.Error("Should receive Status Code 404, got", w.Code)
	}


}
