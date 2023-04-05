package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type ArticlesDb struct {
	postgres *sql.DB
	l        *zap.Logger
}

// a struct to hold all the db connection information
type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDB(l *zap.Logger) *ArticlesDb {
	return &ArticlesDb{nil, l}
}

// InitDB initializes the database connection
func (db *ArticlesDb) Connect() {
	err := godotenv.Load("config/.env")

	db.l.Info("Loaded env ")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	db.l.Info("connInfo ")
	connInfo := connection{
		Host:     os.Getenv("DB_URL"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db.l.Info("Caaling sql.open ")
	fmt.Printf("%s", connInfo)
	// Replace the connection string with your PostgreSQL connection details
	db.postgres, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		log.Fatal(err)
	}
	defer db.postgres.Close()
	db.l.Info("Pinging sql.open ")
	// Ping the database to ensure a connection is established
	err = db.postgres.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

// Take our connection struct and convert to a string for our db connection info
func connToString(info connection) string {
	fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}
