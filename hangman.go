// time spent:  3hrs
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// TODO: call words from a elementary dictionary API
var words = []string{"apple", "banana", "cat", "dilapidated"}

const maxGuesses = 7

// Hangman is the game object you instantiate
// when creating a new Hangman game
type Hangman struct {
	currentWordState, guesses []string
	maxGuesses, numOfTries    int
	word                      string
}

func getWord() string {
	return words[rand.Intn(len(words))]
}

// TODO: use http/template package to pretty this up and standardize it
func (h *Hangman) drawBoard() {
	fmt.Printf("Guesses left: %v \n", h.maxGuesses-h.numOfTries)
	fmt.Printf("Guesses: %v \n", h.guesses)
	fmt.Println(h.currentWordState)
}

func (h *Hangman) getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	guess = strings.TrimSpace(guess)
	if guess == "" {
		fmt.Println("please enter a letter")
		return guess
	}
	if strings.Contains(strings.Join(h.guesses, ""), guess) {
		fmt.Println("You've already guesses that letter. Please select another one.")
		return ""
	}
	fmt.Printf("This guess %v \n", guess)
	h.numOfTries++
	h.guesses = append(h.guesses, guess)
	return guess
}

func (h *Hangman) isMatch(guess string) bool {
	if strings.Contains(h.word, guess) {
		fmt.Printf("%v is a match for %v word \n", guess, h.word)
		return true
	}
	fmt.Printf("%v is NOT a match for %v word \n", guess, h.word)
	return false
}

// updateWordState grabs the indices of each matched letter and
// replaces the "-" with the letter
func (h *Hangman) updateWordState(letter string) {
	// initialize state
	if letter == " " {
		for i := 0; i < len(h.word); i++ {
			h.currentWordState = append(h.currentWordState, "_")
		}
	} else {
		for i, l := range h.word {
			if letter == string(l) {
				h.currentWordState[i] = letter
			}
		}
	}
}

func (h *Hangman) continueGame() bool {
	if len(h.guesses) == h.maxGuesses {
		fmt.Println("you've finished your hangman game, losing.")
		return false
	}
	if strings.Join(h.currentWordState, "") == h.word {
		fmt.Println("you've finished your hangman game, and you've won, congrats!")
		return false
	}
	fmt.Println("continue playing...")
	return true
}

func main() {
	// for choosing a random word
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Welcome to Hangman")

	game := Hangman{
		word:       getWord(),
		maxGuesses: maxGuesses,
	}
	game.updateWordState(" ")
	fmt.Printf("Word: %v \n", game.word)

	for game.continueGame() {
		game.drawBoard()
		guess := game.getGuess()
		if game.isMatch(guess) {
			game.updateWordState(guess)
		}
	}
}
