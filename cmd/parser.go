package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type template struct {
	From     string `json:"from"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}

type email struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}

type customer struct {
	Title     string `csv:"TITLE"`
	FirstName string `csv:"FIRST_NAME"`
	LastName  string `csv:"LAST_NAME"`
	Email     string `csv:"EMAIL"`
}

type parser interface {
	parse()
}

type sender interface {
	sendViaREST()
	sendViaSMTP()
}

func (t *template) parse(file []byte) {
	if err := json.Unmarshal([]byte(file), &t); err != nil {
		fmt.Println(err)
	}

	// fmt.Print(t.From)
}

func (c customer) parse(file *os.File, customers []*customer) {
	if err := gocsv.UnmarshalFile(file, &customers); err != nil {
		fmt.Println(err)
	}
}
