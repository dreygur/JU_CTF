package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

// Form Id of the Google Form
// formID := "1FAIpQLSegHaWfpaYRlhmcnKs_vQRyZCKoVLLKcw16DAA-l8i7CyJtwg"
// Full URL of the Google Form
// url := "https://docs.google.com/forms/d/e/" + formID + "/formResponse"

func uname() (string, error) {
	// Returns Random Bangladeshi University name

	// names := [10] string{
	//   "A",
	// } // Array

	names := make([]string, 0) // Slice
	names = append(names,
		"University of Dhaka",
		"University of Rajshahi",
		"Bangladesh Agricultural University",
		"Agricultural Science",
		"Bangladesh University of Engineering & Technology",
		"University of Chittagong",
		"Jahangirnagar University",
		"Islamic University, Bangladesh",
		"Shahjalal University of Science and Technology",
		"Khulna University",
		"Bangabandhu Sheikh Mujib Medical University",
		"Bangabandhu Sheikh Mujibur Rahman Agricultural University",
		"Agricultural Science",
		"Hajee Mohammad Danesh Science & Technology University",
		"Mawlana Bhashani Science and Technology University",
		"Patuakhali Science and Technology University",
		"Sher-e-Bangla Agricultural University",
		"Agricultural Science",
		"Noakhali Science and Technology University",
		"Chittagong University of Engineering & Technology",
		"Rajshahi University of Engineering & Technology",
		"Khulna University of Engineering & Technology",
		"Dhaka University of Engineering & Technology",
		"Jagannath University",
		"Jatiya Kabi Kazi Nazrul Islam University",
		"Chittagong Veterinary and Animal Sciences University",
		"Sylhet Agricultural University",
		"Agricultural Science",
		"Comilla University",
		"Jessore University of Science & Technology",
		"Pabna University of Science and Technology",
		"Bangladesh University of Professionals",
		"Begum Rokeya University",
		"Bangladesh University of Textiles",
		"Textile Engineering",
		"University of Barisal",
		"Bangabandhu Sheikh Mujibur Rahman Science and Technology University",
		"Rangamati Science and Technology University",
		"Bangabandhu Sheikh Mujibur Rahman Maritime University",
		"Rajshahi Medical University",
		"Chittagong Medical University",
		"Rabindra University, Bangladesh",
		"Bangabandhu Sheikh Mujibur Rahman Digital University",
		"Sheikh Hasina University",
		"General",
		"Sheikh Fazilatunnesa Mujib University of Science & Technology",
		"Sylhet Medical University",
		"Khulna Agricultural University",
		"Bangabandhu Sheikh Mujibur Rahman Aviation and Aerospace University",
		"International University of Business Agriculture and Technology",
		"North South University",
		"University of Science & Technology Chittagong",
		"Central Women's University",
		"Independent University, Bangladesh",
		"American International University-Bangladesh",
		"Ahsanullah University of Science and Technology",
		"Dhaka International University",
		"The University of Comilla, Bangladesh",
		"Asian University of Bangladesh",
		"East West University",
		"Gono Bishwabidyalay",
		"People's University of Bangladesh",
		"Queens University (Bangladesh)",
		"University of Asia Pacific (Bangladesh)",
		"BGMEA University of Fashion & Technology",
		"Textile manufacturing",
		"Chittagong Independent University (CIU)",
		"Bangladesh University",
		"BGC Trust University Bangladesh",
		"BRAC University",
		"Manarat International University",
		"Premier University, Chittagong",
		"Pundra University of Science and Technology",
		"Southern University, Bangladesh",
		"Sylhet International University",
		"City University, Bangladesh",
		"Daffodil International University",
		"Green University of Bangladesh",
		"IBAIS University",
		"Leading University",
		"Northern University, Bangladesh",
		"Prime University",
		"Southeast University (Bangladesh)",
		"Stamford University Bangladesh",
		"State University of Bangladesh",
		"University of Development Alternative",
		"Bangladesh University of Business and Technology",
		"Eastern University, Bangladesh",
		"Metropolitan University, Sylhet",
		"Millennium University",
		"Presidency University, Bangladesh",
		"Primeasia University",
		"Royal University of Dhaka",
		"Shanto-Mariam University of Creative Technology",
		"United International University",
		"University of Information Technology and Sciences",
		"University of South Asia, Bangladesh",
		"Uttara University",
		"Victoria University of Bangladesh",
		"World University of Bangladesh",
		"Atish Dipankar University of Science and Technology",
		"University of Liberal Arts Bangladesh",
		"Asa University Bangladesh",
		"Bangladesh Islami University",
		"East Delta University",
		"Britannia University",
		"Feni University",
		"Khwaja Yunus Ali University",
		"Bangladesh University of Health Sciences",
		"European University of Bangladesh",
		"First Capital University of Bangladesh",
		"Hamdard University Bangladesh",
		"Ishakha International University, Bangladesh",
		"North East University",
		"North Western University, Bangladesh",
		"Port City International University",
		"Sonargaon University",
		"Varendra University",
		"Z H Sikder University of Science & Technology",
		"Rajshahi Science & Technology University",
		"Exim Bank Agricultural University Bangladesh",
		"Fareast International University",
		"German University Bangladesh",
		"Notre Dame University Bangladesh",
		"Sheikh Fazilatunnesa Mujib University",
		"Times University, Bangladesh",
		"Bangladesh Army International University of Science & Technology",
		"Bangladesh Army University of Engineering & Technology",
		"Bangladesh Army University of Science and Technology",
		"CCN University of Science & Technology",
		"Global University Bangladesh",
		"NPI University of Bangladesh",
		"Rabindra Maitree University",
		"The International University of Scholars",
		"University of Creative Technology Chittagong",
		"Anwer Khan Modern University",
		"Central University of Science and Technology",
		"University of Global Village",
		"Z.N.R.F. University of Management Sciences",
		"Islamic University of Technology",
		"Asian University for Women",
		"Bangladesh Open University",
		"National University of Bangladesh",
		"Islamic Arabic University",
		"Patuakhali Science and Technology University",
		"Chittagong University of Engineering & Technology",
		"Bangladesh University of Professionals",
		"Islamic University, Bangladesh",
		"Bangladesh Agricultural University",
		"Pabna University of Science & Technology",
		"Hajee Mohammad Danesh Science & Technology University",
		"Shahjalal University of Science and Technology",
		"Global University Bangladesh",
		"BGC Trust University Bangladesh",
		"Ahsanullah University of Science and Technology",
		"North Western University, Bangladesh",
		"Exim Bank Agricultural University Bangladesh",
		"Leading University")
	return names[rand.Intn(len(names))], nil
}

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

	// Currently shows error.
	// Need to fix
	tx := "TrxID " + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())
	+rand.Int(9) + strings.ToUpper(loweralpha())

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

	fmt.Println(uname())
}
