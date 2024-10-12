package dbfiles

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func CreateFaculty(){
	 db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err !=nil {
		log.Fatal(err)
	 }
	 db.Exec(`
	 CREATE TABLE IF NOT EXISTS FACULTY (
	      faculty_id VARCHAR(10),
		  faculty_name VARCHAR(50),
		  email VARCHAR(50),
		  phone_no VARCHAR(20),
		  PRIMARY KEY(faculty_id)
	);
	 `)
	defer db.Close()
}
func InsertIntoFaculty(faculty_name string , email string , phone_no string ){
	var faculty_id string = ""
	var faculty_count int = lengthOfDatabase("FACULTY")
	if faculty_count < 10 {
		faculty_id = "FACREC0"+strconv.Itoa(faculty_count)
	}else {
		faculty_id = "FACREC"+strconv.Itoa(faculty_count)
	}
	faculty_count++
	db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close()
	_,first_err:= db.Exec(`
	INSERT INTO FACULTY (faculty_id ,faculty_name , email , phone_no )
	VALUES ($1,$2,$3,$4)`,faculty_id , faculty_name,email , phone_no)
	if first_err != nil {
		log.Fatal(first_err)
	}

}