package utils

import (
	"fmt"
	"project_auth_jwt/models"

	"github.com/xuri/excelize/v2"
)

func WriteOrdersToExcel(filePath string, orders []models.Order) error {
	// Create a new Excel file
	file := excelize.NewFile()
	sheetName := "Orders"
	file.NewSheet(sheetName)
	file.DeleteSheet("Sheet1") // Remove the default sheet

	// Define headers for the Excel sheet
	headers := []string{"ID", "Customer Name", "Product Name", "Quantity", "Price", "Created At"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		file.SetCellValue(sheetName, cell, header)
	}

	// Write data rows
	for rowIdx, order := range orders {
		rowNumber := rowIdx + 2 // Start from the second row
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNumber), order.ID)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNumber), order.UserID)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNumber), order.TotalAmount)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNumber), order.PaymentMethod)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNumber), order.ShippingAddress)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNumber), order.Status)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNumber), order.CreatedAt.Format("2006-01-02 15:04:05"))
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNumber), order.UpdatedAt.Format("2006-01-02 15:04:05"))
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

	return nil
}
