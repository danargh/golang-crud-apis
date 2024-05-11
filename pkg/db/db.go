package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// func LoadEnv() (ConfigDB, error) {
// 	var cnf ConfigDB

// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		return cnf, fmt.Errorf("error loading .env file: %v", err)
// 	}

// 	// Set configuration from environment variables
// 	cnf.Host = os.Getenv("DB_HOST")
// 	cnf.Port = 5432 // Default PostgreSQL port
// 	cnf.User = os.Getenv("DB_USER")
// 	cnf.Password = os.Getenv("DB_PASSWORD")
// 	cnf.Name = os.Getenv("DB_NAME")

// 	return cnf, nil
// }

func Connect(cnf ConfigDB) (*sqlx.DB, error) {
	// dsn = data source name
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
	)
	db, err := sqlx.Connect("postgres", dsn)
	return db, err
}
