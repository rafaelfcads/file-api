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
	format := `{"table_name":"createHeaderTable","table_style":"TableStyleMedium2"}`
	err := xlsx.AddTable(sheet, "A5", "C6", format)
	if err != nil {
        fmt.Println(err)
	}
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 5), "Customer name")
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 5), "Customer code")
	xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 5), "Currency")
}

func createHeaderSummary(sheet string, xlsx *excelize.File) {
	format := `{"table_name":"createHeaderSummaryTable","table_style":"TableStyleMedium2"}`
	err := xlsx.AddTable(sheet, "A8", "B12", format)
	if err != nil {
        fmt.Println(err)
	}

	cellStyle := `{"font":{"bold":true, "family":"Berlin Sans FB Demi","color":"#FF0000"}}`
	style, errNewStyle := xlsx.NewStyle(cellStyle)
	if errNewStyle != nil {
		fmt.Println(errNewStyle)
	}

	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 8), "SUMMARY")
	xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 8), "TOTAL")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 9), "Overdue")
	xlsx.SetCellStyle(sheet, fmt.Sprintf("A%d", 9), fmt.Sprintf("A%d", 9), style)
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 10), "Credit")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 11), "Dispute")
	xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 12), "Not Overdue")
}

func createHeaderRow(sheet string, xlsx *excelize.File, len int) {
	const row = 14
	format := `{"table_name":"createHeaderRowTable","table_style":"TableStyleMedium2"}`
	err := xlsx.AddTable(sheet, "A14", fmt.Sprintf("I%d", row + len), format)
	if err != nil {
        fmt.Println(err)
	}
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
	const rowInit = 15

	for keyMain, docItem := range docs.Consolidates {
		sheet := docItem.SalesOrganization + "_" + docItem.Currency
		createSheet(keyMain, sheet, xlsx)
		createHeader(sheet, xlsx)
		xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", 6), docItem.CustomerName)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 6), docItem.CustomerCode)
		xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", 6), docItem.Currency)

		createHeaderSummary(sheet, xlsx)
		cellStyle := `{"font":{"bold":true, "family":"Berlin Sans FB Demi","color":"#FF0000"}}`
		style, errNewStyle := xlsx.NewStyle(cellStyle)
		if errNewStyle != nil {
			fmt.Println(errNewStyle)
		}
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 9), docItem.Overdue)
		xlsx.SetCellStyle(sheet, fmt.Sprintf("B%d", 9), fmt.Sprintf("B%d", 9), style)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 10), docItem.Credit)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 11), docItem.Dispute)
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", 12), docItem.NotOverdue)

		for key, doc := range docItem.Docs {
			createHeaderRow(sheet, xlsx, len(docItem.Docs))
			xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", rowInit+key), doc.CustomerCode)
			xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", rowInit+key), doc.Number)
			xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", rowInit+key), doc.ReferenceNumber)
			xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", rowInit+key), doc.BillingNumber)
			xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", rowInit+key), doc.Division)
			xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", rowInit+key), doc.Dispute)
			xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", rowInit+key), doc.IssuedDate)
			xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", rowInit+key), doc.DueDate)
			xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", rowInit+key), doc.TotalAmount)
		}
	}
	err := xlsx.SaveAs("./Booooooo.xlsx")
	fmt.Println("Salvou!!!!")
    if err != nil {
        fmt.Println(err)
    }
	return xlsx.WriteToBuffer()
}
