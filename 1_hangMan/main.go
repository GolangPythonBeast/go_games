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

// Derive a word we have to guessing
// Printing the game state
// 	* Print word you are guessing
// 	* Print hangman state
// Read user input
// 	* validate it (e.g. only letters)
// Determine if the letter is a correct guess or not
// 	* if correct, update the guessed letter
// 	* if incorrect, update the hangman state
// if word is guessed, game over -> you win
// If hangman is complete -> game over, you lose

var readerInput = bufio.NewReader(os.Stdin)

var dictionary = []string{"Zombie", "Gopher", "United States of America", "Indonesia", "Nazism", "Apple", "Programming"}

func main() {
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWords(targetWord)

	hangmanState := 0
	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		fmt.Print(">")
		input := readInput()

		if len(input) != 1 {
			fmt.Println("Invalid input")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState += 1
		}
	}
	fmt.Print("Game Over........")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You Win!")
	} else {
		fmt.Println("You Lose!")
	}

}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isGameOver(targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}
func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func getRandomWord() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	targerWord := dictionary[r.Intn(len(dictionary))]
	return targerWord

}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(printHangmanState(hangmanState))
}

func readInput() string {
	input, err := readerInput.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func getWordGuessingProgress(targetWord string, guessedLetter map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetter[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	return result
}

func printHangmanState(hangmanState int) string {
	if hangmanState == 0 {
		return state0()
	} else if hangmanState == 1 {
		return state1()
	} else if hangmanState == 2 {
		return state2()
	} else if hangmanState == 3 {
		return state3()
	} else if hangmanState == 4 {
		return state4()
	} else if hangmanState == 5 {
		return state5()
	} else if hangmanState == 6 {
		return state6()
	} else if hangmanState == 7 {
		return state7()
	} else if hangmanState == 8 {
		return state8()
	} else {
		return state9()
	}

}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

func state0() string {
	state := ""
	return state
}
func state1() string {
	state := `
-
|
|
|
|
	`
	return state
}

func state2() string {
	state := `
-----------
|
|
|
|
	`
	return state
}

func state3() string {
	state := `
-----------
|       |
|
|
|
	`
	return state
}

func state4() string {
	state := `
-----------
|       |
|       0
|     
|
	`
	return state
}

func state5() string {
	state := `
-----------
|       |
|       0
|     /
|
	`
	return state
}

func state6() string {
	state := `
-----------
|       |
|       0
|     / |
|
	`
	return state
}

func state7() string {
	state := `
-----------
|       |
|       0
|     / | \
|
	`
	return state
}

func state8() string {
	state := `
-----------
|       |
|       0
|     / | \
|      / 
|      
	`
	return state
}

func state9() string {
	state := `
-----------
|       |
|       0
|     / | \
|      / \
|      
	`
	return state
}
