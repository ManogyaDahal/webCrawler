package main

import (
	"log"
	"os"
	f "ManogyaDahal/webCrawler/cmd/flag"
)

func main(){
	cfg ,err := f.ParseConfig(os.Args[1:])
	if err != nil { 
		log.Fatal("err","error occured while parsing the flags", err)
	}
	/* initiate flag system */
}
