package operations

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var balanceMap = map[string]float64{"Checking": 0, "Savings - Meryll": 0, "Savings - New": 0, "Stocks - Plan": 0, "Stocks - Indv": 0, "Uber Available Credit": 0, "C1 Available Credit": 0, "Liquidity": 0}

var accountsKey = map[string]string{"chck": "Checking", "svngM": "Savings - Meryll", "svngNew": "Savings - New", "stcksP": "Stocks - Plan", "stcksI": "Stocks - Indv"}

const lowBalanceThreshold = 500.00

type File *excelize.File

//Reports out balances
func ReportBalances(detailedReport bool, f *excelize.File) (string, error) {
	initializeBalances(f)
	var sb strings.Builder

	//Just give checking account information
	if !detailedReport {
		balance := fmt.Sprintf("%.2f", balanceMap["Checking"])
		ret := "Checkings balances is: " + balance
		return ret, nil
	}
	//Give a full balance report
	for key, value := range balanceMap {
		balance := fmt.Sprintf("%.2f", value)
		sb.WriteString(key + " balance is: " + balance + "\n")
	}

	return sb.String(), nil
}

//Retrieves the current total balance
func GetTotal(f *excelize.File) (float64, error) {
	initializeBalances(f)

	var total float64
	for _, value := range balanceMap {
		total += value
	}
	ret := fmt.Sprintf("%.2f", total)

	return strconv.ParseFloat(ret, 64)
}

func GetBalanceMap(f *excelize.File) map[string]float64 {
	initializeBalances(f)
	return balanceMap
}

//updates all of the balances selected by the user and reinitializes map after
func UpdateBalances(f *excelize.File) error {
	fmt.Print("\n\nPlease enter which accounts you would like to update separated by commas... here is the legend for account codes: \n")

	for key, value := range accountsKey {
		fmt.Println(value + ": " + key)
	}

	var accounts string
	fmt.Scanf("%s", &accounts)

	accountsList := strings.Split(accounts, ",")
	for _, accountKey := range accountsList {
		account := accountsKey[accountKey]
		if account != "" {
			err := updateBalance(f, account)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Invalid account key: " + accountKey)
			return errors.New("invalid account key provided")
		}

	}
	initializeBalances(f)
	return nil
}

//Updates the balance of the provided account
func updateBalance(f *excelize.File, account string) error {
	cell, err := f.SearchSheet("Sheet1", account, true)
	if err != nil {
		return err
	}
	curPos, err := GetAdjacentCellPos(cell[0], 1)
	if err != nil {
		return err
	}
	fmt.Println("Please input the new balance for " + account + " account: ")
	var newBalance float64
	_, err = fmt.Scanf("%f", &newBalance)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	f.SetCellValue("Sheet1", curPos, newBalance)
	f.Save()
	return nil
}

//fills in the map of various balances
func initializeBalances(file *excelize.File) error {
	for key, _ := range balanceMap {
		keyPos, err := file.SearchSheet("Sheet1", key, true)
		if err != nil {
			return err
		}
		valuePos, err := GetAdjacentCellPos(keyPos[0], 1)
		if err != nil {
			return err
		}
		valueString, err := file.GetCellValue("Sheet1", valuePos)
		if err != nil {
			return err
		}
		value, err := strconv.ParseFloat(valueString, 32)
		if err != nil {
			return err
		}
		balanceMap[key] = value
	}
	return nil
}

func GetAdjacentCellPos(cell string, shift int) (string, error) {
	x, y, _ := excelize.CellNameToCoordinates(cell)
	x += shift
	return excelize.CoordinatesToCellName(x, y)
}

func GetBelowCellPos(cell string) (string, error) {
	x, y, _ := excelize.CellNameToCoordinates(cell)
	y++
	return excelize.CoordinatesToCellName(x, y)
}
