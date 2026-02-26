package main

import (
	"strings"
)
func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	splitAndTrim := strings.Fields(lower)
	return splitAndTrim
}