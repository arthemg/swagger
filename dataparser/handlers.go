package dataparser

import (
	"encoding/json"
	"fmt"
	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	GitRepoURL     = "https://api.github.com/repositories"
	GitRepoPingURL = "https://api.github.com"
)

/*
	Chekcs the URL against the actual one in case
 */
func checkUrl(dataSource string) bool {
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
		fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
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

func JsonGet(dataSource *DataURLs, urlSource *PingUrl) func(params operations.JSONGetParams) middleware.Responder {
	var errorMessages = "Wrong URL Address"
	//var serverError = "There server is down"
	defaultDataSource := dataSource.DataLocation
	defaultPing := urlSource.URLToPing

	return func(params operations.JSONGetParams) middleware.Responder {
		log.Println("params", params.Jsonrepo)
		var dataSource = &defaultDataSource
		var urlPing = &defaultPing
		log.Println("urlPing",*urlPing)

		if params.Jsonrepo != nil {
			dataSource = params.Jsonrepo
			log.Println("dataSource", *dataSource)
			//urlPing =params.Jsonrepo
		}



		//Check if the URL is correct or exists
		if !checkUrl(*dataSource) {
			return operations.NewJSONGetNotFound().WithPayload(
				&models.ErrorResponse{
					Code:    404,
					Message: &errorMessages,
				})
		}

		//Check if the remote server is Up
		//if !pingCheck("https://api.github.com") {
		//	return operations.NewJSONGetNotFound().WithPayload(
		//		&models.ErrorResponse{
		//			Code:    500,
		//			Message: &serverError,
		//		})
		//}
		repos := make(models.Jsonrepo, 0)
		getJSON(*dataSource, &repos)

		return operations.NewJSONGetOK().WithPayload(repos)

	}

}
