package middlewares

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
)
var OTP string

func VerifyEmail(email string) string{
	OTP=""
	pass := os.Getenv("pass")
	//company mail
	from := os.Getenv("email")
	fmt.Println(pass)

	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//generate a otp
	OTP = generateOTP()
	// send the otp
	
	message := []byte(OTP)

	// Authentication.
	auth := smtp.PlainAuth("", from, pass, smtpHost)
	fmt.Println(auth)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully!")
	return OTP
}

func generateOTP() string {
	for i:=0;i<5;i++{
		x := rand.Intn(10)
		OTP = OTP + strconv.Itoa(x)
	}
	return OTP
}
