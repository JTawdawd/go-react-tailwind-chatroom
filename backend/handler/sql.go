package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

var DB *sql.DB

func init() {
	DB = getDBConnection()
}

func getSqlString(queryName string) (string, error) {
	filePath := fmt.Sprintf("./query/%v.sql", queryName)

	sqlContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.New("Error reading " + queryName)
	}

	return string(sqlContent[:]), nil
}

func query(queryName string, args ...any) (*sql.Rows, error) {
	sqlString, err := getSqlString(queryName)
	if err != nil {
		return nil, err
	}

	return DB.Query(sqlString, args...)
}

func insert(queryName string, args ...any) error {
	sqlString, err := getSqlString(queryName)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare(sqlString)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func getDBConnection() *sql.DB {
	f, err := os.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile("DB_PASSWORD=\"(.*?)\"")
	matches := pattern.FindStringSubmatch(string(f[:]))
	if matches == nil {
		log.Fatal("Failed to retrieve database password")
	}
	os.Setenv("DB_PASSWORD", matches[1])
	pgsqlDetails := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, os.Getenv("DB_PASSWORD"), dbname)
	log.Printf(os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", pgsqlDetails)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Connected to database :)")
	return db
}
