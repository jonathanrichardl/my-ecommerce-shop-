package testutil

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func NewTestDb() *sql.DB {
	godotenv.Load("../.env")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("TEST_DB_NAME"))

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
