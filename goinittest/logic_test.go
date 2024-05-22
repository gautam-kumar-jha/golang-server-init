package goinittest

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var path string

func init() {
	flag.StringVar(&path, "arg", "", "first argument")
}
func TestAdd(t *testing.T) {
	fmt.Println("Test:", path)

	// Get a list of files in the directory
	files, err := getFiles(path)
	if err != nil {
		log.Fatal(err)
	}

	// Print the list of file names
	for _, file := range files {
		fmt.Println(file)
	}

	result := SayHello()
	expected := "hello"
	if result != expected {
		t.Errorf("SayHello() = %s; want %s", result, expected)
	}
}

func getFiles(dir string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}
