package crawl

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

/* finds the link in the provided page */
func findLinks(url string) []string { 
	var FoundUrls []string

	resp, err:= http.Get(url)
	if err != nil { /*Todo*/ }
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil { /*Todo*/ }

	for n := range doc.Descendants(){ 
		if n.Type != html.ElementNode || n.Data != "a" { 
			continue 
		}

		for _, val := range n.Attr{ 
			if val.Key == "href"{
				value := strings.TrimSpace(val.Val)
				if value=="" || 
				strings.HasPrefix(value, "#") ||
				strings.HasPrefix(value, "mailto:")||
				strings.HasPrefix(value, "javacript:"){
					continue
				}
				FoundUrls = append(FoundUrls, value)	
			}
		}
	}
	return FoundUrls
}
