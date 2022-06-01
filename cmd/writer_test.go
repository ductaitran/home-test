package main

import (
	"os"
	"testing"
)

const testFilePath = "../resource/"
const testFileJson = "test.json"
const testFileCsv = "test.csv"

func setupWriterTest(t *testing.T) func(t *testing.T) {
	t.Log("setup writer test")

	return func(t *testing.T) {
		t.Log("teardown writer test")

		if _, err := os.Stat(testFilePath + testFileJson); err == nil {
			os.Remove(testFilePath + testFileJson)
		}
		if _, err := os.Stat(testFilePath + testFileCsv); err == nil {
			os.Remove(testFilePath + testFileCsv)
		}
	}
}

func TestWrite(t *testing.T) {
	var testCases = []struct {
		name string
		in   interface{}
		file string
	}{
		{"TC-WriteEmailJson", []email{}, testFilePath + testFileJson},
		{"TC-WriteCustomerCsv", []customer{}, testFilePath + testFileCsv},
		// {"TC-WriteUnsupportType", "unsupport", "default", "default"},
	}

	teardown := setupWriterTest(t)
	defer teardown(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			write(tc.in, tc.file)
			_, err := os.Stat(tc.file)
			if err != nil {
				t.Error("output file was not created", err)
			}
		})
	}
}
