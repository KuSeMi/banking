package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	custmers := []Customer{
		{"Siro", "Kyiv", "01001"},
		{"Eugene", "Kyiv", "01001"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(custmers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(custmers)
	}

}
