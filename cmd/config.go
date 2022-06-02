package main

import "os"

type Config struct {
	templatePath  string
	customersPath string
	outputPath    string
	errorsPath    string
}

var instance *Config

func getConfig() *Config {
	if instance == nil {
		instance = &Config{
			templatePath:  getOsArg(1, "resource/email_template.json"),
			customersPath: getOsArg(2, "resource/customers.csv"),
			outputPath:    getOsArg(3, "resource/output_emails.json"),
			errorsPath:    getOsArg(4, "resource/errors.csv"),
		}
	}
	return instance
}

func getOsArg(arg int, defaultValue string) string {
	if len(os.Args) <= arg || os.Args[arg] == "" {
		return defaultValue
	}

	return os.Args[arg]
}
