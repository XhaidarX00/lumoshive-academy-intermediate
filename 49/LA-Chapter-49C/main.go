// send email standar
package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "youremail@gmail.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is a simple email using gomail.")

	// // Menetapkan subjek email
	// m.SetHeader("Subject", "Welcome to Our Service")

	// // Menetapkan isi email dengan HTML
	// m.SetBody("text/html", `
	// 		<html>
	// 		<body>
	// 			<h1>Welcome to Our Service</h1>
	// 			<p>Dear User,</p>
	// 			<p>Thank you for registering with our service. We are excited to have you on board!</p>
	// 			<p>Please <a href="https://example.com/activate">click here</a> to activate your account.</p>
	// 			<p>Best regards,<br/>The Team</p>
	// 		</body>
	// 		</html>
	// 	`)

	// // Menetapkan subjek email
	// m.SetHeader("Subject", "Invoice for Your Purchase")

	// // Menetapkan isi email
	// m.SetBody("text/plain", "Dear Customer,\n\nPlease find attached the invoice for your recent purchase.\n\nThank you for your business!")

	// // Menambahkan lampiran dokumen
	// m.Attach("./invoice.pdf")

	// // Menetapkan penerima email (banyak penerima)
	// m.SetHeader("To", "recipient1@example.com", "recipient2@example.com", "recipient3@example.com")

	// // Menetapkan subjek email
	// m.SetHeader("Subject", "Important Update")

	// // Menetapkan isi email
	// m.SetBody("text/plain", "Dear All,\n\nThis is an important update. Please read it carefully.\n\nBest regards,\nYour Name")

	d := gomail.NewDialer("smtp.gmail.com", 587, "youremail@gmail.com", "yourpassword")

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent successfully!")
}
