package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	var name = "Rafi"
	var want = regexp.MustCompile(`\b` + name + `\b`)
	var msg, err = Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v; want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	var msg, err = Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v; want "", error`, msg, err)
	}
}
