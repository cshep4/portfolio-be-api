package database

import (
	//"fmt"
	"database/sql"
	//"os"
	"github.com/lib/pq"
)

//func ConnectDb(host string, port string, user string, password string, dbname string) (*sql.DB, error) {
//	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
//
//	return sql.Open("postgres", sqlInfo)
//}

func ConnectDb(url string) (*sql.DB, error) {
	//url := os.Getenv("DATABASE_URL")
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