package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func IsUUID(str string) bool {
    uuidRegex := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
    return uuidRegex.MatchString(str)
}