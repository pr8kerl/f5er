package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
func prettifyString(string input) (string, error) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	tabs := 0

	for scanner.Scan() {
		tok := scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", tabs)

}
*/
func prettifyScanner(input string) {

	printtabs := func(m int) {
		for i := 0; i < m; i++ {
			fmt.Printf("\t")
		}
	}

	tabs := 0
	open := false
	for _, tok := range input {

		switch {
		case tok == '"':
			if open {
				open = false
			} else {
				open = true
			}
		case tok == '{':
			if !open {
				fmt.Printf("\n")
				printtabs(tabs)
				tabs++
				fmt.Println(string(tok))
				printtabs(tabs)
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == '}':
			if !open {
				fmt.Printf("\n")
				tabs--
				printtabs(tabs)
				fmt.Printf("%s", string(tok))
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == '[':
			if !open {
				fmt.Printf("\n")
				printtabs(tabs)
				tabs++
				fmt.Println(string(tok))
				printtabs(tabs)
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == ']':
			if !open {
				fmt.Printf("\n")
				tabs--
				printtabs(tabs)
				fmt.Printf("%s", string(tok))
			} else {
				fmt.Printf("%s", string(tok))
			}
		case tok == ',':
			fmt.Println(string(tok))
			printtabs(tabs)
		case tok == '\n':
		default:
			fmt.Printf("%s", string(tok))
		}
	}
	fmt.Println()

}

func prettifyBytes(input string) {

	f := func(c rune) bool {
		return (c == '{' || c == '}')
	}
	substrings := strings.FieldsFunc(input, f)
	for _, v := range substrings {
		fmt.Printf("pretty: %s\n", v)
	}

}

func printResponse(input interface{}) {

	jsonresp, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonresp))

}
