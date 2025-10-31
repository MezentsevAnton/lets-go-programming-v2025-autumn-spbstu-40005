package parsers

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML[T any](filePath string) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error: opening a file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("Error: closing xml file")
		}
	}()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = charset.NewReaderLabel

	var result T
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("Error: xml decoding: %w", err)
	}

	return &result, nil
}
