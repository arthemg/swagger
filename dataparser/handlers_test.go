package dataparser

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

type Message struct{}

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

func TestCheckStatusCode(t *testing.T) {
	var statusOk = "https://httpstat.us/200"
	returnStatusCode, err := checkStatusCode(statusOk)
	if err != nil || returnStatusCode != 200 {
		t.Error("Should have received Status code 200, got", returnStatusCode)
	}
	var brokenURl = "https://"
	_, err = checkStatusCode(brokenURl)
	if err == nil {
		t.Error("Should have received an error, got", err)
	}
}

var dp = &models.Jsonrepo{{
	FullName: "artsem",
	HTMLURL:  "artsemURL",
	ID:       12345,
	Login:    "art",
	Name:     "artsemH",
	URL:      "artsem.com",
}}

var tcs = []struct {
	description        string
	reposClient        *models.Jsonrepo
	statusCode         int
	expectedStatusCode int
	expectedBody       string
}{
	{
		description:        "Internal Server Error 500 ",
		reposClient:        &models.Jsonrepo{},
		statusCode:         http.StatusInternalServerError,
		expectedStatusCode: http.StatusInternalServerError,
		expectedBody:       "{\"code\":500,\"message\":\"INTERNAL_SERVER_ERROR\"}\n",
	}, {
		description:        "Resource Not Found",
		reposClient:        &models.Jsonrepo{},
		statusCode:         http.StatusNotFound,
		expectedStatusCode: http.StatusNotFound,
		expectedBody:       "{\"code\":404,\"message\":\"RESOURCE_NOT_FOUND\"}\n",
	}, {
		description:        "No Data Available",
		reposClient:        &models.Jsonrepo{},
		statusCode:         200,
		expectedStatusCode: http.StatusNotFound,
		expectedBody:       "{\"code\":404,\"message\":\"NO_DATA_AVAILABLE\"}\n",
	}, {
		description:        "Get Valid Json",
		reposClient:        dp,
		statusCode:         http.StatusOK,
		expectedStatusCode: http.StatusOK,
		expectedBody:       "[{\"full_name\":\"artsem\",\"html_url\":\"artsemURL\",\"id\":12345,\"login\":\"art\",\"name\":\"artsemH\",\"url\":\"artsem.com\"}]\n",
	},
}

func TestJSONGetMock2(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range tcs {
		resp, _ := json.Marshal(tc.reposClient)
		mockts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(tc.statusCode)
			w.Write(resp)
		}))
		defer mockts.Close()
		testData := &DataURLs{
			DataLocation: mockts.URL,
			URLToPing:    mockts.URL,
		}
		req, err := http.NewRequest("GET", mockts.URL, bytes.NewBuffer(resp))
		assert.NoError(err)
		r := JSONGet(testData)
		w := httptest.NewRecorder()
		params := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{mockts.URL, mockts.URL}}
		r(params).WriteResponse(w, runtime.JSONProducer())
		assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
		assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
	}
}
