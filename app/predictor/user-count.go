package predictor

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"portfolio-be-api/app/database"
	"portfolio-be-api/app/service"
	"net/http"
	"strconv"
)

func (t *PredictorService) GetUserCount(r *http.Request, args *PredictorArgs, result *service.Response) error {
	db, err, query := database.ConnectDb(args.AppName)

	err = database.TestConnection(db, err)
	defer db.Close()

	if err != nil {
		return err
	}

	rows, err := db.Query(query)
	count, err := checkCount(rows)

	if err != nil {
		return err
	}

	*result = service.Response{StatusCode: 200, Body: strconv.Itoa(count)}

	return nil
}

func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}