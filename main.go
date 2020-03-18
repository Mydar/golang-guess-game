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

	for n := 10; n < 1000000; n *= 10 {
		number := rand.NewSource(time.Now().UnixNano())
		random := rand.New(number)
		correctGuess := random.Intn(n)

		var inputValue int
		var loop bool = true

		for i := 0; i < 3; i++ {
			fmt.Printf("Pick a number between zero to %d: ", n)
			fmt.Scan(&inputValue)
			if inputValue > correctGuess && i < 2 {
				fmt.Println("You might want to pick a lesser number!")
			} else if inputValue < correctGuess && i < 2 {
				fmt.Println("You should try a higher number!")
			} else if inputValue == correctGuess {
				fmt.Println("Congratulations, You won!")
				fmt.Println("You made it to the next Stage!")
				break;
			}
			if i == 2 && inputValue != correctGuess {
				fmt.Println("Wrong!!! Better luck next time.")
				fmt.Println("Game Over!")
				loop = false;
		
			}
		
		}

		if loop == true {
			continue
		} else {
			break
		}
	}
}