package operations

import "github.com/360EntSecGroup-Skylar/excelize"

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
func getScheduledPayments(f *excelize.File) {

}
