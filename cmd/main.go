package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var (
	config = getConfig()
)

func main() {
	jsonFile, err := os.Open(config.templatePath)
	if err != nil {
		log.Panic(err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panic(err.Error())
	}
	var tpl template
	tpl.parse(byteValue)

	csvFile, err := os.OpenFile(config.customersPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}
	defer csvFile.Close()
	customers := []*customer{}
	if err := gocsv.UnmarshalFile(csvFile, &customers); err != nil {
		log.Panic(err)
	}

	csvErrorFile, err := os.OpenFile(config.errorsPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}
	defer csvErrorFile.Close()

	currentTime := time.Now()
	var mail []email
	var errorCustomers = []customer{}
	for _, customer := range customers {
		body := strings.ReplaceAll(tpl.Body, "{{TODAY}}", currentTime.Format("02 Jan 2006"))
		body = strings.ReplaceAll(body, "{{TITLE}}", customer.Title)
		body = strings.ReplaceAll(body, "{{FIRST_NAME}}", customer.FirstName)
		body = strings.ReplaceAll(body, "{{LAST_NAME}}", customer.LastName)
		fmt.Println(body)
		mail = append(mail, email{
			From:     tpl.From,
			To:       customer.Email,
			Subject:  tpl.Subject,
			MimeType: tpl.MimeType,
			Body:     body,
		})

		//write error customer to errors.csv
		if customer.Email == "" {
			errorCustomers = append(errorCustomers, *customer)
		}
	}
	err = gocsv.MarshalFile(&errorCustomers, csvErrorFile)

	for _, m := range mail {
		// fmt.Println(m)
		output, err := json.MarshalIndent(m, "", " ")
		if err != nil {
			log.Panic(err.Error())
		}
		err = os.WriteFile(config.outputPath, output, 0644)
	}

}
