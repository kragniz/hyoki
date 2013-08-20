package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

type Notes map[string][]string

func HyokiFile() []byte {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(usr.HomeDir, ".hyoki", "notes")
	notes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(path, "doesn't seem to be a notes file")
	}
	return notes
}

func notes() Notes {
	notes := HyokiFile()

	sections := make(Notes)
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

func PrintSections(notes Notes, exp string) {
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

func ListSections(notes Notes) {
	i := 0
	for section := range notes {
		fmt.Print(section, "")
		i++
		if i < len(notes) {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	notes := notes()
	args := os.Args

	if len(args) > 1 {
		firstArg := args[1]
		switch {
		case firstArg == "list-sections":
			ListSections(notes)
			return
		}

		PrintSections(notes, args[1])
	} else {
		PrintSections(notes, "")
	}

}
