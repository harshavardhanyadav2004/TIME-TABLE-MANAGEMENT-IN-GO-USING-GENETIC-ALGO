package dbfiles

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)
func CreateCourses(){
	 db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err!=nil {
		 log.Fatal(err)
	 }
	 db.Exec(`
	 CREATE TABLE IF NOT EXISTS COURSES (
	  course_id VARCHAR(10),
	  course_name VARCHAR(100),
	  branch_id VARCHAR(10),
	  semester INTEGER,
	  PRIMARY KEY (course_id),
	  CONSTRAINT fk_courses 
	  FOREIGN KEY (branch_id) REFERENCES BRANCHES(branch_id) 
	)
	 `)
	 defer db.Close()
}
func InsertIntoCourses(course_name string , branch_id string , semester int ){
	var course_id string = ""
	var course_count int = lengthOfDatabase("COURSES")

	if course_count < 10 {
		 course_id = "RECC0"+strconv.Itoa(course_count)
	}else {
		 course_id = "RECC"+strconv.Itoa(course_count)
	}
	 db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err != nil {
		log.Fatal(err)
	 }
	 defer db.Close()
	 _,new_err := db.Exec(`INSERT INTO COURSES (course_id , course_name , branch_id , semester)
	                VALUES ($1,$2,$3,$4) `,course_id,course_name, branch_id , semester)
	if new_err != nil {
		log.Fatal(err)
	}

}