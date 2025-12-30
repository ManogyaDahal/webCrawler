/* 
	This package defines all the necessary Flags that our webCrawler needs 
*/
package flags 

import (
	"flag"
	"time"
)

type Config struct {
	URL 		 		string
	Depth   		uint
	Limit 			uint
	Delay 			time.Duration
	Concurrency uint
	Output 			string
	Verbose 		bool
}

/* initialize the config of the webcrawler with default values*/
func initConifg() (*Config){ 
	return &Config{}
}

/* This function takes args from command line and updates the 
config struct accordingly */
func ParseConfig(args []string) (*Config, error){ 
	fs := flag.NewFlagSet("crawler", flag.ContinueOnError)

	cfg := initConifg()

	fs.StringVar(&cfg.URL, "url", "", "Starting URL to crawl")
	fs.UintVar(&cfg.Depth, "depth", 2, "Crawl depth")
	fs.UintVar(&cfg.Limit, "limit", 100, "Max pages to crawl")
	fs.DurationVar(&cfg.Delay, "delay", 500*time.Millisecond, "Delay between requests")
	fs.UintVar(&cfg.Concurrency, "concurrency", 4, "Concurrent requests")
	fs.StringVar(&cfg.Output, "out", "output.json", "Output file")
	fs.BoolVar(&cfg.Verbose, "verbose", false, "Verbose logging")

	if err := fs.Parse(args); err != nil { 
	return nil, err
	}

	return cfg, nil
}

func ValidateUserInput(cfg *Config) []error {
	var errors []error
	
	return nil
}
