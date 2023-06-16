package entity

import (
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClientRaw(t *testing.T) {
	client := entity.ClientRaw{
		ID:                 "123",
		Document:           "123456789",
		IsPrivate:          "true",
		IsIncomplete:       "false",
		LastPurchaseDate:   nil,
		AverageTicket:      nil,
		LastPurchaseTicket: nil,
		MostFrequentStore:  nil,
		LastPurchaseStore:  nil,
		Status:             "active",
		CreatedAt:          "2023-01-01",
	}

	assert.Equal(t, "123", client.ID)
	assert.Equal(t, "123456789", client.Document)
	assert.Equal(t, "true", client.IsPrivate)
	assert.Equal(t, "false", client.IsIncomplete)
	assert.Nil(t, nil, client.LastPurchaseDate)
	assert.Nil(t, nil, client.AverageTicket)
	assert.Nil(t, nil, client.LastPurchaseTicket)
	assert.Nil(t, nil, client.MostFrequentStore)
	assert.Nil(t, nil, client.LastPurchaseStore)
	assert.Equal(t, "active", client.Status)
	assert.Equal(t, "2023-01-01", client.CreatedAt)
}

func parseLastPurchaseDate(date string) (*time.Time, error) {

	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

func TestClient(t *testing.T) {
	parsedValue, _ := parseLastPurchaseDate("2006-01-02")

	client := entity.Client{
		ID:                 "123",
		Document:           "123456789",
		IsPrivate:          true,
		IsIncomplete:       false,
		LastPurchaseDate:   parsedValue,
		AverageTicket:      100.50,
		LastPurchaseTicket: 150.75,
		MostFrequentStore:  "79.379.491/0001-83",
		LastPurchaseStore:  "79.379.491/0001-83",
		Status:             "active",
		CreatedAt:          "2023-01-01",
	}

	assert.Equal(t, "123", client.ID)
	assert.Equal(t, "123456789", client.Document)
	assert.Equal(t, true, client.IsPrivate)
	assert.Equal(t, false, client.IsIncomplete)
	assert.Equal(t, parsedValue, client.LastPurchaseDate)
	assert.Equal(t, 100.50, client.AverageTicket)
	assert.Equal(t, 150.75, client.LastPurchaseTicket)
	assert.Equal(t, "79.379.491/0001-83", client.MostFrequentStore)
	assert.Equal(t, "79.379.491/0001-83", client.LastPurchaseStore)
	assert.Equal(t, "active", client.Status)
	assert.Equal(t, "2023-01-01", client.CreatedAt)
}
