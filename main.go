package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Options struct {
	lang  string
	input string
	help  bool
}

func (p *Options) validate() error {
	l := map[string]bool{"it": true, "international": true}
	_, ok := l[p.lang]

	if !ok {
		return fmt.Errorf("invalid language: %v. Valid languages: [it, international]", p.lang)
	}
	return nil
}

func loadAlphabet(lang string) map[string]string {
	jsonFile, err := os.Open("alphabet.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	alphabet := make(map[string]map[string]string)
	json.Unmarshal(byteValue, &alphabet)
	return alphabet[lang]
}

func promptForString(prompt string) string {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(line, "\n")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getOptions() Options {
	options := Options{}

	lang := flag.String("l", "international", "language to use for spelling")
	input := flag.String("i", "", "input text")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	options.lang = *lang
	options.input = *input
	options.help = *help

	err := options.validate()
	if err != nil {
		log.Fatalf("Invalid Options: %v", err)
	}
	return options
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("gospell [-l language=(international|it)] [-h] [-i \"Your input text 1234!\"]")
}

func main() {
	options := getOptions()

	if options.help {
		printHelp()
		return
	}

	alphabet := loadAlphabet(options.lang)

	if options.input == "" {
		options.input = promptForString("What do you to spell? ")
	}
	clearScreen()

	fmt.Println(strings.ToUpper(options.input) + "\n")
	for i, letter := range options.input {
		l := string(letter)
		L := strings.ToUpper(l)
		_, ok := alphabet[L]
		char := alphabet[L]
		if !ok {
			char = l
		}
		fmt.Printf("%v.\t %v\n", i, char)
	}
	fmt.Println()
}
