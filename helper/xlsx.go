package helper

import (
	"C"
	"bytes"
	"fmt"
	_ "image/png"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/rafaelfcads/file-api/model"
)

func createHeader(sheet string, xlsx *excelize.File) {
	const config = `{"x_scale": 0.5, "y_scale": 0.5, "hyperlink": "#Sheet2!D8", "hyperlink_type": "Location"}`
	errPicture := xlsx.AddPicture(sheet, "A1", "./image/icon-embraer.png", config)
	if errPicture != nil {
        fmt.Println(errPicture)
	}
	errColWidth := xlsx.SetColWidth(sheet, "A", "I", 20)
	if errColWidth != nil {
        fmt.Println(errColWidth)
	}
	
	headerCellStyle := `{"font":{"bold":true, "color":"#FFFFFF"},"fill":{"type":"pattern","color":["#100690"],"pattern":1}}`
	headerStyle, errHeaderStyle := xlsx.NewStyle(headerCellStyle)
	if errHeaderStyle != nil {
		fmt.Println(errHeaderStyle)
	}

	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 5), "Customer name")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 5), fmt.Sprintf("A%d", 5), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 5), "Customer code")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 5),  fmt.Sprintf("B%d", 5), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 5), "Currency")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("C%d", 5), fmt.Sprintf("C%d", 5), headerStyle)
}

func createHeaderSummary(sheet string, xlsx *excelize.File) {
	cellStyle := `{"font":{"bold":true, "color":"#9A0511"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1},{"type":"right","color":"#000000","style":1}]}`
	style, errNewStyle := xlsx.NewStyle(cellStyle)
	if errNewStyle != nil {
		fmt.Println(errNewStyle)
	}

	headerCellStyle := `{"font":{"bold":true, "color":"#FFFFFF"},"fill":{"type":"pattern","color":["#100690"],"pattern":1}}`
	headerStyle, errHeaderStyle := xlsx.NewStyle(headerCellStyle)
	if errHeaderStyle != nil {
		fmt.Println(errHeaderStyle)
	}

	borderCellStyle := `{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1},{"type":"right","color":"#000000","style":1}]}`
	borderStyle, err := xlsx.NewStyle(borderCellStyle)
	if err != nil {
		fmt.Println(err)
	}
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 8), "SUMMARY")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 8), fmt.Sprintf("A%d", 8), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 8), "TOTAL")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 8), fmt.Sprintf("B%d", 8), headerStyle)

	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 9), "Overdue")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 9), fmt.Sprintf("A%d", 9), style)
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 10), "Credit")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 10), fmt.Sprintf("A%d", 10), borderStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 11), "Dispute")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 11), fmt.Sprintf("A%d", 11), borderStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 12), "Not Overdue")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 12), fmt.Sprintf("A%d", 12), borderStyle)
}

func createHeaderRow(sheet string, xlsx *excelize.File, len int) {
	const row = 14
	headerCellStyle := `{"font":{"bold":true, "color":"#FFFFFF"},"fill":{"type":"pattern","color":["#100690"],"pattern":1}}`
	headerStyle, errHeaderStyle := xlsx.NewStyle(headerCellStyle)
	if errHeaderStyle != nil {
		fmt.Println(errHeaderStyle)
	}
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", row), "CUSTOMER")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", row), "INVOICE #")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", row), "REFERENCE #")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", row), "BILLING DOC")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("D%d", row), fmt.Sprintf("D%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", row), "PRODUCT")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("E%d", row), fmt.Sprintf("E%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", row), "DISPUTE")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("F%d", row), fmt.Sprintf("F%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", row), "ISSUED DATE")
	xlsx.SetCellStyle(sheet,fmt.Sprintf("G%d", row), fmt.Sprintf("G%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", row), "DUE DATE")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("H%d", row), fmt.Sprintf("H%d", row), headerStyle)
	xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", row), "TOTAL")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("I%d", row), fmt.Sprintf("I%d", row), headerStyle)
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
	const rowInit = 15
	borderCellStyle := `{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1},{"type":"right","color":"#000000","style":1}]}`
	borderStyle, err := xlsx.NewStyle(borderCellStyle)
	if err != nil {
		fmt.Println(err)
	}

	for keyMain, docItem := range docs.Consolidates {
		sheet := docItem.SalesOrganization + "_" + docItem.Currency
		createSheet(keyMain, sheet, xlsx)
		createHeader(sheet, xlsx)
		xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 6), docItem.CustomerName)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 6), fmt.Sprintf("A%d", 6), borderStyle)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 6), docItem.CustomerCode)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 6), fmt.Sprintf("B%d", 6), borderStyle)
		xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 6), docItem.Currency)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("C%d", 6), fmt.Sprintf("C%d", 6), borderStyle)

		createHeaderSummary(sheet, xlsx)
		cellStyle := `{"font":{"bold":true, "color":"#FF0000"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1},{"type":"right","color":"#000000","style":1}]}`
		style, errNewStyle := xlsx.NewStyle(cellStyle)
		if errNewStyle != nil {
			fmt.Println(errNewStyle)
		}
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 9), docItem.Overdue)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 9), fmt.Sprintf("B%d", 9), style)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 10), docItem.Credit)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 10), fmt.Sprintf("B%d", 10), borderStyle)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 11), docItem.Dispute)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 11), fmt.Sprintf("B%d", 11), borderStyle)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 12), docItem.NotOverdue)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 12), fmt.Sprintf("B%d", 12), borderStyle)

		for key, doc := range docItem.Docs {
			createHeaderRow(sheet, xlsx, len(docItem.Docs))
			xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", rowInit+key), doc.CustomerCode)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", rowInit+key), fmt.Sprintf("A%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", rowInit+key), doc.Number)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", rowInit+key), fmt.Sprintf("B%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", rowInit+key), doc.ReferenceNumber)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("C%d", rowInit+key), fmt.Sprintf("C%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", rowInit+key), doc.BillingNumber)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("D%d", rowInit+key), fmt.Sprintf("D%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", rowInit+key), doc.Division)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("E%d", rowInit+key), fmt.Sprintf("E%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", rowInit+key), doc.Dispute)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("F%d", rowInit+key), fmt.Sprintf("F%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", rowInit+key), doc.IssuedDate)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("G%d", rowInit+key), fmt.Sprintf("G%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", rowInit+key), doc.DueDate)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("H%d", rowInit+key), fmt.Sprintf("H%d", rowInit+key), borderStyle)
			xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", rowInit+key), doc.TotalAmount)
			xlsx.SetCellStyle(sheet, fmt.Sprintf("I%d", rowInit+key), fmt.Sprintf("I%d", rowInit+key), borderStyle)
		}
	}

	return xlsx.WriteToBuffer()
}
