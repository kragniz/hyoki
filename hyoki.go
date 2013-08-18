package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"
)

type Section struct {
	name  string
	lines []string
}

func HyokiFile() []byte {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(usr.HomeDir, ".hyoki", "notes")
	notes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Can't open notes file")
	}
	return notes
}

func Notes() []string {
	notes := HyokiFile()

	sections := make(map[string][]string)
	currentSection := ""
	for _, line := range strings.Split(string(notes), "\n") {
		if !strings.HasPrefix(line, "  ") && len(line) > 0 {
			sections[line] = []string{}
			currentSection = line
		} else if currentSection != "" {
			sections[currentSection] = append(sections[currentSection], line)
		}
	}
	return strings.Split(string(notes), "\n")
}

func main() {
	for _, line := range Notes() {
		fmt.Println(line)
	}
}
