package goinittest

import "testing"

func TestAdd(t *testing.T) {
	result := SayHello()
	expected := "helloy"
	if result != expected {
		t.Errorf("SayHello() = %s; want %s", result, expected)
	}
}
