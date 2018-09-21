package dataparser

type PingUrl struct {
	URLToPing string `long:"urltoping" short:"p" description:"Ping URL to check if server is up" default:"https://api.github.com1"`
}