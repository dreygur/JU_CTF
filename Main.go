package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func formData() {
	req, err := json.Marshal(map[string]string{
		"entry.838304731": "name",
	})

	fmt.Printf("%v\n\n%v", req, err)
}

func main() {
	/*
		Traditional main function
	*/

	// Form Id of the Google Form
	formID := "1FAIpQLSegHaWfpaYRlhmcnKs_vQRyZCKoVLLKcw16DAA-l8i7CyJtwg"
	// Full URL of the Google Form
	url := "https://docs.google.com/forms/d/e/" + formID + "/formResponse"
	// Make request and get the Response
	resp, err := http.Get(url)
	if resp.StatusCode == 200 {
		fmt.Printf("[+] Succcess! -> %v\n", resp.StatusCode)
	} else {
		fmt.Println(err.Error())
	}

	formData()
}
