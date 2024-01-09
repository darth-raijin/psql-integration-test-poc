package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
)

var (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	dbname     = "poc"
	password   = "1234"
	sslmode    = "disable"
	maxRetries = 25
	lock       sync.Once
)

// CreateTestDB creates a new test database from the custom template.
func CreateTestDB(testDBName string) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, "template1", port)
	templateDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil && !strings.Contains(err.Error(), "too many clients") {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// Connect to the template1 database
	// Keep in mind that this sync only blocks for each respective package so you might need some retry mechanic
	lock.Do(func() {
		err = migrateTemplateDatabase(templateDB)
		if err != nil {
			log.Fatalf("failed to create databases: %v", err)
		}
	})

	// Create a new test database from template1 and close connection as we don't need it anymore
	err = templateDB.Exec(fmt.Sprintf("CREATE DATABASE %s TEMPLATE template1", testDBName)).Error
	if err != nil {
		log.Fatalf("failed to create test database '%s' from template1: %v", testDBName, err)
	}

	sql, err := templateDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sql.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the new test database and return the connection
	testDBDsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=%s", host, user, testDBName, port, password, sslmode)
	testDB, err := gorm.Open(postgres.Open(testDBDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	return testDB
}

func migrateTemplateDatabase(templateDB *gorm.DB) error {
	// Run migration on template1
	if err := Migrate(templateDB); err != nil {
		log.Fatalf("failed to migrate template1: %v", err)
	}

	return nil
}

// CreateDefaultTestDB creates a new test database without using a template.
func CreateDefaultTestDB(testDBName string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=%s", host, user, dbname, port, password, sslmode)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// Create a new test database
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName)).Error
	if err != nil {
		log.Fatalf("failed to create test database '%s': %v", testDBName, err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sql.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the new test database
	testDBDsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=%s", host, user, testDBName, port, password, sslmode)
	testDB, err := gorm.Open(postgres.Open(testDBDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	return testDB
}
