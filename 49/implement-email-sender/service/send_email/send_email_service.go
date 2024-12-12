package sendemail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	"github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
)

type EmailSenderInterface interface {
	GenerateOTP() string
	SendEmail(receiped, otp string) error
	SendEmailRegis(recipient, name string) error
}

type EmailSender struct {
	Log *zap.Logger
}

type EmailData struct {
	OTP string
}

type EmailDataGreat struct {
	Name string
	Year int
}

func NewEmailSenderService(log *zap.Logger) EmailSenderInterface {
	return &EmailSender{Log: log}
}

// var APIKey = "mlsn.7b9801ccf7f5ddee350bab5536d83e0106f35fa0c6eb09deb6fcd2874059d032"

func (s *EmailSender) GenerateOTP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000)) // Generate 6 digit OTP
	return otp
}

// SendOTPEmail mengirimkan email OTP melalui MailerSend
func (s *EmailSender) SendEmail(receiped, otp string) error {
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

	// Konfigurasi MailerSend
	apiKey := os.Getenv("MAILERSEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("MAILERSEND_API_KEY is not set")
	}

	fromEmail := os.Getenv("MAILERSEND_FROM_EMAIL")
	if fromEmail == "" {
		return fmt.Errorf("MAILERSEND_FROM_EMAIL is not set")
	}

	ms := mailersend.NewMailersend(apiKey)

	// Konfigurasi konten email
	subject := "Your OTP Code"
	htmlContent := body.String()

	from := mailersend.From{
		Name:  "Darmi Ecommers",
		Email: fromEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Recipient",
			Email: receiped,
		},
	}

	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(htmlContent)

	// Kirim email
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	_, err = ms.Email.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}

func (s *EmailSender) SendEmailRegis(recipient, name string) error {
	// Load template HTML
	tmplPath := "email/regis.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to load email template from %s: %v", tmplPath, err)
	}

	// Data yang akan dimuat ke dalam template
	data := EmailDataGreat{
		Name: name,
		Year: time.Now().Year(),
	}

	// Apply template dengan data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	// Konfigurasi MailerSend API
	apiKey := os.Getenv("MAILERSEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("MAILERSEND_API_KEY is not set in environment variables")
	}

	fromEmail := os.Getenv("MAILERSEND_FROM_EMAIL")
	if fromEmail == "" {
		return fmt.Errorf("MAILERSEND_FROM_EMAIL is not set in environment variables")
	}

	ms := mailersend.NewMailersend(apiKey)

	// Konfigurasi konten email
	subject := "Welcome!"
	htmlContent := body.String()

	from := mailersend.From{
		Name:  "Darmi Ecommers",
		Email: fromEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  name,
			Email: recipient,
		},
	}

	// Membuat pesan email
	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(htmlContent)

	// Kirim email dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = ms.Email.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send email to %s: %v", recipient, err)
	}

	return nil
}
