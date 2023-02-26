package pkg
import (
    "database/sql"
    _ "github.com/lib/pq"
)
func Db() *sql.DB{
	connStr := "user=postgres password=601246 dbname=golang1 sslmode=disable"
    DB, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
	return DB
}