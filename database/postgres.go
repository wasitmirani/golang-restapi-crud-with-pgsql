package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Book struct {
	gorm.Model
	Title           string  `json:"title"`
	Author          string  `json:"author"`
	PublicationYear int     `json:"publication_year"`
	Genre           string  `json:"genre"`
	ISBN            string  `json:"isbn" gorm:"unique"`
	Language        string  `json:"language"`
	PageCount       int     `json:"page_count"`
	Price           float64 `json:"price"`
	IsAvailable     bool    `json:"is_available" gorm:"default:true"`
}

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "golang_crud"
	dbUser := "root"
	password := ""
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Book{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
