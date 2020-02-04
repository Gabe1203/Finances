package operations

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type payments []*struct {
	Name     string
	Category string
	Amount   float64
}

/* Updates checking balance
takes in user input of payments as well as d\goes through scheduled payments
returns a map of payments and their category to be calculated into new total and cached
*/
func makePayments(f *excelize.File) map[string]float64 {

	return map[string]float64{}
}

/*
Runs through the A column and gets the payment name
func returns map of spend category to amount
if the last logged on date is before the pay date
and the current date is after then include that transaction
if auto pay is on otherwise ask the user
if the cost is v then ask user what the cost was
*/
func getScheduledPayments(f *excelize.File) (payments, error) {
	startingCell := "A2"
	var scheduledPayments payments
	index := 0
	for {
		name, err := f.GetCellValue("Sheet1", startingCell)
		if err != nil {
			return scheduledPayments, err
		}
		if name == "" {
			break
		}
		//TODO: check date with current date to see if we need to include this payment
		scheduledPayments[index].Name = name
		amountCell, err := GetAdjacentCellPos(startingCell, 1)
		if err != nil {
			return scheduledPayments, err
		}
		amountStr, err := f.GetCellValue("Sheet1", amountCell)
		if err != nil {
			return scheduledPayments, err
		}
		var amount float64
		if amountStr == "v" {
			fmt.Println("Please enter how much you paid for: " + name)

			for {
				_, err := fmt.Scanf("%.2f", &amount)
				if err != nil {
					fmt.Println("Invalid input, please try again.")
				} else {
					break
				}
			}
		} else {
			amount, err = strconv.ParseFloat(amountStr, 64)
			if err != nil {
				return scheduledPayments, err
			}
		}
		scheduledPayments[index].Amount = amount

		categoryCell, err := GetAdjacentCellPos(startingCell, 5)
		if err != nil {
			return scheduledPayments, err
		}
		category, err := f.GetCellValue("Sheet1", categoryCell)
		if err != nil {
			return scheduledPayments, err
		}
		//TODO: check forr valid cat
		scheduledPayments[index].Category = category

		index++
		startingCell, err = GetBelowCellPos(startingCell)
		if err != nil {
			return scheduledPayments, err
		}

	}
	return scheduledPayments, nil
}
