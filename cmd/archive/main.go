package main

import (
	"io/ioutil"
	"log"
	"strings"
)

//go run cmd/archive/main.go
func GetCurrDir() {
	dir, err := ioutil.ReadDir("./")
	if err != nil {
		return
	}
	for _, file := range dir {
		if file.IsDir() {
			log.Printf("%s is dir", file.Name())
		} else {
			log.Printf("%s is file", file.Name())
			if strings.HasSuffix(file.Name(), "go") {
				log.Printf("%s is go file", file)
			}
		}
	}
}

func main() {
	GetCurrDir()
}
