package main

import (
	"fmt"
	"github.com/darth-raijin/psql-integration-test-poc/models"
	"github.com/google/uuid"
	"log"
	"strings"
	"testing"
	"time"
)

// BenchmarkCreateTestDB benchmarks the database creation using CreateTestDB (from template1)
func BenchmarkCreateTestDB(b *testing.B) {
	for i := 0; i < 100; i++ {
		currentTime := time.Now()
		testDB := CreateTestDB(generateDatabaseName())

		testDB.Save(&models.User{
			Username:    "some username",
			Email:       "some email",
			Password:    "some password",
			FirstName:   "John",
			LastName:    "Doe",
			DateOfBirth: time.Now(),
		})

		// Finally close connection
		sql, err := testDB.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = sql.Close()
		if err != nil {
			return
		}

		duration := time.Since(currentTime) // Calculate the duration
		fmt.Printf("Completed iteration in %v\n", duration)
	}
}

// BenchmarkCreateDefaultTestDB benchmarks the database creation using CreateDefaultTestDB (without a template)
func BenchmarkCreateDefaultTestDB(b *testing.B) {
	for i := 0; i < 100; i++ {
		currentTime := time.Now()
		// Create a new test database without using a template
		testDB := CreateDefaultTestDB(generateDatabaseName())

		err := Migrate(testDB)
		if err != nil {
			return
		}

		testDB.Save(&models.User{
			Username:    "some username",
			Email:       "some email",
			Password:    "some password",
			FirstName:   "John",
			LastName:    "Doe",
			DateOfBirth: time.Now(),
		})

		// Finally close connection
		sql, err := testDB.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = sql.Close()
		if err != nil {
			log.Fatal(err)
		}
		duration := time.Since(currentTime) // Calculate the duration
		fmt.Printf("Completed iteration in %v\n", duration)
	}
}

func generateDatabaseName() string {
	return fmt.Sprintf("test_%v", strings.Replace(uuid.NewString(), "-", "", -1))
}
