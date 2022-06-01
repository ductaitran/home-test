package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// write writes input data to file
// input could be in any types
// current support types: struct(email), struct(customer)
func write(in interface{}, file string) {
	switch in.(type) {
	case []email:
		outputFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer outputFile.Close()

		outputData, err := json.MarshalIndent(in, "", " ")
		if err != nil {
			log.Fatalln(err.Error())
		}

		if _, err = outputFile.Write(outputData); err != nil {
			log.Fatalln(err.Error())
		}
	case []customer:
		csvErrorFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer csvErrorFile.Close()

		if err = gocsv.MarshalFile(in, csvErrorFile); err != nil {
			log.Fatalln(err.Error())
		}
	default:
		log.Panicln("input data is not in a support type")
	}

}
