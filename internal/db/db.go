package db

import (
	"fmt"
	"os"
	"github.com/GlebZigert/go-grpc-services-course/internal/rocket"
	

	"github.com/jmoiron/sqlx"
)

// DB
type Store struct {
	db *sqlx.DB
}

// New - returns a new Store
func New() (Store, error) {

	fmt.Println("1")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println("2")
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectString)
	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		fmt.Println("Failed to sqlx.Connect")
		return Store{}, err
	}
	fmt.Println("4")	
	return Store{
		db: db,
	}, nil
}

// GetRocketByID - returns a rocket from the database by a given ID
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// DeleteRocket - deletes a rocket from the database by it's ID
func (s Store) DeleteRocket(id string) error {
	return nil
}
