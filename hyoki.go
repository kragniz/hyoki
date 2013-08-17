package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    content, err := ioutil.ReadFile("notes")
    if err != nil {
        fmt.Println("Can't open notes file")
    }
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        fmt.Println(line)
    }
}
