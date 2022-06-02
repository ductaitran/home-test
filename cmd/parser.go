package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

func (t *template) parse(file []byte) {
	if err := json.Unmarshal([]byte(file), &t); err != nil {
		fmt.Println(err)
	}
}

// merge customer data into template
// valid customers are added to a mail list,
// customers without email are filtered
func merge(tpl template, customers []*customer) ([]email, []customer) {
	currentTime := time.Now()
	var mailList []email
	var errorCustomers = []customer{}
	for _, customer := range customers {
		body := strings.ReplaceAll(tpl.Body, "{{TODAY}}", currentTime.Format("02 Jan 2006"))
		body = strings.ReplaceAll(body, "{{TITLE}}", customer.Title)
		body = strings.ReplaceAll(body, "{{FIRST_NAME}}", customer.FirstName)
		body = strings.ReplaceAll(body, "{{LAST_NAME}}", customer.LastName)
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
