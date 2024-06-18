package main

import (
	"fmt"

	"github.com/google/uuid"
)

func generateUUID() {
	newUUID := uuid.New()
	fmt.Println("Generated UUID:", newUUID)
}

func parseUUID() {
	str := "123e4567-e89b-12d3-a456-426614174000"
	parsedUUID, err := uuid.Parse(str)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}
	fmt.Println("Parsed UUID:", parsedUUID)
}

func isValidUUID(str string) {
	_, err := uuid.Parse(str)
	if err != nil {
		fmt.Println("Invalid UUID")
	} else {
		fmt.Println("Valid UUID")
	}
}

/*
go run base.go
*/
func main() {
	generateUUID()
	parseUUID()

	isValidUUID("123e4567-e89b-12d3-a456-426614174000")
}
