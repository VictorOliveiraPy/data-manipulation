package database

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClientRepository struct {
	Db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{Db: db}
}

func (repository *ClientRepository) Create(clients []*entity.Client) error {
	const insertQuery = `INSERT INTO client_data (id, document, document_type, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	batch := &pgx.Batch{}
	for _, client := range clients {
		batch.Queue(
			insertQuery,
			client.ID,
			client.Document,
			client.DocumentType,
			client.IsPrivate,
			client.IsIncomplete,
			client.LastPurchaseDate,
			client.AverageTicket,
			client.LastPurchaseTicket,
			client.MostFrequentStore,
			client.LastPurchaseStore,
			client.Status,
			client.CreatedAt,
		)

	}

	br := repository.Db.SendBatch(context.Background(), batch)
	defer br.Close()

	for range clients {
		_, err := br.Exec()
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Errorf("error %v", err)
		}
	}
	return nil
}
