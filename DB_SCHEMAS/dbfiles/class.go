package dbfiles
import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)
func CreateClass(){
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`
	CREATE TABLE IF NOT EXISTS CLASS 
	(
	class_id VARCHAR(100) ,
	 faculty_course_id VARCHAR(10), 
	 year INTEGER ,
	 semester INTEGER , 
	 time_Slot VARCHAR(50),
	 PRIMARY KEY(class_id),
	CONSTRAINT fk_class
	  FOREIGN KEY (faculty_course_id) REFERENCES FACULTY_COURSE(faculty_course_id)
	)`)
	defer db.Close()
}
func InsertIntoClass(faculty_course_id string  , year int , semester int , time_slot string){
	var class_id string = ""
	var class_count int = lengthOfDatabase("CLASS")
	if class_count < 10 {
		class_id = "CCREC" + strconv.Itoa(class_count)
	}else {
		class_id = "CCREC" + strconv.Itoa(class_count)

	}
	class_count++
	   db,err:=sql.Open("postgres","user=postgres dbname= mydb sslmode = disable")
	   if err !=nil {
	       log.Fatal(err)
	   }
	   _,another_err:= db.Exec(`
	    INSERT INTO CLASS (class_id , faculty_course_id , year , semester, time_slot ) VALUES (?,?,?,?,?) 
	   ` , class_id , faculty_course_id,year , semester ,time_slot)
	   if another_err!=nil {
		log.Fatal(err)
	   }
	   defer db.Close()
}