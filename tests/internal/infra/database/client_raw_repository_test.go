package database

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestCreateClientRawDataWhenSuccessful(t *testing.T) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@db:5432/dataloader?search_path=dataloader_test")
	println(conn)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer conn.Close()

	clientRepository := database.NewClientRawRepository(conn)

	clients := []*entity.ClientRaw{
		{
			ID:                 uuid.New().String(),
			Document:           "1234567890",
			IsPrivate:          "0",
			IsIncomplete:       "0",
			LastPurchaseDate:   nil,
			AverageTicket:      nil,
			LastPurchaseTicket: nil,
			MostFrequentStore:  nil,
			LastPurchaseStore:  nil,
			Status:             "Active",
			CreatedAt:          time.Now().Format(time.RFC3339),
		},
	}

	err = clientRepository.Create(clients)
	fmt.Println(err)
	assert.NoError(t, err)

	query := `TRUNCATE TABLE raw_client_data;`
	_, err = conn.Exec(context.Background(), query)

}
