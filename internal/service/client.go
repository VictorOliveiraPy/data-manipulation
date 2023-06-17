package service

import (
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/parser"
)

type ClientService struct {
	clientRawRepository entity.ClientRawRepositoryInterface
	clientRepository    entity.ClientRepositoryInterface
}

func NewClientService(clientRawRepository entity.ClientRawRepositoryInterface, clientRepository entity.ClientRepositoryInterface) *ClientService {
	return &ClientService{
		clientRawRepository: clientRawRepository,
		clientRepository:    clientRepository,
	}
}

func (service *ClientService) LoadRawDataFromFile(filePath string) error {
	allClients, err := parser.ParseFile(filePath)
	if err != nil {
		return err
	}

	err = service.clientRawRepository.Create(allClients)
	if err != nil {
		return err
	}

	return nil
}

func (service *ClientService) CleanAndLoadData(limit int, status string) error {

	for {
		var allClientsClean []*entity.Client
		result, err := service.clientRawRepository.GetClients(limit, status)

		if err != nil {
			return err
		}
		if len(result) == 0 {
			fmt.Println("Processing Finish")
			break
		}
		cleaned, err := parser.ParseClient(result)
		if err != nil {
			return err
		}

		estimatedSize := len(cleaned) * 2
		allClientsClean = make([]*entity.Client, 0, estimatedSize)

		allClientsClean = append(allClientsClean, cleaned...)

		err = service.clientRepository.Create(allClientsClean)
		if err != nil {
			return err
		}
		err = service.clientRawRepository.UpdateStatusClient(allClientsClean)
		if err != nil {
			return err
		}
	}
	return nil
}
