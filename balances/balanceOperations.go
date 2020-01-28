package balances

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var balanceMap = map[string]float64{"Checking": 0, "Savings - Meryll": 0, "Savings - New": 0, "Stocks - Plan": 0, "Stocks - Indv": 0, "Total": 0}

const balanceSheet = "balances/balances.xlsx"

func ReportBalances(detailedReport bool) (string, error) {
	f, err := excelize.OpenFile(balanceSheet)
	if err != nil {
		return "", err
	}
	//go through the balances and initialize their values
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

func initializeBalances(file *excelize.File) error {
	for key, _ := range balanceMap {
		keyPos, err := file.SearchSheet("Sheet1", key, true)
		if err != nil {
			return err
		}
		valuePos, err := getAdjacentCellPos(keyPos[0])
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

func getAdjacentCellPos(cell string) (string, error) {
	x, y, _ := excelize.CellNameToCoordinates(cell)
	x++
	return excelize.CoordinatesToCellName(x, y)
}
