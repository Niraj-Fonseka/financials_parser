package models

import (
	"errors"
	"fmt"
	"os"
)

func MonthlySpeditures(month string) (map[string]string, error) {
	//var billsList []string
	payload := make(map[string]string)

	spreadsheetId := os.Getenv("SHEETID")
	spreadsheetSheetName := MonthToAbbrMonthly[month]

	if spreadsheetSheetName == "" {
		err := errors.New("Not a valid month")
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Spreadsheet id %s \n", spreadsheetId)
	readRange := spreadsheetSheetName + "!J2:K22"

	fmt.Println(readRange)
	resp, err := SheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Values[0][0].(string))
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {

		fmt.Println(resp.Values)
		for _, row := range resp.Values {

			payload[row[0].(string)] = row[1].(string)
		}
	}

	return payload, nil
}
