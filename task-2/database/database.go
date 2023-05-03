package Db

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"fmt"
	"log"
	"os"

	"github.com/ashokdivi/task-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db  user=%s password=%s dbname=%s port=5432 sslmode=disable ",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database.\n", err)
		os.Exit(2)
	}
	fmt.Println("connected")
	fmt.Println("Run migrations")
	db.AutoMigrate(&models.Employee{})

	DB = DbInstance{
		Db: db,
	}
}
