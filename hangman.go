// time spent: 33 minutes
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var words = []string{"apple", "banana", "cat"}

type Hangman struct {
	word         string
	numOfGuesses int
}

func start() {
	game := Hangman
	word := getWord()
	fmt.Printf("Word: %v \n", word)
	drawBoard(word)
	guessLetter := getGuess()
	fmt.Printf("%v", guessLetter)
	// redrawBoard(checkLetter)

}

func getWord() string {
	// rand.Intn doesn't seem random after it runs once
	return words[rand.Intn(len(words))]
}

func drawBoard(word string) {
	for i := 0; i < len(word); i++ {
		fmt.Printf("_")
	}
	fmt.Printf("\n")
}

func redrawBoard(guess bool) {

}

func checkLetter(letter, word string) bool {
	if strings.Contains(word, letter) {
		return true
	}
	return false
}

func getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return guess
}

func main() {
	start()
}
