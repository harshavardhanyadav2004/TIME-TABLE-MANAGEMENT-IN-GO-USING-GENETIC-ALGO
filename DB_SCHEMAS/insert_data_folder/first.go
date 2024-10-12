package insertdatafolder

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
func ReadFromTheFacultyCourses() map[string][]string  {
  facultyCourses :=make(map[string][]string) 
	db,err:=sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	if err !=nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows,new_err:=db.Query("SELECT * FROM FACULTY_COURSE")
	if new_err != nil {
		log.Fatal(new_err)
	}
	for rows.Next(){
		var faculty_course_id string 
		var faculty_id string 
		var course_id string 
		rows.Scan(&faculty_course_id,&faculty_id,&course_id)
		facultyCourses[faculty_id]=append(facultyCourses[faculty_id],course_id)
	}
	return facultyCourses
}