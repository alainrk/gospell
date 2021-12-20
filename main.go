package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func loadAlphabet() map[string]string {
	jsonFile, err := os.Open("alphabet.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	alphabet := make(map[string]string)
	json.Unmarshal(byteValue, &alphabet)
	return alphabet
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

func main() {
	alphabet := loadAlphabet()
	input := promptForString("What do you to spell? ")
	clearScreen()
	fmt.Println(strings.ToUpper(input) + "\n")
	for i, letter := range input {
		l := string(letter)
		L := strings.ToUpper(l)
		fmt.Printf("%v.\t %v\n", i, alphabet[L])
	}
	fmt.Println()
}
