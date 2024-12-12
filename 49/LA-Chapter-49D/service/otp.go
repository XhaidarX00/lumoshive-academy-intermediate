package service

import (
	"be-golang-chapter-49/LA-Chapter-49D/config"
	"bytes"
	"fmt"
	"math/rand"
	"text/template"
	"time"

	"gopkg.in/gomail.v2"
)

// Struct untuk memuat data ke dalam template
type EmailData struct {
	OTP string
}

func GenerateOTP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000)) // Generate 6 digit OTP
	return otp
}

func SendOTPEmail(to string, otp string) error {
	// Load template HTML
	tmpl, err := template.ParseFiles("email/template.html")
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	// Data yang akan dimuat ke dalam template
	data := EmailData{OTP: otp}

	// Apply template dengan data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.AppConfig.FromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(config.AppConfig.SMTPHost, config.AppConfig.SMTPPort, config.AppConfig.SMTPUser, config.AppConfig.SMTPPassword)

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
