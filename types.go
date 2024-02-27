package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount float64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type UpdateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Balance float64 `json:"balance"`
}

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Number int64 `json:"number"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(100000000)),
		CreatedAt: time.Now().UTC(),
	}
}

func UpdateAccount(firstName, lastName string, balance float64) *Account {
	return &Account{
		FirstName: firstName,
		LastName: lastName,
		Balance: balance,
	}
}