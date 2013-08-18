package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"regexp"
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

func Notes() map[string][]string {
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
	return sections
}

func PrintSections(notes map[string][]string, exp string) {
	rexp, _ := regexp.Compile(exp)
	for section := range notes {
		if rexp.Match([]byte(section)) {
			fmt.Println(section)
			for _, line := range notes[section] {
				fmt.Println(line)
			}
		}
	}
}

func main() {
	notes := Notes()
	PrintSections(notes, "^s")
}
