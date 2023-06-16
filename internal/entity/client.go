package entity

import "time"

type ClientRaw struct {
	ID                 string
	Document           string
	IsPrivate          string
	IsIncomplete       string
	LastPurchaseDate   *string
	AverageTicket      *string
	LastPurchaseTicket *string
	MostFrequentStore  *string
	LastPurchaseStore  *string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}

type Client struct {
	ID                 string
	Document           string
	DocumentType       string
	IsPrivate          bool
	IsIncomplete       bool
	LastPurchaseDate   *time.Time
	AverageTicket      float64
	LastPurchaseTicket float64
	MostFrequentStore  string
	LastPurchaseStore  string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}
