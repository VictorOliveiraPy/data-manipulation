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

func TestCreateClientDataWhenSuccessful(t *testing.T) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@db:5432/dataloader?search_path=dataloader_test")
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer conn.Close()

	clientRepository := database.NewClientRepository(conn)

	clients := []*entity.Client{
		{
			ID:                 uuid.New().String(),
			Document:           "1234567890",
			IsPrivate:          false,
			IsIncomplete:       true,
			LastPurchaseDate:   nil,
			AverageTicket:      19.0,
			LastPurchaseTicket: 20.9,
			MostFrequentStore:  "79.379.491/0001-83",
			LastPurchaseStore:  "79.379.491/0001-83",
			Status:             "Active",
			CreatedAt:          time.Now().Format(time.RFC3339),
		},
	}

	err = clientRepository.Create(clients)
	fmt.Println(err)
	assert.NoError(t, err)

	query := `TRUNCATE TABLE client_data;`
	_, err = conn.Exec(context.Background(), query)

}
