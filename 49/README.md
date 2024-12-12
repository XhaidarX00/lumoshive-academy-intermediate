# be-golang-chapter-49
this repo for golang chapter 49


### exampel use app.mailjet.com
```go
/*
This call sends a message to one recipient.
*/
package main

import (
	"fmt"
	"log"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func main() {
	// Inisialisasi client Mailjet
	mailjetClient := mailjet.NewMailjetClient("api-key", "secreat-key")

	// Menyiapkan data pesan
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "lumoshivetesting@gmail.com",
				Name:  "lumos",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "lumoshive.academy@gmail.com",
					Name:  "lumos_testing",
				},
			},
			Subject:  "OTP",
			TextPart: "ini adalah pesan otp",
		},
	}

	// Menyiapkan payload untuk pengiriman email
	messages := mailjet.MessagesV31{Info: messagesInfo}

	// Mengirim email
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil
	fmt.Printf("Data: %+v\n", res)
}
```