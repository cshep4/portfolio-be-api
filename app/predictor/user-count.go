package predictor

import (
	"net/http"
	"portfolio-be-api/app/service"
	"portfolio-be-api/app/database"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

type PredictorService struct{}

type PredictorArgs struct {
	AppName string
}

const QUERY_GET_COUNT = "SELECT COUNT(*) as count FROM users"

func (t *PredictorService) GetUserCount(r *http.Request, args *PredictorArgs, result *service.Response) error {
	url := getDbCredentials(args.AppName)

	db, err := database.ConnectDb(url)

	err = database.TestConnection(db, err)
	defer db.Close()

	if err != nil {
		return err
	}

	rows, err := db.Query(QUERY_GET_COUNT)
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

func (t *PredictorService) Path() string {
	return "/predictor"
}

func (t *PredictorService) Name() string {
	return "predictor"
}