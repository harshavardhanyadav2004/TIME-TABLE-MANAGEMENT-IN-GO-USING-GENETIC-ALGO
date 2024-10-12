package insertdatafolder
import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"

)
func ReadFromTheFacultyTable() map[string][]string {
	facultyMap := make(map[string][]string)
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err !=nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows ,new_err := db.Query("SELECT * FROM  FACULTY")
	if new_err != nil {
		log.Fatal(new_err)
	}
	for rows.Next(){
		var faculty_id string 
		var faculty_name string 
		var email string 
		var phone_no string 
		rows.Scan(&faculty_id,&faculty_name,&email,&phone_no)
		facultyMap[faculty_id]=append(facultyMap[faculty_id], faculty_name)
	}
	return facultyMap

}