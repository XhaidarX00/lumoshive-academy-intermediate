package utils

import (
	"encoding/json"
	"errors"
	"project-voucher-team3/models"

	"time"
)

func ValidateVoucher(voucherValidateInput models.VoucherDTO, voucher models.Voucher) (models.ValidateVoucherResponse, error) {
	var result models.ValidateVoucherResponse

	if voucher.ID != 0 {
		return result, errors.New("voucher not found")
	}
	result.VoucherID = voucher.ID
	result.VoucherCode = voucher.VoucherCode

	if voucher.StartDate.Before(voucherValidateInput.FormatedTransactionDate) || voucher.EndDate.Before(voucherValidateInput.FormatedTransactionDate) {
		return result, errors.New("voucher is already out of date")
	} else {
		result.VoucherStatus = "Active"
	}

	if voucherValidateInput.TotalTransaction < voucher.MinPurchase {
		return result, errors.New("your total transaction must be greater than the minimum purchase")
	} else {
		result.TotalTransaction = voucherValidateInput.TotalTransaction
		result.TotalShippingCost = voucherValidateInput.TotalShippingCost
	}
	if voucher.PaymentMethod != "" && voucherValidateInput.PaymentMethod != voucher.PaymentMethod {
		return result, errors.New("payment method does not match")
	}
	if voucher.ApplicableAreas != "" {
		var areas []string
		err := json.Unmarshal([]byte(voucher.ApplicableAreas), &areas)
		if err != nil {
			return result, errors.New("invalid applicable areas format")
		}

		// Check if the input area is in the applicable areas list
		isValidArea := false
		for _, area := range areas {
			if area == voucherValidateInput.Area {
				isValidArea = true
				break
			}
		}
		if !isValidArea {
			return result, errors.New("area is not valid")
		}
	}

	benefitAmount := CalculateBenefit(voucherValidateInput.TotalTransaction, voucherValidateInput.TotalShippingCost, voucher.DiscountAmount, voucher.VoucherCategory)
	result.BenefitAmount = benefitAmount
	return result, nil
}

func isValidArea(applicableAreasJSON, inputArea string) bool {
	var areas []string
	if err := json.Unmarshal([]byte(applicableAreasJSON), &areas); err != nil {
		return false
	}
	for _, area := range areas {
		if area == inputArea {
			return true
		}
	}
	return false
}

func isVoucherActive(voucher models.Voucher, transactionDate time.Time) bool {
	return voucher.StartDate.Before(transactionDate) && voucher.EndDate.After(transactionDate)
}
