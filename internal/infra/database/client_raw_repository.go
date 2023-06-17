package database

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
)

type ClientRawRepository struct {
	Db *pgxpool.Pool
}

func NewClientRawRepository(db *pgxpool.Pool) *ClientRawRepository {
	return &ClientRawRepository{Db: db}
}

func (repository *ClientRawRepository) Create(clients []*entity.ClientRaw) error {
	const insertQuery = `INSERT INTO raw_client_data (id, document, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	batch := &pgx.Batch{}
	for _, client := range clients {
		batch.Queue(
			insertQuery,
			client.ID,
			client.Document,
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
			return err
		}
	}

	return nil
}

func (repository *ClientRawRepository) GetClients(limit int, status string) ([]*entity.ClientRaw, error) {
	query := "SELECT id, document, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status FROM raw_client_data WHERE status = $1 LIMIT " + strconv.Itoa(limit)

	var clients []*entity.ClientRaw

	rows, err := repository.Db.Query(context.Background(), query, status)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var client entity.ClientRaw
		err := rows.Scan(
			&client.ID,
			&client.Document,
			&client.IsPrivate,
			&client.IsIncomplete,
			&client.LastPurchaseDate,
			&client.AverageTicket,
			&client.LastPurchaseTicket,
			&client.MostFrequentStore,
			&client.LastPurchaseStore,
			&client.Status,
		)

		if rows.Err() != nil {
			return nil, err
		}
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		clients = append(clients, &client)
	}

	return clients, nil

}

func (repository *ClientRawRepository) UpdateStatusClient(clients []*entity.Client) error {
	query := "UPDATE raw_client_data SET status = $2, updated_at = $3 WHERE id = $1"
	batch := &pgx.Batch{}
	for _, client := range clients {
		batch.Queue(
			query,
			client.ID,
			client.Status,
			client.UpdatedAt,
		)
	}

	br := repository.Db.SendBatch(context.Background(), batch)

	defer br.Close()

	for range clients {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}
