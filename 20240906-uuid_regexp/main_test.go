package main

import (
	"regexp"
	"testing"
)

const (
	UuidText = "123e4567-e89b-12d3-a456-426614174000"
	TestText = "v4/" + UuidText + "/UUID/"
)

func TestFullSupport(t *testing.T) {
	re, err := regexp.Compile("([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})")
	if err != nil {
		t.Error(err.Error())
	}

	match := re.FindStringSubmatch(TestText)
	if len(match) == 0 {
		t.Error("No match found")
	}

	if match[1] != UuidText {
		t.Error("Match found but not the expected value")
	}
}

func TestEasyText(t *testing.T) {
	re, err := regexp.Compile(`([0-9a-f\-]{36})`)
	if err != nil {
		t.Error(err.Error())
	}

	match := re.FindStringSubmatch(TestText)
	if len(match) == 0 {
		t.Error("No match found")
	}

	if match[1] != UuidText {
		t.Error("Match found but not the expected value")
	}
}

func TestHexText(t *testing.T) {
	re, err := regexp.Compile(`([[:xdigit:]\-]{36})`)
	if err != nil {
		t.Error(err.Error())
	}

	match := re.FindStringSubmatch(TestText)
	if len(match) == 0 {
		t.Error("No match found")
	}

	if match[1] != UuidText {
		t.Error("Match found but not the expected value")
	}
}
