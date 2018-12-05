package main

import (
	"https://github.com/oripio/respreader"
	"log"
	"net/http"
)

func main() {
	url := "https://google.com/"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bresp, err := respreader.Decode(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(bresp))
}
