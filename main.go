package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type stringarray []string

func (t *stringarray) String() string {
	return "string array"
}

func (t *stringarray) Set(value string) error {
	*t = append(*t, value)
	return nil
}

var tablenames stringarray

func contains(list []string, value string) bool {
	for _, entry := range list {
		if entry == value {
			return true
		}
	}

	return false
}

func init() {
	flag.Var(&tablenames, "table", "table name where the INSERT should be stripped")
}

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		println("Please use a pipe to insert data into sqlstrip")
		os.Exit(1)
	}

	flag.Parse()

	r, _ := regexp.Compile("INSERT INTO `([a-zA-Z0-9_-]*)` VALUES")

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "INSERT") {
			name := r.FindStringSubmatch(line)[1]
			if contains(tablenames, name) {
				// This is an insert to one of the table names we should skip
				continue
			}
		}

		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
