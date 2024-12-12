package handler

import (
	"be-golang-chapter-49/LA-Chapter-49D/service"
	"fmt"
	"net/http"
)

func OtpHandler(w http.ResponseWriter, r *http.Request) {
	to := r.URL.Query().Get("email")
	if to == "" {
		http.Error(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	otp := service.GenerateOTP()

	if err := service.SendOTPEmail(to, otp); err != nil {
		http.Error(w, fmt.Sprintf("Failed to send email: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OTP sent to %s", to)
}
