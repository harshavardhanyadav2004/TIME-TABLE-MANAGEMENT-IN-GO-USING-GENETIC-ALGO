package insertdatafolder
import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"

)
func ReadFromCoursesTable() map[string][]string {
	 coursesMap := make(map[string][]string)
	 db,err:= sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err != nil {
		log.Fatal(err)
	 }
	 defer db.Close()
	 rows, err := db.Query("SELECT * FROM courses")
	 if err != nil {
		log.Fatal(err)
	 }
	 for rows.Next(){
		var course_id string 
		var course_name string 
		var branch_id string 
		var semester int
		rows.Scan(&course_id,&course_name,&branch_id ,&semester)
		//inserting data into the map
		coursesMap[branch_id] = append(coursesMap[branch_id],course_id)

	 }
	 return coursesMap
}