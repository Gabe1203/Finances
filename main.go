package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		exitMessage(f)
		return
	} else {
		exitMessage(f)
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

//Message returned when app exits
//TODO: persist net worth to be a greeting
func exitMessage(f *excelize.File) {
	fmt.Println("Thanks for checking up on your finances!")
	netWorth := GetTotal(f)
	fmt.Println("Your net worth is: " + netWorth)
	persistTotal(f)
	fmt.Println("Application terminating... come back for more features.")
}

type netWorthEntry struct {
	Date             string  `json:"date"`
	NetWorth         float64 `json:"netWorth"`
	PreviousNetWorth float64 `json:"newWorth"`
}

//Writes total net worth to a json file if it has a different value
//TODO: add logic to check if we should add a new entry
//TODO: eventually want to store this by date for easy sorting
//TODO: add error signature
func persistTotal(f *excelize.File) {
	//Read current values from file
	jsonFile, err := os.Open("output.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var netWorthEntries []netWorthEntry
	err = json.Unmarshal(byteValue, &netWorthEntries)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Logic of adding next entry here:

	newEntry := netWorthEntry{
		Date:             "01/04/2020",
		NetWorth:         0.0,
		PreviousNetWorth: 0.0,
	}

	//Append the latest entry
	netWorthEntries = append(netWorthEntries, newEntry)

	jsonData, err := json.Marshal(netWorthEntries)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Write new list to the file
	err = ioutil.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
