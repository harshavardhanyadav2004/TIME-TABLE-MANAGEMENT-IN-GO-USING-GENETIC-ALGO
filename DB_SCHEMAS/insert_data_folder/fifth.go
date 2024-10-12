package insertdatafolder

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
func MapTheCourses()map[string]string {
	MappedCourses := make(map[string]string )
	db,err := sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}	
	defer db.Close()
	rows,new_err := db.Query("SELECT * FROM COURSES")
	if new_err != nil {
		log.Fatal(new_err)
	}
	for rows.Next(){
		var course_id string 
		var course_name string 
		var branch_id string 
		var semester int 
		rows.Scan(&course_id,&course_name,&branch_id,&semester)
		MappedCourses[course_id] = course_name
	}
	return MappedCourses
}