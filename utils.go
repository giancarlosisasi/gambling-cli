package main

import "fmt"

func GetBet(balance uint) uint {
	var bet uint

	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scan((&bet))
		if err != nil {
			fmt.Println("There was an error when storing the bet value. Please try again")
		}

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}

	return bet
}

func GetName() string {
	var name string

	for {
		fmt.Printf("Hi there! What's your name?: ")
		_, err := fmt.Scan(&name)

		if err != nil {
			fmt.Println("There was an error when trying to get your name. Please try again.")
		}

		if name == "" || name == " " {
			fmt.Println("Name can't be blank.")
		} else {
			break
		}
	}

	return name
}
