package dataparser

type DataURLs struct {
	DataLocation string `long:"datalocation" short:"d" description:"The default Data URl" default:"https://api.github.com/repositories"`
	URLToPing string `long:"urltoping" short:"p" description:"Ping URL to check if server is up" default:"https://api.github.com1"`
}

