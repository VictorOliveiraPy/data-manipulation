package main

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/cmd/configs"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	appConfigs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfigs.DBHost,
		appConfigs.DBPort,
		appConfigs.DBUser,
		appConfigs.DBPassword,
		appConfigs.DBName,
	)

	conn, err := pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		panic(err)
	}

	clientRepository := database.NewClientRepository(conn)
	clientRawRepository := database.NewClientRawRepository(conn)

	clientService := service.NewClientService(clientRawRepository, clientRepository)

	clientService.LoadRawDataFromFile("base_49994.txt")
	clientService.CleanAndLoadData(1000, "Waiting")
	elapsed := time.Since(startTime)
	message := fmt.Sprintf("[Done] exited with code=0 in %s", elapsed)
	fmt.Println(message)

}
