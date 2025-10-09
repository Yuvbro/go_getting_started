package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Nevo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf("Hello(%q) = %q, %v, want match %v", name, msg, err, want)
	}
}

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf("Hello(%q) = %q, %v, want match", "", msg, err)
	}
}
