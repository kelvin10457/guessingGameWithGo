package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func getInput() (int, error) {
	valueEntered, err := reader.ReadString('\n')
	if err != nil {
		return -1, errors.New("there was an error reading the input")
	}
	valueEntered = strings.TrimSpace(valueEntered)

	intValueEntered, err := strconv.Atoi(valueEntered)
	if err != nil {
		return -1, errors.New("you should enter a number")
	}
	return intValueEntered, nil
}

func guess(chances int) {
	fmt.Printf("You have %d chances to guess the correct number.\n", chances)
	randomNumber := rand.IntN(100) + 1 //numeros aleatorios entre 0 y 100
	isGuessed := false
	attempts := 0
	for attempts < chances {
		fmt.Print("Enter your choice: ")
		numberGuess, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		attempts++
		if numberGuess == randomNumber {
			isGuessed = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts. \n", attempts)
			break
		} else if numberGuess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %d. \n", numberGuess)
		} else {
			fmt.Printf("Incorrect! The number is less than %d. \n", numberGuess)
		}
	}
	if !isGuessed {
		fmt.Printf("Game Over. Try again. The number was %d \n", randomNumber)
	}
}
func main() {
	fmt.Println(
		"Welcome to the Number Guessing Game!\n" +
			"I'm thinking of a number between 1 and 100.\n" +
			"You have chances to guess the correct number.\n" +
			"Please select the difficulty level:\n" +
			"1. Easy (10 chances)\n" +
			"2. Medium (5 chances)\n" +
			"3. Hard (3 chances)",
	)

	validEntry := false
	level := -1

	for !validEntry {
		fmt.Print("->")
		intValue, err := getInput()
		if err == nil && (intValue == 1 || intValue == 2 || intValue == 3) {
			validEntry = true
			level = intValue
		} else {
			fmt.Println("You should enter a valid level (1,2,3)")
		}
	}

	switch level {
	case 1:
		fmt.Println("Great! You have selected the Easy difficulty level.")
		guess(10)
	case 2:
		fmt.Println("Great! You have selected the Medium difficulty level.")
		guess(5)
	case 3:
		fmt.Println("Great! You have selected the Hard difficulty level.")
		guess(3)
	}
}
