package balances

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var balanceMap = map[string]float64{"Checking": 0, "Savings - Meryll": 0, "Savings - New": 0, "Stocks - Plan": 0, "Stocks - Indv": 0, "Total": 0}

type File *excelize.File

//Reports out balances
func ReportBalances(detailedReport bool, f *excelize.File) (string, error) {
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

//updates all of the balances and reinitializes map after
func UpdateBalances(f *excelize.File) error {
	err := updateCheckingsBalance(f)
	if err != nil {
		return err
	}
	initializeBalances(f)
	return nil
}

//Connect to Bank of America API to update the value
//onlineId1
//passcode1
//Idea: use selenium to log into the website
//sign in link: https://staticweb.bankofamerica.com/cavmwebbactouch/common/index.html#home?app=signonv2
//btCustomOnlineId
//TODO: change this from manual update to update with web scraping
func updateCheckingsBalance(f *excelize.File) error {
	cell, err := f.SearchSheet("Sheet1", "Checking", true)
	if err != nil {
		return err
	}
	curPos, err := getAdjacentCellPos(cell[0])
	if err != nil {
		return err
	}
	fmt.Println("Please input the new balance: ")
	// reader := bufio.NewReader(os.Stdin)
	// newValue, err := reader.ReadString('\n')
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
