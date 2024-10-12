package dbfiles

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"strconv"
)
func CreateStudents(){
	 db, err:= sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err!=nil {
		log.Fatal(err)
	 }
	 db.Exec(`
	 CREATE TABLE IF NOT EXISTS STUDENTS (
	      student_id VARCHAR(10),
		  student_name VARCHAR(50),
		  email VARCHAR(50),
		  phone_no VARCHAR(20),
		  branch_id VARCHAR(10),
		  year INTEGER , 
		  semester INTEGER ,
		  PRIMARY KEY(student_id),
		  CONSTRAINT fk_students
		  FOREIGN KEY(branch_id) REFERENCES BRANCHES(branch_id)
	)
	 `)
	defer db.Close()
}
func InsertIntoStudents(student_name string , email string , phone_no string ,branch_id string , year int ,semester int){
	var student_id string = ""
	var student_count int = lengthOfDatabase("STUDENTS")
	if student_count < 10 {
		 student_id = "21981A420"+strconv.Itoa(student_count)
	}else {
		 student_id = "21981A42"+strconv.Itoa(student_count)
	}
	 db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err!=nil {
		 log.Fatal(err)
	 }
	 _,first_err := db.Exec(`
	 INSERT INTO STUDENTS 
	 (student_id , student_name , email , phone_no , branch_id , year ,semester) VALUES (?,?,?,?,?,?,?)
	 `, student_id,student_name,email,phone_no,branch_id,year,semester)
	 if first_err!=nil {
		log.Fatal(first_err)
	 }
	 defer db.Close()
}