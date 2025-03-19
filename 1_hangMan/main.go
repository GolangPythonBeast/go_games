package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var dictionary = []string{"zombie", "gopher", "golang", "apple", "programming"}
var maxAttempts = 9 // Maximum incorrect guesses allowed
var reader = bufio.NewReader(os.Stdin)

func main() {
	targetWord := getRandomWord()
	guessedLetters := make(map[rune]bool)
	attempts := 0

	// Reveal first and last letters to help the player
	guessedLetters[rune(targetWord[0])] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true

	for !gameOver(targetWord, guessedLetters, attempts) {
		printGameState(targetWord, guessedLetters, attempts)
		fmt.Print("Enter a letter: ")
		letter := readLetter()

		if guessedLetters[letter] {
			fmt.Println("You've already guessed this letter.")
			continue
		}

		if strings.ContainsRune(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			attempts++
			fmt.Println("Incorrect guess!")
		}
	}

	// Final game result
	printGameState(targetWord, guessedLetters, attempts)
	fmt.Println("Game Over!")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("🎉 You Win! 🎉")
	} else {
		fmt.Println("😞 You Lose! The word was:", targetWord)
	}
}

// Select a random word from the dictionary
func getRandomWord() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	targerWord := dictionary[r.Intn(len(dictionary))]
	return targerWord
}

// Check if the game is over
func gameOver(targetWord string, guessedLetters map[rune]bool, attempts int) bool {
	return isWordGuessed(targetWord, guessedLetters) || attempts >= maxAttempts
}

// Check if the player has guessed all letters correctly
func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, letter := range targetWord {
		if letter != ' ' && !guessedLetters[letter] {
			return false
		}
	}
	return true
}

// Print the current game state
func printGameState(targetWord string, guessedLetters map[rune]bool, attempts int) {
	fmt.Println("\nWord: " + getWordProgress(targetWord, guessedLetters))
	fmt.Println(drawHangman(attempts))
	fmt.Printf("Attempts left: %d\n", maxAttempts-attempts)
}

// Get the current word progress with guessed letters
func getWordProgress(targetWord string, guessedLetters map[rune]bool) string {
	var progress strings.Builder
	for _, letter := range targetWord {
		if letter == ' ' {
			progress.WriteString("  ") // Space remains visible
		} else if guessedLetters[letter] {
			progress.WriteRune(letter) // Show guessed letters
		} else {
			progress.WriteString("_ ") // Hide unguessed letters
		}
	}
	return progress.String()
}

// Read and validate user input
func readLetter() rune {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if len(input) != 1 {
		fmt.Println("Invalid input! Enter a single letter.")
		return readLetter()
	}
	return unicode.ToLower(rune(input[0]))
}

// Draw hangman stages
func drawHangman(attempts int) string {
	states := []string{
		`
		   
		   
		   
		   
		`,
		`
		---
		|
		|
		|
		`,
		`
		---
		|   |
		|
		|
		`,
		`
		---
		|   |
		|   O
		|
		`,
		`
		---
		|   |
		|   O
		|   |
		`,
		`
		---
		|   |
		|   O
		|  /|
		`,
		`
		---
		|   |
		|   O
		|  /|\
		`,
		`
		---
		|   |
		|   O
		|  /|\
		|  /
		`,
		`
		---
		|   |
		|   O
		|  /|\
		|  / \
		`,
	}
	if attempts >= len(states) {
		return states[len(states)-1] // Last hangman stage
	}
	return states[attempts]
}
