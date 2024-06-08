package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	_"time"
)

type Details struct {
	Pass    string   `json:"pass"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message []byte   `json:"msg"`
}

func UserEmailData(w http.ResponseWriter, r *http.Request) {
	var userDetails Details
	// time.Sleep(time.Second*3)
	json.NewDecoder(r.Body).Decode(&userDetails)
	fmt.Println("DETAILS: ")
	fmt.Println(userDetails)
	fmt.Println("---------------")
	json.NewEncoder(w).Encode(userDetails)
	err := SendMail(&userDetails)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Email Sent Successfully!")
}
//Send mail 

func SendMail(d *Details) error {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"


	// Authentication.
	auth := smtp.PlainAuth("", d.From, d.Pass, smtpHost)
	fmt.Println(auth)

	// Sending email.

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, d.From,d.To, d.Message)
	if err != nil {

		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}
