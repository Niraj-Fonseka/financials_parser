package models

import (
	"fmt"
	"log"
	"os"
)

func DashboardBillsList() []string {

	var billsList []string

	spreadsheetId := os.Getenv("SHEETID")

	fmt.Printf("Spreadsheet id %s \n", spreadsheetId)
	readRange := "Dashboard!B15:B22"
	resp, err := SheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.

			//billsList = append(billsList, string(row[0]))
			fmt.Printf("%s, %s\n", row[0].(string))
			billsList = append(billsList, row[0].(string))
		}
	}

	return billsList
}

type BillsMonth struct {
	bill   string `json:"bill"`
	amount string `json:"amount"`
}

func DashboardBillsListMonthly(month string) map[string]string {
	//var billsList []string
	payload := make(map[string]string)

	spreadsheetId := os.Getenv("SHEETID")

	fmt.Printf("Spreadsheet id %s \n", spreadsheetId)
	readRange := "Dashboard!B16:" + month + "22"

	fmt.Println(readRange)
	resp, err := SheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	fmt.Println(resp.Values[0][0].(string))
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {

		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.

			//billsList = append(billsList, string(row[0]))
			fmt.Printf("%s - %s \n", row[0].(string), row[1].(string))
			payload[row[0].(string)] = row[1].(string)
			//billsList = append(billsList, row[0].(string))
		}
	}

	return payload
}
