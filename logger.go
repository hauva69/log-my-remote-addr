package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(body))
}
