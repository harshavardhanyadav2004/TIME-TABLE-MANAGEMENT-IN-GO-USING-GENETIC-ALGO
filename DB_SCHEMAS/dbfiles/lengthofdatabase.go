package dbfiles

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

func lengthOfDatabase(tableName string) int {
    db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
    var count int
    err = db.QueryRow(query).Scan(&count)
    if err != nil {
        log.Fatal(err)
    }

    return count + 1
}

