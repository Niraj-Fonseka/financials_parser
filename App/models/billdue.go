package models

import (
	"os"
)

func BillDueDates() (map[string]map[string]string, error) {
	finalPayload := make(map[string]map[string]string)
	xAxis := [5]string{"category", "Account", "Automated", "DueDate"}

	spreadsheetId := os.Getenv("SHEETID")

	readRange := "BillDueDates!D7:H19"
	resp, err := SheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		return nil, err
	}

	for index := 1; index < len(resp.Values); index++ {
		payload := make(map[string]string)
		var billName string
		for rows := 1; rows < len(resp.Values[index]); rows++ {

			payload[xAxis[rows-1]] = resp.Values[1][rows].(string) //category as the key and the values
			billName = resp.Values[index][0].(string)              // extrating out the name of the bill
		}
		finalPayload[billName] = payload //using the bill name as the key and the map of categories and values as the value
	}

	return finalPayload, nil
}
