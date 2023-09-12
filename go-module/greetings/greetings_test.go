package greetings

import (
	"regexp"
	"testing"
)

func TestGreetingName(t *testing.T) {
	name := "Vincent"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting("Vincent")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Vincent") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
