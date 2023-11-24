package hello

import (
	"testing"
)

func TestGreetingName(t *testing.T) {
	testName := "Antonio"
	msg := Greet(&testName)
	if msg != "Hello, Antonio" {
		t.Fatalf("Test failed, wrong return value %s, expected `Hello, Antonio`", msg)
	}
}

func TestGreetingEmpty(t *testing.T) {
	testName := ""
	msg := Greet(&testName)
	if msg != "Hello, Anonymous" {
		t.Fatalf("Test failed, wrong return value %s, expected `Hello, Anonymous`", msg)
	}
}
