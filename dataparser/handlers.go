package dataparser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// const (
// 	errorMessages = "Wrong URL Address"
// 	serverError   = "There server is down"
// )

/*
	Chekcs the URL against the actual one in case
*/
//func checkURL(dataSource string) bool {
//	if !(dataSource == "https://api.github.com/repositories") {
//		return false
//	}
//	return true
//}

/* TODO: Update or remove this function.
	Check the remote access point to check if it is responding
	retuns boolean
*/
//func pingCheck(dataSource string) bool {
//	p := fastping.NewPinger()
//	ra, err := net.ResolveIPAddr("ip4:icmp", dataSource)
//	if err != nil {
//		//fmt.Println(err)
//		return false
//	}
//	p.AddIPAddr(ra)
//	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
//		// fmt.Printf("Server is UP! IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
//	}
//	p.OnIdle = func() {
//		//fmt.Println("finish")
//	}
//	err = p.Run()
//	if err != nil {
//		fmt.Println(err)
//	}
//	return true
//}

func checkStatusCode(url string) (int, error){
	resp, err := http.Get(url)
	if err != nil {
		//log.Fatal(err)
		return -1, err
	}


	// Print the HTTP Status Code and Status Name
	//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	//
	//if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
	//	fmt.Println("HTTP Status is in the 2xx range")
	//} else {
	//	fmt.Println("Argh! Broken")
	//}
	return resp.StatusCode, nil
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
	fmt.Println("our R", r.Body)
	return json.NewDecoder(r.Body).Decode(target)
}

//JSONGet gets JSON data from source URL and parses neessary
func JSONGet(dataSource *DataURLs) func(params operations.JSONGetParams) middleware.Responder {
	defaultDataSource := dataSource.DataLocation
	defaultPing := dataSource.URLToPing

	return func(params operations.JSONGetParams) middleware.Responder {
		//errorMessages := "Wrong URL Address"
		//serverError := "There server is down"
		//missingJSON := "MISSING_JSON_DATA"
		internalError :="INTERNAL_SERVER_ERROR"
		resourceNotFound := "RESOURCE_NOT_FOUND"
		noDataAvailable := "NO_DATA_AVAILABLE"
		var dataSource = &defaultDataSource
		var urlPing = &defaultPing
		if params.Jsonrepo != nil {
			dataSource = &params.Jsonrepo[0]
			//fmt.Println("dataSource " , dataSource)
			urlPing = &params.Jsonrepo[1]
			//fmt.Println("urlPing", urlPing)
			}

		//Check if the URL is correct or exists
		//if !checkURL(*dataSource) {
		//	return operations.NewJSONGetNotFound().WithPayload(
		//		&models.ErrorResponse{
		//			Code:    400,
		//			Message: &errorMessages,
		//		})
		//}

		//Check if the remote server is Up
		statusCode,_ := checkStatusCode(*urlPing)
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
			repos := make(models.Jsonrepo, 0)
			getJSON(*dataSource, &repos)
			if len(repos) <= 0 {
				return operations.NewJSONGetNotFound().WithPayload(
					&models.ErrorResponse{
						Code:    404,
						Message: &noDataAvailable,
					})
			}



		}
		repos := make(models.Jsonrepo, 0)
		getJSON(*dataSource, &repos)
		return operations.NewJSONGetOK().WithPayload(repos)
		//repos := make(models.Jsonrepo, 0)
		//
		//err :=getJSON(*dataSource, &repos)
		//
		//fmt.Println("LENGTH:", repos)
		////if  repos {
		////	return operations.NewJSONGetNotFound().WithPayload(
		////		&models.ErrorResponse{
		////			Code:    404,
		////			Message: &missingJSON,
		////		})
		////
		////}
		//if err != nil {
		//	return operations.NewJSONGetNotFound().WithPayload(
		//		&models.ErrorResponse{
		//			Code:    500,
		//			Message: &internalError,
		//		})
		//}
		//
		//return operations.NewJSONGetOK().WithPayload(repos)

	}

}
