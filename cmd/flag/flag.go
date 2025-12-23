/* 
	This package defines all the necessary Flags that our webCrawler needs 
*/
package main

import (
	"flag"
	"time"
)

type Config struct {
	URL 		 		string
	Depth   		int
	Limit 			int
	Delay 			time.Duration
	Concurrency int
	Output 			string
	Verbose 		bool
}

/* initialize the config of the webcrawler */
func initConifg() (*Config){ 
	return &Config{}
}

func ParseConfig(args []string) (*Config, error){ 
	fs := flag.NewFlagSet("crawler", flag.ContinueOnError)

	cfg := initConifg()

	fs.StringVar(&cfg.URL, "url", "", "Starting URL to crawl")
	fs.IntVar(&cfg.Depth, "depth", 2, "Crawl depth")
	fs.IntVar(&cfg.Limit, "limit", 100, "Max pages to crawl")
	fs.DurationVar(&cfg.Delay, "delay", 500*time.Millisecond, "Delay between requests")
	fs.IntVar(&cfg.Concurrency, "concurrency", 4, "Concurrent requests")
	fs.StringVar(&cfg.Output, "out", "output.json", "Output file")
	fs.BoolVar(&cfg.Verbose, "verbose", false, "Verbose logging")

	return cfg, nil
}
