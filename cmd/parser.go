package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type fileType interface {
	parse(filePath string) error
}

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

// func (t *template) parse(file []byte) {
// 	if err := json.Unmarshal([]byte(file), &t); err != nil {
// 		fmt.Println(err)
// 	}
// }

func (t *template) parse(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &t); err != nil {
		return err
	}

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// if byteData, ok := d.([]byte); ok {
	// 	if err := json.Unmarshal(byteData, &t); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	return errors.New("Invalid type")
	// }

	return nil
}

// func (c customer) parse(file *os.File, customers []*customer) {
// 	if err := gocsv.UnmarshalFile(file, &customers); err != nil {
// 		fmt.Println(err)
// 	}
// }

func (c *customer) parse(filePath string) error {
	csvCustomerFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer csvCustomerFile.Close()

	if err := gocsv.UnmarshalFile(csvCustomerFile, &customers); err != nil {
		return err
	}

	return nil
}
