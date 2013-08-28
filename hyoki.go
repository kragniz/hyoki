package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

type Notes map[string][]string
type NotesFiles map[string]Notes
type NoteFileList map[string][]byte

func HyokiPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".hyoki")
}

func HyokiFilenames() []string {
	file, _ := os.Open(HyokiPath())
	names, _ := file.Readdirnames(0)
	hykFiles := []string{}
	for _, name := range names {
		fmt.Println(name)
		if strings.HasSuffix(name, ".hyk") {
			hykFiles = append(names, name)
		}
	}
	return hykFiles
}

func HyokiFiles() NoteFileList {
	path := HyokiPath()

	files := make(NoteFileList)
	for _, filename := range HyokiFilenames() {
		fmt.Println("wooo", filename)
		notes, err := ioutil.ReadFile(filepath.Join(path, filename))
		files[filename] = notes

		if err != nil {
			fmt.Println(path, "doesn't seem to be a notes file")
		}

	}
	return files
}

func notes() NotesFiles {
	notes := HyokiFiles()
	notesFiles := make(NotesFiles)
	for _, filename := range notes {
		sections := make(Notes)
		currentSection := ""
		filename := string(filename)
		for _, line := range strings.Split(string(notes[filename]), "\n") {
			if !strings.HasPrefix(line, "  ") && len(line) > 0 {
				sections[line] = []string{}
				currentSection = line
			} else if currentSection != "" {
				sections[currentSection] = append(sections[currentSection], line)
			}
		}
		notesFiles[filename] = sections
	}
	return notesFiles
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

func PrintFileSections(notesFiles NotesFiles, exp string) {
	for filename, _ := range notesFiles {
		fmt.Println("hi there!")
		PrintSections(notesFiles[filename], exp)
	}
}

func listSections(notes Notes) {
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

func ListFiles(notes NotesFiles) {
	for filename, _ := range notes {
		listSections(notes[string(filename)])
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

func main() {
	notes := notes()
	args := os.Args

	if len(args) > 1 {
		firstArg := args[1]
		switch {
		case firstArg == "list-sections":
			ListFiles(notes)
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
		PrintFileSections(notes, args[1])
	} else {
		PrintFileSections(notes, "")
	}
}
