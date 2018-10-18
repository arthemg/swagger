package dataparser

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tatsushid/go-fastping"
)

// const (
// 	errorMessages = "Wrong URL Address"
// 	serverError   = "There server is down"
// )

/*
	Chekcs the URL against the actual one in case
*/
func checkURL(dataSource string) bool {
	if !(dataSource == "https://api.github.com/repositories") {
		return false
	}
	return true
}

/*
	Check the remote access point to check if it is responding
	retuns boolean
*/
func pingCheck(dataSource string) bool {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", dataSource)
	if err != nil {
		//fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		// fmt.Printf("Server is UP! IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		//fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

/*
	get data form the remote server to be processed
*/
func getJSON(url string, target interface{}) error {

	var httpClient = &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//JSONGet gets JSON data from source URL and parses neessary
func JSONGet(dataSource *DataURLs) func(params operations.JSONGetParams) middleware.Responder {
	defaultDataSource := dataSource.DataLocation
	defaultPing := dataSource.URLToPing

	return func(params operations.JSONGetParams) middleware.Responder {
		errorMessages := "Wrong URL Address"
		serverError := "There server is down"
		var dataSource = &defaultDataSource
		var urlPing = &defaultPing

		if params.Jsonrepo != nil {
			dataSource = &params.Jsonrepo[0]

			urlPing = &params.Jsonrepo[1]

		}

		//Check if the URL is correct or exists
		if !checkURL(*dataSource) {
			return operations.NewJSONGetNotFound().WithPayload(
				&models.ErrorResponse{
					Code:    404,
					Message: &errorMessages,
				})
		}

		//Check if the remote server is Up
		if !pingCheck(*urlPing) {
			return operations.NewJSONGetNotFound().WithPayload(
				&models.ErrorResponse{
					Code:    500,
					Message: &serverError,
				})
		}
		repos := make(models.Jsonrepo, 0)
		getJSON(*dataSource, &repos)

		return operations.NewJSONGetOK().WithPayload(repos)

	}

}
