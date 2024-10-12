package dbfiles

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func CreateBranches() {
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err!=nil {
		log.Fatal(err)
	}
	db.Exec(
	`
	CREATE TABLE IF NOT EXISTS BRANCHES(
		branch_id VARCHAR(10) ,
		branch_name VARCHAR(50) , 
		PRIMARY KEY(branch_id)
		);`)
	defer db.Close()
}
func InsertIntoTable(branch_name string) {
	var branch_id string = ""
	var branch_count int = lengthOfDatabase("BRANCHES")
	if branch_count < 10 {
		branch_id = "BREC0"+strconv.Itoa(branch_count)
	}else {
		 branch_id = "BREC"+strconv.Itoa(branch_count)
	}
	branch_count++
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	_,new_err:=db.Exec("INSERT INTO BRANCHES (branch_id , branch_name) VALUES ($1,$2)",branch_id,branch_name)
	if new_err != nil {
		log.Fatal(err)
	}
}
