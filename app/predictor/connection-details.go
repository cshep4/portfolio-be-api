package predictor

import "os"

func getDbCredentials(appName string) string {
	if appName == "PremiersPredictor" {
		return os.Getenv("DATABASE_URL")
	} else {
		return os.Getenv("HEROKU_POSTGRESQL_GRAY_URL")
	}
}