package main

import (
	"os"
	"testing"
)

func setupConfigTest(t *testing.T) func(t *testing.T) {
	t.Log("setup config test")
	oldArgs := os.Args

	// set custom os arguments
	os.Args = []string{"", "arg1", "arg2"}

	return func(t *testing.T) {
		t.Log("teardown config test")

		// restore Args
		defer func() { os.Args = oldArgs }()
	}
}

func TestGetOsArg(t *testing.T) {
	var testCases = []struct {
		name         string
		arg          int
		defaultValue string
		want         string
	}{
		{"TC-Argument1", 1, "default", "arg1"},
		{"TC-Argument2", 2, "default", "arg2"},
		{"TC-NoArgument", 1, "default", "default"},
	}

	teardown := setupConfigTest(t)
	defer teardown(t)

	for _, tc := range testCases {
		// setup for TC-NoArgument
		if tc.name == "TC-NoArgument" {
			os.Args = []string{""}
		}
		t.Run(tc.name, func(t *testing.T) {
			got := getOsArg(tc.arg, tc.defaultValue)
			if got != tc.want {
				t.Errorf("getOsArg(%d, %s) got %v, want %v", tc.arg, tc.defaultValue, got, tc.want)
			}
		})
	}
}
