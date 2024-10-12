package hosting 
import (
		"context"
		"fmt"
		"log"
	
		"github.com/jackc/pgx/v5/pgxpool"
)
func RetreiveFromTheDataBase(){
	 dbUrl:= "postgresql://postgres:dRBlRNYqWCNynfNhsKyQRLgPSwwcJDoA@autorack.proxy.rlwy.net:33642/railway"
	 conn,err:= pgxpool.New(context.Background(),dbUrl)
	 if err != nil {
		log.Fatal(err)
	 }
	 defer conn.Close()
	 rows , new_err:= conn.Query(context.Background(),"SELECT * FROM names")
	 if new_err != nil {
		 log.Fatal(new_err)
	 }
	 fmt.Println("Printing the records from the table")
	 for rows.Next(){
		var id int 
		var name string 
		var subject string 
		rows.Scan(&id,&name,&subject)
		fmt.Println(id,name ,subject)
	 }
	 fmt.Println("Succesfully Printed")

}