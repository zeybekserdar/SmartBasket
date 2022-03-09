package db

import (
	"SmartBasket/helpers"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user = helpers.GetEnvKey("DB_USER")
var password = helpers.GetEnvKey("DB_PASS")
var host = helpers.GetEnvKey("DB_HOST")
var port = helpers.GetEnvKey("DB_PORT")
var db = helpers.GetEnvKey("DB_NAME")

func Connect() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
