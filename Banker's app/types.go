package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TransferReq struct {
	ToAccount string `json:"to_account"`
	Amount    uint64 `json:"amount"`
}

type LoginReq struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AccountReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Account struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    uint64    `json:"number"`
	Balance   uint64    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstname string, lastname string, email string, password string) *Account {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &Account{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: firstname,
		LastName:  lastname,
		Number:    uint64(rand.Intn(1000)),
		Balance:   uint64(rand.Intn(1000)),
		CreatedAt: time.Now().UTC(),
	}
}
