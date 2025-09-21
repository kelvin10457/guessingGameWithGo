package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
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
	randomNumber := rand.Intn(100)
	isGuessed := false
	for i := 0; i < chances; i++ {
		fmt.Print("Enter your choice: ")
		numberGuess, err := getInput()

		if err != nil {
			log.Fatal(err)
		}

		if numberGuess == randomNumber {
			isGuessed = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts. \n", (i + 1))
			break
		} else if numberGuess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %d. \n", numberGuess)
		} else {
			fmt.Printf("Incorrect! The number is less than %d. \n", numberGuess)
		}
	}
	if !isGuessed {
		fmt.Println("Game Over. Try again.")
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
	fmt.Print("->")
	level, err := getInput()
	if err != nil {
		log.Fatal(err)
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
