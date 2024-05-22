package main

import "testing"

func TestAdd(t *testing.T) {
	result := SayHello()
	expected := "hellor"
	if result != expected {
		t.Errorf("SayHello() = %s; want %s", result, expected)
	}
}
