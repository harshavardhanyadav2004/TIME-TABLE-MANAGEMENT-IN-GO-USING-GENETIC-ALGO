package insertdatafolder

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
)

func sanitizeTableName(tableName string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	return re.ReplaceAllString(tableName, "_")
}

func InsertIntoDirectlyIntoHost(total_time_table_map map[string]map[string][]string) {
	var days = []string{"Mon", "Tue", "Wed", "Thurs", "Fri"}
	dbUrl := "postgresql://postgres:dRBlRNYqWCNynfNhsKyQRLgPSwwcJDoA@autorack.proxy.rlwy.net:33642/railway"
	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	for branch, branch_days := range total_time_table_map {
		sanitizedBranch := sanitizeTableName(branch)
		createTableQuery := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS "%s" (
				day VARCHAR(50),
				"8-9"   VARCHAR(50),
				"9-10"  VARCHAR(50),
				"10-11" VARCHAR(50),
				"11-12" VARCHAR(50),
				"12-01" VARCHAR(50),
				"01-02" VARCHAR(50),
				"02-03" VARCHAR(50),
				"03-04" VARCHAR(50),
				PRIMARY KEY(day)
			)`, sanitizedBranch)

		_, err := conn.Exec(context.Background(), createTableQuery)
		if err != nil {
			log.Printf("Error creating table for branch %s: %v\n", branch, err)
			continue
		}

		for i := 0; i < len(days); i++ {
			array := branch_days[days[i]]
			insertQuery := fmt.Sprintf(`
				INSERT INTO "%s" (day, "8-9", "9-10", "10-11", "11-12", "12-01", "01-02", "02-03", "03-04")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, sanitizedBranch)

			_, err := conn.Exec(context.Background(), insertQuery,
				days[i], array[0], array[1], array[2], "Lunch_Period", array[3], array[4], array[5], array[6])
			if err != nil {
				log.Printf("Error inserting data for branch %s, day %s: %v\n", branch, days[i], err)
			}
		}
	}

	fmt.Println("Inserted Successfully")
}
