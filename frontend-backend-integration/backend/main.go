package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v76"
	_ "github.com/stripe/stripe-go/v76"
	_ "github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func main() {
	stripe.Key = "sk_test_51OBdmnSBYMn4HDSRhQf13fFGrat7wiQoplJOWpVpS6tLZBqxBg5FUGbxwUQs03weXZYJ5192QqbYICSKkFkhIAc200LEv3pM5L"
	http.HandleFunc("/create-payment-intent",handleCreatePaymentIntent)
	http.HandleFunc("/health",handleHealth)
	fmt.Println("Listening at localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242",nil)
	if err!=nil{
		log.Fatal(err)
	}
}
func handleCreatePaymentIntent(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST"{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City string `json:"city"`
		State string `json:"state"`
		Zip string `json:"zip"`
		Country string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	params := &stripe.PaymentIntentParams{
		Amount:stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency:stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent,err := paymentintent.New(params)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	fmt.Println(paymentIntent.ClientSecret)
}
func handleHealth(w http.ResponseWriter,r *http.Request){   //writing in the response and returning it to response
	fmt.Println("health is ok")
}
func calculateOrderAmount(productId string) int64{
	switch productId {
	case "Forever Pants":
		return 2300
	case "Forever Shirt":
		return 8900
	case "Forever Shorts":
		return 1200
	}
	return 0
}