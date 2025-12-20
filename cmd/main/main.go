package main

import(
	"crypto/tls"
	"fmt"
	"flag" //flag helps to parse command line args
	"os"
	"net/http"
	"net/url"
	"github.com/jackdanger/collectlinks"
)

var visited = make(map[string]bool)

func main(){
	// Parsing the arguments(webUrl) from the command line
	flag.Parse()
	args := flag.Args()

	/* needs a rework here for adding a valid url */
	if len(args) < 1 {
		fmt.Println("Please specify the start page to crawl from")
		os.Exit(1)
	}

	//This gives us new channel that recieves and delivers strings
	queue := make(chan string)

	go func(){	//async function
		queue <- args[0] 
	}()

	for url := range queue {
		enqueue(url,queue)
	}
}

func enqueue(url string, queue chan string){
	fmt.Println("fetching", url)
	visited[url] = true

	tlsConfig := &tls.Config{
		InsecureSkipVerify:true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{ Transport: transport}

	resp, err := client.Get(url)
	if err != nil{
		return 
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := fixUrl(link, url)

		if url != "" {
			if !visited[absolute]{ //not to enqueue page twice
				go func() {queue <- absolute}()
			}
		}
	}
}

func fixUrl(href, base string)(string){
	url , err := url.Parse(href)
	if err != nil {
		return ""
	}

	baseUrl , err := url.Parse(base)
	if err != nil {
		return ""
	}

	url = baseUrl.ResolveReference(url)
	return url.String()
}
