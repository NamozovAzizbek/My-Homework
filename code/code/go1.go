package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)
// json malumot jsonString o'zgaruvchisiga string shaklida berildi
func Main0() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}                   
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {					   					// agar struct json o'qish uchun bo'lsa maydonlari katt harflar bilan
		Ism string `json:"name"`
		Matn string  `json:"Text"`                             // bo'lishi kerak. struct json malumotni o'qish uchun kerak
	}							   								// agar maydonlar kichik harflarda bo'ladigan bo'lsa json malumotlar
	dec := json.NewDecoder(strings.NewReader(jsonStream))	    // o'qilmay qoladi.
								   								// decoder json ni o'qiysdi
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {	   			// decode jsondagi malumotni go a'zgaruvchilarga o'zlashtiradi
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println( m.Ism,": ", m.Matn)
	}
}
