package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"reflect"
	"strings"
)

type Details struct {
	Pass    string   `json:"pass"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message string   `json:"msg"`
}

// Get data from html
var html_file = "render.html"

func SendEmailDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(html_file)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, html_file, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetSendDetails(w http.ResponseWriter, r *http.Request) {
	pass := r.FormValue("pass")
	from := r.FormValue("from")
	toRaw := r.FormValue("to")
	msg := r.FormValue("msg")

	to := strings.FieldsFunc(toRaw, func(r rune) bool {
		return r == ',' || r == '\n'
	})

	for i := range to {
		to[i] = strings.TrimSpace(to[i])
	}

	fmt.Println("to[]:",to,reflect.TypeOf(to))

	fmt.Println(pass, from)
	fmt.Println(toRaw, msg)
	fmt.Println(reflect.TypeOf(toRaw))
	fmt.Println(reflect.TypeOf(from))
	userDetails := Details{
		pass,
		from,
		to,
		msg,
	}
	err := SendMail(&userDetails)

	if err != nil {
		log.Fatal(err)
	}

}

func UserEmailData(w http.ResponseWriter, r *http.Request) {
	var userDetails Details
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

	message := []byte(d.Message)

	// Sending email.

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, d.From, d.To, message)
	if err != nil {

		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}
