package main

import (
	"bufio"
	"fmt"
	"github.com/Gabe1203/Finances/balances"
	"os"
)

func main() {
	fmt.Println("            Starting application... ")
	fmt.Println("|--------------------------------------------|\n\n\n\n\n")
	fmt.Println("Do you want to see your current balances? (y/n)")

	//Wait for valid input
	checkBalance, err := readInput()
	if err != nil {
		fmt.Printf("Error reading input: %s", err.Error())
	}

	if checkBalance {
		fmt.Println("Do you want to see a detailed view? (y/n)")
		detailedView, err := readInput()
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}
		balances, err := balances.ReportBalances(detailedView)
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}
		fmt.Println("Here are your balances: ")
		fmt.Println(balances)
		return
	} else {
		fmt.Println("Application terminating... come back for more features.")
		return
	}

	//fmt.Println(char)
	return
}

func readInput() (bool, error) {
	var char rune
	var err error
	//Spin until the user inputs a valid input
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err = reader.ReadRune()
		if err != nil {
			fmt.Printf("Errorr on input: %s.... exiting program", err.Error())
			return false, err
		}
		if char != 'y' && char != 'n' {
			fmt.Println("Please enter a valid character, y or n.")
		} else {
			break
		}
	}
	if char == 'y' {
		return true, nil
	} else {
		return false, nil
	}
}
