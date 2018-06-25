package models

import (
	"fmt"
	"os"
)

func BillDueDates() (map[string]map[string]string, error) {
	finalPayload := make(map[string]map[string]string)
	xAxis := [5]string{"category", "Account", "Automated", "DueDate"}

	//Category	Account	Automated	DueDate
	spreadsheetId := os.Getenv("SHEETID")
	// columnID := MothToColumnDashboard[month]
	// columnIndex := ColumnToIndexDashboard[columnID]

	fmt.Printf("Spreadsheet id %s \n", spreadsheetId)
	readRange := "BillDueDates!D7:H19"

	fmt.Println(readRange)
	resp, err := SheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}

	fmt.Println(len(resp.Values))
	for index := 1; index < len(resp.Values); index++ {
		// fmt.Println(resp.Values[index][0].(string))
		payload := make(map[string]string)
		var billName string
		for rows := 1; rows < len(resp.Values[index]); rows++ {
			fmt.Println(resp.Values[1][rows].(string))
			payload[xAxis[rows-1]] = resp.Values[1][rows].(string)
			billName = resp.Values[index][0].(string)
		}
		finalPayload[billName] = payload
		fmt.Println(payload)
		fmt.Println()
	}

	return finalPayload, nil
}
