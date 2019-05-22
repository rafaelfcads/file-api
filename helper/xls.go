package helper

import (
	"C"
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/rafaelfcads/file-api/model"
)

func createHeader(sheet string, xlsx *excelize.File) {
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 1), "Customer name")
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 1), "Customer code")
	xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 1), "Currency")
}

func createHeaderSummary(sheet string, xlsx *excelize.File) {
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 4), "SUMMARY")
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 4), "TOTAL")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 5), "Overdue")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 6), "Credit")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 7), "Dispute")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 8), "Not Overdue")
}

func createHeaderRow(sheet string, xlsx *excelize.File) {
	const row = 10
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", row), "CUSTOMER")
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", row), "INVOICE #")
	xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", row), "REFERENCE #")
	xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", row), "BILLING DOC")
	xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", row), "PRODUCT")
	xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", row), "DISPUTE")
	xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", row), "ISSUED DATE")
	xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", row), "DUE DATE")
	xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", row), "TOTAL")
}

func createSheet(row int, name string, xlsx *excelize.File) {
	if row == 0 {
		xlsx.SetSheetName("Sheet1", name)
	}

	if row > 0 {
		xlsx.NewSheet(name)
	}
}

func JsonToXlsx(docs model.DocumentFile) (*bytes.Buffer, error) {

	xlsx := excelize.NewFile()
	const rowInit = 11

	for keyMain, docItem := range docs.Consolidates {
		sheet := docItem.SalesOrganization + "_" + docItem.Currency
		createSheet(keyMain, sheet, xlsx)
		createHeader(sheet, xlsx)
		xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 2), docItem.CustomerName)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 2), docItem.CustomerCode)
		xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 2), docItem.Currency)

		createHeaderSummary(sheet, xlsx)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 5), docItem.Overdue)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 6), docItem.Credit)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 7), docItem.Dispute)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 8), docItem.NotOverdue)

		for key, doc := range docItem.Docs {
			createHeaderRow(sheet, xlsx)
			xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", rowInit+key), doc.CustomerCode)
			xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", rowInit+key), doc.Number)
			xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", rowInit+key), doc.ReferenceNumber)
			xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", rowInit+key), doc.BillingNumber)
			xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", rowInit+key), doc.Division)
			xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", rowInit+key), doc.IsDispute)
			xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", rowInit+key), doc.IssuedDate)
			xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", rowInit+key), doc.DueDate)
			xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", rowInit+key), doc.TotalAmount)
		}
	}

	return xlsx.WriteToBuffer()
}
