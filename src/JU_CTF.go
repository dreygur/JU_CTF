package main

import (
	"encoding/json"
	"fmt"
	"juctf"
	"net/http"
)

func formData() (string, error) {
	// This Function isn't implemented yet!

	req, err := json.MarshalIndent(map[string]string{
		"entry.838304731":  "name",
		"entry.2146648422": "2",
		"entry.2129567656": "0",
		"entry.1114548755": "0",
		"entry.773910481":  "0",
	}, "", "\t")

	// _ = ioutil.WriteFile("test.json", req, 0644)

	return string(req), err
}

func main() {
	// Form Id of the Google Form
	formID := "1FAIpQLSegHaWfpaYRlhmcnKs_vQRyZCKoVLLKcw16DAA-l8i7CyJtwg"
	// Full URL of the Google Form
	url := "https://docs.google.com/forms/d/e/" + formID + "/formResponse"
	// Make request and get the Response
	resp, err := http.Get(url)
	if resp.StatusCode == 200 {
		fmt.Printf("[+] %v -> %v\n", http.StatusText(resp.StatusCode), resp.StatusCode)
	} else {
		fmt.Println(err.Error())
	}
	data, err := formData()
	fmt.Println(data, err)

	fmt.Println(juctf.Uname())
}
