package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	. "github.com/Gabe1203/Finances/balances"
)

const balanceSheet = "balances/balances.xlsx"

func main() {
	//----------------OUTPUT-------------------------
	fmt.Println("            Starting application... ")
	fmt.Print("|--------------------------------------------|\n\n\n\n\n\n\n\n\n")
	fmt.Println("Do you want to see your current balances? (y/n)")
	//-----------------------------------------------

	//Wait for valid input
	f, err := excelize.OpenFile(balanceSheet)
	checkBalance, err := readInput()
	if err != nil {
		fmt.Printf("Error reading input: %s", err.Error())
	}

	if checkBalance {
		//----------------OUTPUT-------------------------
		fmt.Println("Do you want to see a detailed view? (y/n)")
		//-----------------------------------------------

		detailedView, err := readInput()
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}
		balances, err := ReportBalances(detailedView, f)
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}

		//----------------OUTPUT-------------------------
		fmt.Println("Here are your balances: ")
		fmt.Println(balances)
		fmt.Println("Do you want to update checking balance before you quit? (y/n)")
		//-----------------------------------------------

		updateBalance, err := readInput()
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}
		if updateBalance {
			err := UpdateBalances(f)
			if err != nil {
				fmt.Printf("Error updating balances... %s", err.Error())
				return
			}
			fmt.Println("Balance updated correctly.")
		}

		fmt.Println("Application terminating... come back for more features.")
		return
	} else {
		fmt.Println("Application terminating... come back for more features.")
		return
	}
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
