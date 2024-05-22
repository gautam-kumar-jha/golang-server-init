package goinittest

import (
	"flag"
	"fmt"
	"testing"
)

var path string

func init() {
	flag.StringVar(&path, "arg", "", "first argument")
}
func TestAdd(t *testing.T) {
	fmt.Println(path)
	result := SayHello()
	expected := "hello"
	if result != expected {
		t.Errorf("SayHello() = %s; want %s", result, expected)
	}
}
