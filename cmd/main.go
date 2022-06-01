package main

import (
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

var (
	config    = getConfig()
	tpl       template
	cm        customer
	customers []*customer
)

func main() {
	// parse template
	jsonFile, err := os.Open(config.templatePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	tpl.parse(byteValue)

	// parse customers
	csvCustomerFile, err := os.OpenFile(config.customersPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}
	defer csvCustomerFile.Close()
	if err := gocsv.UnmarshalFile(csvCustomerFile, &customers); err != nil {
		log.Panic(err)
	}

	// merge customers data to template
	mailList, errorCustomers := merge(tpl, customers)

	// write output
	write(errorCustomers, config.errorsPath)
	write(mailList, config.outputPath)
}
