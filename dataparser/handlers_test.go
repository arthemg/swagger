package dataparser

import (
	"net/http/httptest"
	"testing"
)

const (
	CorrectURL   string = "https://api.github.com/repositories"
	IncorrectURL string = "https://api.github.com/repositorie"
)

func TestCorrectcheckURL(t *testing.T) {
	reusult := checkURL(CorrectURL)
	if reusult != true {
		t.Error("Testing Correct URlExpected true, got ", reusult)
	}
}

func TestIncorrectURL(t *testing.T) {
	reusult := checkURL(IncorrectURL)
	if reusult == true {
		t.Error("Testing Incorrect URL Expected false, got ", reusult)
	}
}

const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]`

const jsonStreamCorrectSample = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`

type Message struct {
	Name, Text string
}

func TestgetJSON(t *testing.T) {
	url := "https://api.github.com/repositories"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	// repos := []interface{}

}
