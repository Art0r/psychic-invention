package utils

import (
	"fmt"
	"log"
	"os"
)

func GetQueryAsString(file string) string {
	queryFile := fmt.Sprintf("sql/%s.sql", file)
    queryBytes, err := os.ReadFile(queryFile)
    if err != nil {
        log.Fatalf("Error reading query file: %v", err)
    }

	query := string(queryBytes)

	return query
}