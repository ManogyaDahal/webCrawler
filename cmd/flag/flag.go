/*
This package defines all the necessary Flags that our webCrawler needs
*/
package flags

import (
	"errors"
	"flag"
	"time"
)

type Config struct {
	URL 		 		string				 //path starting
	Depth   		uint					 //no. of hops to reach the end website
	Limit 			uint					 //Stop after visiting N number of websites in total
	Delay 			time.Duration  // within how much duration launch a goroutine
	Concurrency uint						//no. of webworker (goroutines) to launch
	Output 			string					//what is your output file
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

/* This function validates the user's input */
func ValidateUserInput(cfg *Config) []error {
	var errors []error = nil
	
	//validation start
	if err := validateURL(cfg); err != nil{ 
		errors = append(errors, err)
	}
	if err := validateDepth(cfg); err != nil{ 
		errors = append(errors, err)
	}
	if err := validateLimit(cfg); err != nil{ 
		errors = append(errors, err)
	}
	if err := validateConcurrency(cfg); err != nil{ 
		errors = append(errors, err)
	}
	if err := validateOutput(cfg); err != nil{ 
		errors = append(errors, err)
	}
	//validation ends

	return errors
}

/* function for validating URL's */
func validateURL(cfg *Config) error { 
	if cfg.URL == ""  { return errors.New("URL is needed, Empty URL error") } 
	if len(cfg.URL)<5 { return errors.New("URL with very small length") } 
	if cfg.URL[:4] != "http" && (cfg.URL[4:5] != "s") { return errors.New("URL should have http or https") } 
	/* Todo: validate for host lenght */
	return nil
}

/* depth means the number of hops requires to reach a specific url from starting page*/
func validateDepth(cfg *Config) error { 
	if cfg.Depth > 50 { 
		return errors.New("Depth should be between 50 and 0")
	}
	return nil
}

func validateLimit(cfg *Config) error { 
	if cfg.Limit == 0 { 
		return errors.New("Limit should be >0")
	}
	return nil
}

func validateConcurrency(cfg *Config) error { 
	if cfg.Concurrency == 0 || cfg.Concurrency > 100{ 
		return errors.New("Concurrency should lie between 1-100")
	}
	return nil
}

func validateOutput(cfg *Config) error { 
	if len(cfg.Output) < 6 {
		return errors.New("file should have .json extension with some name eg:x.json")
	} 
	if cfg.Output[len(cfg.Output)-5:] != ".json"{ 
		return errors.New("The file provided should be of .json type")
	}
	return nil
}


