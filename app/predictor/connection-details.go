package predictor

import "os"

func getDbCredentials(appName string) string {
	if appName == "PremierPredictor" {
		return os.Getenv("DATABASE_URL")
	} else {
		return os.Getenv("HEROKU_POSTGRESQL_GRAY_URL")
	}
}