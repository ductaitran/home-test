package main

import (
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	var (
		now   = time.Now()
		today = now.Format("02 Jan 2006")
	)

	var testCases = []struct {
		name      string
		tpl       template
		customers []*customer
		want      []email
	}{
		{
			"TC-MergeToMail",
			template{
				From:     "a@ex.com",
				Subject:  "X",
				MimeType: "text/plain",
				Body:     "{{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, {{TODAY}}",
			},
			[]*customer{
				{Title: "Mr", FirstName: "Smith", LastName: "John", Email: "john.smith@example.com"},
				{Title: "Mrs", FirstName: "Smith", LastName: "Michelle", Email: "michelle.smith@example.com"},
				{Title: "Mrs", FirstName: "Smith", LastName: "Michelle", Email: ""},
			},
			[]email{
				{From: "a@ex.com", To: "john.smith@example.com", Subject: "X", MimeType: "text/plain", Body: "Mr Smith John, " + string(today)},
				{From: "a@ex.com", To: "michelle.smith@example.com", Subject: "X", MimeType: "text/plain", Body: "Mrs Smith Michelle, " + string(today)},
				{From: "a@ex.com", To: "", Subject: "X", MimeType: "text/plain", Body: "Mrs Smith Michelle, " + string(today)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := merge(tc.tpl, tc.customers)
			for i, v := range got {
				if tc.want[i] != v {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}
