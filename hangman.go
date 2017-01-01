// time spent: 90 minutes
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

const maxGuesses = 7

// Hangman is the game object you instantiate
// when creating a new Hangman game
type Hangman struct {
	guesses        string
	matchedLetters int
	maxGuesses     int
	numOfTries     int
	word           string
}

func (h *Hangman) play() {
	h.drawBoard()
	h.getGuess()
}

func getWord() string {
	// rand.Intn doesn't seem random after it runs once
	return words[rand.Intn(len(words))]
}

func (h *Hangman) drawBoard() {
	fmt.Println("inside drawBoard")
	h.matchedLetters = 0
	fmt.Printf("Guesses left: %v \n", h.maxGuesses-h.numOfTries)
	for _, letter := range h.guesses {
		fmt.Printf("Letters you've tried: %v", letter)
	}
	fmt.Printf("\n")

	for _, l := range h.word {
		if strings.ContainsRune(h.guesses, l) {
			h.matchedLetters++
			fmt.Printf("%#v", l)
		} else {
			fmt.Printf("_")
		}
	}
	fmt.Printf("\n")
}

func (h *Hangman) getGuess() {
	reader := bufio.NewReader(os.Stdin)
	guess, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	h.numOfTries++
	h.guesses += guess
}

func (h *Hangman) continueGame() bool {
	if len(h.guesses) == h.maxGuesses {
		fmt.Println("you've finished your hangman game, losing.")
		return false
	}
	if h.matchedLetters == len(h.word) {
		fmt.Println("you've finished your hangman game, and you've won, congrats!")
		return false
	}
	fmt.Println("continue playing...")
	return true
}

func main() {
	fmt.Println("Welcome to Hangman")

	game := Hangman{
		word:       getWord(),
		maxGuesses: maxGuesses,
		numOfTries: 0,
	}
	fmt.Printf("Word: %v \n", game.word)

	continueGame := game.continueGame()

	if continueGame {
		game.play()
		fmt.Println("at end of continueGame loop")
		continueGame = game.continueGame()
	}
}
