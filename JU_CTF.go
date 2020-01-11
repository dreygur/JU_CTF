package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/JU_CTF/uni"
)

// Form Id of the Google Form
// formID := "1FAIpQLSegHaWfpaYRlhmcnKs_vQRyZCKoVLLKcw16DAA-l8i7CyJtwg"
// Full URL of the Google Form
// url := "https://docs.google.com/forms/d/e/" + formID + "/formResponse"

func name(n int) string {
	// This function returns some random names
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_."
	name := make([]byte, n) // Making Slices
	for i := range name {
		name[i] = chars[rand.Intn(len(chars))]
	}

	fmt.Println(string(name))
	return string(name) // Stringify the output and return
}

func email(n int) string {
	// Returns a email address

	// Need to define more email addresses
	domains := []string{"@gmail.com", "@yahoo.com", "@outlook.com", "@yandex.com", "@hotmail.com"}
	fmt.Println(domains)
	email := name(rand.Intn(n)) + domains[rand.Intn(len(domains))]
	return email
}

func loweralpha() string {
	// Generating alphanumeric charecters
	p := make([]byte, 26) // Slice
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p) // Return chars as string
}

func text() string {
	// Generating random & dynamic Transaction ID's

	txt := "TrxID " + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha()) +
		strconv.Itoa(rand.Intn(9)) + strings.ToUpper(loweralpha())

	return txt[:16] // The first 16 chars are useful to us the rest are junk
}

func formData() (string, error) {
	// This Function isn't implemented yet!
	// Makes data organized and sendable as POST request t google form

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

	fmt.Println(email(6))

	fmt.Println(uni.Uname())
}
