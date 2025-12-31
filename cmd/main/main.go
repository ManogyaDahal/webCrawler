package main

import (
	"log"
	"os"
	f "ManogyaDahal/webCrawler/cmd/flag"
)

func main(){
	/*flags*/
	cfg ,err := f.ParseConfig(os.Args[1:])
	if err != nil { 
		log.Fatal("err:","error occured while parsing the flags", err)
	}
	if errors := f.ValidateUserInput(cfg); errors != nil { 
		for _, err = range errors {
			log.Println("err:",err)
		}
		log.Fatal("err:", "error occured while validating the user input check for --help flag")
	}
	/*flags*/
}
