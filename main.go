package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	fmt.Println("WELCOME!!!")
	fmt.Println("This game gives you the chance to make only three guesses")
	fmt.Println("Ready...")
	fmt.Println("Set...")

	number := rand.NewSource(time.Now().UnixNano())
	random := rand.New(number)
	correctGuess := random.Intn(10)

	var inputValue int

	for i := 0; i < 3; i++ {
		fmt.Println("Pick a number between zero to 10: ")
		fmt.Scan(&inputValue)
	if inputValue > correctGuess && i < 2 {
		fmt.Println("You might want to pick a lesser number!")
	} else if inputValue < correctGuess && i < 2 {
		fmt.Println("You should try a higher number!")
	} else if inputValue == correctGuess {
		fmt.Println("Congratulations, You won!")
		fmt.Println("Game Over!")
		break;
	} 
	if i == 2 && inputValue != correctGuess {
		fmt.Println("Wrong!!! Better luck next time.")
		fmt.Println("Game Over!")
	}
		
	}
}