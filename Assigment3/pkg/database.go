package pkg
import (
    // "database/sql"
	// "gorm.io/driver/sqlite"
    // _ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=601246 dbname=golang sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	} else {
		fmt.Println("ok")
	}
	return db
}