package parser

import (
	"bufio"
	"fmt"
	"github.com/Nhanderu/brdoc"
	"github.com/google/uuid"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/VictorOliveiraPy/internal/constants"
	"github.com/VictorOliveiraPy/internal/entity"
)

func TrimValue(value string) string {
	value = strings.Trim(value, " ")
	if value == "NULL" {
		return ""
	}
	return value
}

func ParseFile(filePath string) ([]*entity.ClientRaw, error) {
	file, err := os.Open("tmp/" + filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	allClients := make([]*entity.ClientRaw, 0)
	fileScanner.Scan()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		document := TrimValue(line[constants.DocumentIndexStart:constants.DocumentIndexEnd])
		private := line[constants.PrivateIndexStart:constants.PrivateIndexEnd]
		isIncomplete := line[constants.IsIncompleteIndexStart:constants.IsIncompleteIndexEnd]
		lastPurchaseDate := TrimValue(line[constants.LastPurchaseDateIndexStart:constants.LastPurchaseDateIndexEnd])
		averageTicket := TrimValue(line[constants.AverageTicketIndexStart:constants.AverageTicketIndexEnd])
		lastPurchaseTicket := TrimValue(line[constants.LastPurchaseTicketIndexStart:constants.LastPurchaseTicketIndexEnd])
		mostFrequentStore := TrimValue(line[constants.MostFrequentStoreIndexStart:constants.MostFrequentStoreIndexEnd])
		lastPurchaseStore := TrimValue(line[constants.LastPurchaseStoreIndexStart:])
		now := time.Now().Format(time.RFC3339)

		client := &entity.ClientRaw{
			ID:                 uuid.New().String(),
			Document:           document,
			IsPrivate:          private,
			IsIncomplete:       isIncomplete,
			LastPurchaseDate:   &lastPurchaseDate,
			AverageTicket:      &averageTicket,
			LastPurchaseTicket: &lastPurchaseTicket,
			MostFrequentStore:  &mostFrequentStore,
			LastPurchaseStore:  &lastPurchaseStore,
			Status:             "Waiting",
			CreatedAt:          now,
		}
		allClients = append(allClients, client)
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	return allClients, nil
}

func ParseDocumentValue(document string) (string, error) {
	if brdoc.IsCPF(document) {
		return "CPF", nil
	} else if brdoc.IsCNPJ(document) {
		return "CNPJ", nil
	}
	return "", fmt.Errorf("documento inválido")

}

func RemoveNonNumericCharacters(str string) string {
	re := regexp.MustCompile("[^0-9]+")
	cleanedStr := re.ReplaceAllString(str, "")

	return cleanedStr
}

func ParseFloat(value string) (float64, error) {
	str := strings.Replace(value, ",", ".", 1)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return f, nil
	}
	return f, nil
}

func ParseBoolValue(value string) (bool, error) {
	parsedValue, err := strconv.ParseBool(TrimValue(value))
	if err != nil {
		return false, err
	}

	return parsedValue, nil
}

func parseDate(dateString string) (*time.Time, error) {
	if TrimValue(dateString) == "" {
		return nil, nil
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return &time.Time{}, err
	}

	utcDate := date.UTC()
	return &utcDate, nil
}

func PointerToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func ParseClient(clients []*entity.ClientRaw) ([]*entity.Client, error) {
	allClients := make([]*entity.Client, 0)

	for _, c := range clients {
		documentType, err := ParseDocumentValue(c.Document)
		if err != nil {
			continue
		}

		averageTicket, err := ParseFloat(TrimValue(PointerToString(c.AverageTicket)))
		if err != nil {
			return nil, err
		}

		lastPurchaseTicket, err := ParseFloat(TrimValue(PointerToString(c.LastPurchaseTicket)))
		if err != nil {
			return nil, err
		}
		isPrivate, err := ParseBoolValue(c.IsPrivate)
		if err != nil {
			return nil, err
		}

		isIncomplete, err := ParseBoolValue(c.IsIncomplete)
		if err != nil {
			return nil, err
		}

		lastDate, err := parseDate(PointerToString(c.LastPurchaseDate))
		if err != nil {
			fmt.Println("Erro ao converter a string para data:", err)
			return nil, err
		}

		client := &entity.Client{
			ID:                 TrimValue(c.ID),
			Document:           RemoveNonNumericCharacters(TrimValue(c.Document)),
			DocumentType:       documentType,
			IsPrivate:          isPrivate,
			IsIncomplete:       isIncomplete,
			LastPurchaseDate:   lastDate,
			AverageTicket:      averageTicket,
			LastPurchaseTicket: lastPurchaseTicket,
			MostFrequentStore:  RemoveNonNumericCharacters(TrimValue(PointerToString(c.MostFrequentStore))),
			LastPurchaseStore:  RemoveNonNumericCharacters(TrimValue(PointerToString(c.LastPurchaseStore))),
			Status:             "concluded",
			CreatedAt:          time.Now().Format(time.RFC3339),
			UpdatedAt:          time.Now().Format(time.RFC3339),
		}

		allClients = append(allClients, client)

	}
	return allClients, nil

}
