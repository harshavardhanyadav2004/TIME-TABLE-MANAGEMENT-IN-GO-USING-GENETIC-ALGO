package insertdatafolder
import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"

)
func ReadFromTheBranchesTable() map[string][]string {
	 branchMap := make(map[string][]string)
	 db,err:= sql.Open("postgres","user=postgres dbname=mydb sslmode=disable")
	 if err != nil {
		log.Fatal(err)
	 }
	 defer db.Close()
	 rows,new_err := db.Query("SELECT * FROM BRANCHES")
	 if new_err != nil {
		log.Fatal(new_err)
	 }
	 for rows.Next(){
		var branch_id string 
		var branch_name string 
		rows.Scan(&branch_id ,&branch_name)
		branchMap[branch_id] = append(branchMap[branch_id], branch_name)
	 }
	 return branchMap
}