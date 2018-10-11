package dataparser

//DataURLs Contains the flags definitions to be passed as defaults.
type DataURLs struct {
	DataLocation string `long:"datalocation" short:"d" description:"The default Data URl" default:"https://api.github.com/repositories"`
	URLToPing    string `long:"urltoping" short:"p" description:"Ping URL to check if server is up" default:"api.github.com"`
}