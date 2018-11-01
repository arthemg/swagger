package dataparser

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func checkStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, nil
}

/*
	get data form the remote server to be processed
*/
func getJSON(url string, target interface{}) error {
	var httpClient = &http.Client{Timeout: 10 * time.Minute}
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

var internalError =  "INTERNAL_SERVER_ERROR"
var resourceNotFound = "RESOURCE_NOT_FOUND"
var noDataAvailable = "NO_DATA_AVAILABLE"

//JSONGet gets JSON data from source URL and parses neessary
func JSONGet(dataSource *DataURLs) func(params operations.JSONGetParams) middleware.Responder {
	defaultDataSource := dataSource.DataLocation
	defaultPing := dataSource.URLToPing
	repos := make(models.Jsonrepo, 0)
	return func(params operations.JSONGetParams) middleware.Responder {
		var dataSource = &defaultDataSource
		var urlPing = &defaultPing
		if params.Jsonrepo != nil {
			dataSource = &params.Jsonrepo[0]
			urlPing = &params.Jsonrepo[1]
		}
		//Check if the remote server is Up
		statusCode, _ := checkStatusCode(*urlPing)
		switch statusCode {
		case 500:
			return operations.NewJSONGetInternalServerError().WithPayload(
				&models.ErrorResponse{
					Code:    500,
					Message: &internalError,
				})
		case 404:
			return operations.NewJSONGetNotFound().WithPayload(
				&models.ErrorResponse{
					Code:    404,
					Message: &resourceNotFound,
				})
		case 200:
			getJSON(*dataSource, &repos)
			if len(repos) <= 0 {
				return operations.NewJSONGetNotFound().WithPayload(
					&models.ErrorResponse{
						Code:    404,
						Message: &noDataAvailable,
					})
			} else {
				break
			}
		}
		getJSON(*dataSource, &repos)
		return operations.NewJSONGetOK().WithPayload(repos)
	}
}
