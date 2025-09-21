package main

import (
	"bufio"
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
		return -1, err
	}
	valueEntered = strings.Replace(valueEntered, "\n", "", 1)
	intValueEntered, err := strconv.Atoi(valueEntered)
	if err != nil {
		return -1, err
	}
	return intValueEntered, nil
}

func guess(chances int) {
	randomNumber := rand.Intn(100)
	for i := 0; i < chances; i++ {
		fmt.Print("Enter your choice: ")
		numberGuess, err := getInput()
		fmt.Println()
		if err != nil {

		}
		if numberGuess == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts. \n", (i + 1))
		} else if numberGuess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %d. \n", numberGuess)
		} else {
			fmt.Printf("Incorrect! The number is less than %d. \n", numberGuess)
		}
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
		log.Fatal("Hubo un error al seleccionar el nivel")
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
