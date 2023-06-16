package parser

import (
	"github.com/VictorOliveiraPy/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFileCountClients(t *testing.T) {

	clients, err := parser.ParseFile("test.txt")
	if err != nil {
		t.Fatalf("Erro ao analisar o arquivo: %s", err)
	}
	expectedCount := 2
	assert.Equal(t, len(clients), expectedCount)
}

func TestParseFileClientFields(t *testing.T) {
	filePath := "test.txt"

	clients, err := parser.ParseFile(filePath)
	if err != nil {
		t.Fatalf("Erro ao analisar o arquivo %s: %s", filePath, err)
	}

	for _, client := range clients {
		assert.NotEmpty(t, client.ID)
		assert.NotEmpty(t, client.Document)
		assert.NotEmpty(t, client.IsPrivate)
	}
}

func TestParseFileErrorHandling(t *testing.T) {
	filePath := "nonexistent.txt"
	_, err := parser.ParseFile(filePath)
	assert.Error(t, err, "Erro esperado ao analisar um arquivo inexistente")
}

func TestRemoveNonNumericCharactersWithNonNumericCharacters(t *testing.T) {
	str := "1a2b3c4d5e6f"
	expected := "123456"
	result := parser.RemoveNonNumericCharacters(str)
	assert.Equal(t, expected, result)
}

func TestRemoveNonNumericCharactersWithOnlyNumericCharacters(t *testing.T) {
	str := "1234567890"
	expected := "1234567890"
	result := parser.RemoveNonNumericCharacters(str)
	assert.Equal(t, expected, result)
}

func TestRemoveNonNumericCharactersWithNoNumericCharacters(t *testing.T) {
	str := "abcd"
	expected := ""
	result := parser.RemoveNonNumericCharacters(str)
	assert.Equal(t, expected, result)
}

func TestParseFloatWithValidValue(t *testing.T) {
	value := "10.5"
	expected := 10.5
	result, err := parser.ParseFloat(value)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result)
}

func TestParseFloatWithCommaSeparator(t *testing.T) {
	value := "100,25"
	expected := 100.25
	result, err := parser.ParseFloat(value)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result)
}

func TestParseBoolValueWithValidTrueValue(t *testing.T) {
	value := "true"
	expected := true
	result, err := parser.ParseBoolValue(value)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result)
}

func TestParseBoolValueWithValidFalseValue(t *testing.T) {
	value := "false"
	expected := false
	result, err := parser.ParseBoolValue(value)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result)
}

func TestParseBoolValueWithInvalidValue(t *testing.T) {
	value := "invalid"
	_, err := parser.ParseBoolValue(value)
	assert.NotNil(t, err)
}

func TestTrimValueWithTrimmingSpaces(t *testing.T) {
	value := "  value  "
	expected := "value"
	result := parser.TrimValue(value)
	assert.Equal(t, expected, result)
}

func TestTrimValueWithNullValue(t *testing.T) {
	value := "NULL"
	expected := ""
	result := parser.TrimValue(value)
	assert.Equal(t, expected, result)
}

func TestTrimValueWithEmptyValue(t *testing.T) {
	value := ""
	expected := ""
	result := parser.TrimValue(value)
	assert.Equal(t, expected, result)
}
