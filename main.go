package main

import (
	"fmt"
)

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	// how many x times money we will get for each symbol
	multiplies := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	balance := uint(200)
	name := GetName()
	fmt.Printf("Welcome %s!\n", name)

	symbolSlice := GenerateSymbolArray(symbols)

	for int(balance) >= 0 {
		bet := GetBet(balance)

		if bet == 0 {
			break
		}

		balance -= bet

		spin := GetSpin(symbolSlice, 3, 3)
		PrintSpin(spin)
		winningLines := checkWin(spin, multiplies)

		for i, multi := range winningLines {
			win := multi * bet
			balance += win

			if multi > 0 {
				fmt.Printf("Won %d, (%dx) on Line #%d\n", win, multi, i+1)
			}
		}
	}

}

func checkWin(spin [][]string, multiplies map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}

		if win {
			lines = append(lines, multiplies[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
}
