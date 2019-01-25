package database

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const QUERY_GET_COUNT = "SELECT COUNT(*) as count FROM users"

func ConnectDb(appName string) (*sql.DB, error, string) {
	db, err := connect(appName)

	return db, err, QUERY_GET_COUNT
}

func connect(appName string) (*sql.DB, error) {
	url := getDbCredentials(appName)

	return connectPostgres(url)
}


func getDbCredentials(appName string) string {
	if appName == "PremierPredictor" {
		return os.Getenv("DATABASE_URL")
	} else {
		return os.Getenv("HEROKU_POSTGRESQL_GRAY_URL")
	}
}

func connectPostgres(url string) (*sql.DB, error) {
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	return sql.Open("postgres", connection)
}

func TestConnection(db *sql.DB, err error) error {
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return err
}