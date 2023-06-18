package main

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	clientRepository := database.NewClientRepository(conn)
	clientRawRepository := database.NewClientRawRepository(conn)

	clientService := service.NewClientService(clientRawRepository, clientRepository)

	clientService.LoadRawDataFromFile("base.txt")
	clientService.CleanAndLoadData(1000, "Waiting")
	elapsed := time.Since(startTime)
	message := fmt.Sprintf("[Done] exited with code=0 in %s", elapsed)
	fmt.Println(message)

}
