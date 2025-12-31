package main

import (
	"os"
)

func CreateNewFile(path string) (*os.File, error){ 
	/* Todo: parse with the directory path */
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}	
	return file, nil
}
