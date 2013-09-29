package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

type Notes map[string][]string

func HyokiPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".hyoki", "notes.hyk")
}

func HyokiFile() []byte {
	path := HyokiPath()
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
			fmt.Println(SectionString(notes, section))
		}
	}
}

func SectionString(notes Notes, section string) string {
	for name := range notes {
		if name == section {
			str := section
			for _, line := range notes[section] {
				str = str + "\n" + line
			}
			return str
		}
	}
	return ""
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

func EditSection(filename string, section string) error {
	cmd := exec.Command("vim", "-c", "/^"+section, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Edit(filename string) error {
	cmd := exec.Command("vim", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GenerateJsonRequest(contents string, filename string) string {
	json := `{
  "description": "hyoki exported gist",
  "public": false,
  "files": {
    "%s": {
      "content": "%s"
    }
  }
}`
	return fmt.Sprintf(json, filename, contents)
}

func PostGist(file string, filename string) string {
	resp, _ := http.Post("https://api.github.com/gists", "text/json",
		strings.NewReader(GenerateJsonRequest(file, filename)))
	body, _ := ioutil.ReadAll(resp.Body)
	htmlRegex := regexp.MustCompile(`"html_url":.?"https://gist.github.com/[0-9a-f]+"`)
	url := htmlRegex.Find(body)
	url = url[12 : len(url)-1]
	return string(url)
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
		case firstArg == "edit":
			if len(args) > 2 {
				section := args[2]
				EditSection(HyokiPath(), section)
			} else {
				Edit(HyokiPath())
			}
			return
		}
		PrintSections(notes, args[1])
	} else {
		Edit(HyokiPath())
	}
}
