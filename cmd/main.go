package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var (
	config    = getConfig()
	tpl       template
	cm        customer
	customers []*customer
)

func merge(tpl template, customers []*customer) ([]email, []customer) {
	currentTime := time.Now()
	var mailList []email
	var errorCustomers = []customer{}
	for _, customer := range customers {
		body := strings.ReplaceAll(tpl.Body, "{{TODAY}}", currentTime.Format("02 Jan 2006"))
		body = strings.ReplaceAll(body, "{{TITLE}}", customer.Title)
		body = strings.ReplaceAll(body, "{{FIRST_NAME}}", customer.FirstName)
		body = strings.ReplaceAll(body, "{{LAST_NAME}}", customer.LastName)
		// fmt.Println(body)
		if customer.Email != "" {
			mailList = append(mailList, email{
				From:     tpl.From,
				To:       customer.Email,
				Subject:  tpl.Subject,
				MimeType: tpl.MimeType,
				Body:     body,
			})
		} else {
			errorCustomers = append(errorCustomers, *customer)
		}
	}

	return mailList, errorCustomers
}

func main() {
	// jsonFile, err := os.Open(config.templatePath)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer jsonFile.Close()

	// byteValue, err := io.ReadAll(jsonFile)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }

	// tpl.parse(byteValue)
	tpl.parse(config.templatePath)

	// csvCustomerFile, err := os.OpenFile(config.customersPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// defer csvCustomerFile.Close()

	// if err := gocsv.UnmarshalFile(csvCustomerFile, &customers); err != nil {
	// 	log.Panic(err)
	// }

	csvErrorFile, err := os.OpenFile(config.errorsPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}
	defer csvErrorFile.Close()

	mailList, errorCustomers := merge(tpl, customers)

	err = gocsv.MarshalFile(&errorCustomers, csvErrorFile)

	outputFile, err := os.OpenFile(config.outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err.Error())
	}
	defer outputFile.Close()
	outputData, err := json.MarshalIndent(mailList, "", " ")
	if err != nil {
		log.Panic(err.Error())
	}

	_, err = outputFile.Write(outputData)
	if err != nil {
		log.Panic(err.Error())
	}
}
