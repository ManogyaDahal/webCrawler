package main

import(
	"fmt"
	"flag" //flag helps to parse command line args
	"os"
	"net/http"
	"io/ioutil"

	"github.com/jackdanger/collectlinks"
)

func main(){
	// Parsing the arguments(webUrl) from the command line
	flag.Parse()
	args := flag.Args()

	/* needs a rework here for adding a valid url */
	if len(args) < 1 {
		fmt.Println("Please specify the start page to crawl from")
		os.Exit(1)
	}
	// Retreving the page html information 
	retrieve(args[0])
}

func retrieve(url string){

	resp, err := http.Get(url)	
	if err != nil {
		return
	}
	//closes the tcp connection to some web servers (after the function ends)
	//LIFO
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body) //using collect links package
	for _,links := range(links){
		fmt.Println(links)
	}

	body, _:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
