package utils

import (
	"fmt"
	"log"
	"project_auth_jwt/models"

	"github.com/xuri/excelize/v2"
)

func WriteReportsToExcel(filePath string, reports []models.Report) error {
	// Create a new Excel file
	file := excelize.NewFile()
	sheetName := "Reports"
	file.NewSheet(sheetName)
	file.DeleteSheet("Sheet1") // Remove the default sheet

	// Define headers for the Excel sheet
	headers := []string{
		"ID", "Customer Name", "Product Name", "Quantity", "Total Amount",
		"Payment Method", "Shipping Address", "Status", "Created At",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		file.SetCellValue(sheetName, cell, header)
	}

	// Write data rows
	for rowIdx, report := range reports {
		rowNumber := rowIdx + 2 // Start from the second row
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNumber), report.ID)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNumber), report.CustomerName)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNumber), report.ProductName)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNumber), report.Quantity)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNumber), fmt.Sprintf("%.2f", report.TotalAmount))
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNumber), report.PaymentMethod)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNumber), report.ShippingAddress)
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNumber), report.Status)
		file.SetCellValue(sheetName, fmt.Sprintf("I%d", rowNumber), report.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// Auto-adjust column widths for better readability
	for col := 0; col < len(headers); col++ {
		column := string(rune('A' + col))
		err := file.SetColWidth(sheetName, column, column, 20)
		if err != nil {
			return fmt.Errorf("failed to set column width: %v", err)
		}
	}

	// Save the file
	if err := file.SaveAs(filePath); err != nil {
		return fmt.Errorf("failed to save Excel file: %v", err)
	}

	log.Println("Export success")

	return nil
}
