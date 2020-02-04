package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	. "github.com/Gabe1203/Finances/operations"
)

const balanceSheet = "sheets/balances.xlsx"

var startBalance float64

func main() {
	//----------------OUTPUT-------------------------
	fmt.Println("            Starting application... ")
	fmt.Print("|--------------------------------------------|\n\n\n\n\n\n\n\n\n")
	fmt.Println("Do you want to see your current balances? (y/n)")
	//-----------------------------------------------

	//Wait for valid input
	f, err := excelize.OpenFile(balanceSheet)
	startBalance, err = GetTotal(f)
	if err != nil {
		fmt.Printf("Error getting starting balance: %s", err.Error())
	}
	//fmt.Println(startBalance)
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
		fmt.Println("Do you want to update balances before you quit? (y/n)")
		//-----------------------------------------------

		updateBalance, err := readInput()
		if err != nil {
			fmt.Printf("Error reading input: %s", err.Error())
		}
		attempts := 0
		if updateBalance {
			for {
				attempts++
				err := UpdateBalances(f)
				if err == nil {
					fmt.Println("Balance updated correctly.")
					break
				} else {
					fmt.Printf("Error updating balances... %s. Please try again.\n", err.Error())
				}
				if attempts > 4 {
					fmt.Printf("Sorry, that was too many failed attempts(%d) to update balances. Try again next time.\n", attempts)
					break
				}
			}
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
func exitMessage(f *excelize.File) error {
	fmt.Println("Thanks for checking up on your finances!")
	netWorth, err := GetTotal(f)
	if err != nil {
		fmt.Printf("Error getting total... %s", err.Error())
		return err
	}
	delta := netWorth - startBalance
	if delta > 0 {
		fmt.Printf("Congrats! Your worth went up.\n")
		fmt.Printf("Your worth increased by %f from %f to %.2f\n", delta, startBalance, netWorth)
		err := persistTotal(f, netWorth)
		if err != nil {
			fmt.Printf("Error persisting total... %s", err.Error())
			return err
		}
	} else if delta < 0 {
		fmt.Printf("Oof... looks like your worth went down. Try saving more money\n")
		fmt.Printf("Your worth decreased by %f from %f to %.2f\n", delta, startBalance, netWorth)
		err := persistTotal(f, netWorth)
		if err != nil {
			fmt.Printf("Error persisting total... %s", err.Error())
			return err
		}
	} else {
		fmt.Printf("Your net worth did not change: %.2f\n", netWorth)
	}
	fmt.Println("Application terminating... come back for more features.")
	return nil
}

type netWorthEntry struct {
	Date             string  `json:"date"`
	NetWorth         float64 `json:"netWorth"`
	PreviousNetWorth float64 `json:"previousNetWorth"`
}

//Writes total net worth to a json file if it has a different value -> only records changes ie changeLog of total
func persistTotal(f *excelize.File, currentWorth float64) error {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var netWorthEntries []netWorthEntry
	err = json.Unmarshal(byteValue, &netWorthEntries)
	if err != nil {
		return err
	}

	previousEntry := netWorthEntries[len(netWorthEntries)-1]
	newEntry := netWorthEntry{
		Date:             time.Now().String(),
		NetWorth:         currentWorth,
		PreviousNetWorth: previousEntry.NetWorth,
	}

	netWorthEntries = append(netWorthEntries, newEntry)

	jsonData, err := json.Marshal(netWorthEntries)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
