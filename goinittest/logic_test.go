package goinittest

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var path string

func init() {
	flag.StringVar(&path, "arg", "", "first argument")
}
func TestAdd(t *testing.T) {
	te, _ := os.Getwd()
	fmt.Println("Test:", path+te)
	result := SayHello()
	expected := "hello"
	if result != expected {
		t.Errorf("SayHello() = %s; want %s", result, expected)
	}
}
