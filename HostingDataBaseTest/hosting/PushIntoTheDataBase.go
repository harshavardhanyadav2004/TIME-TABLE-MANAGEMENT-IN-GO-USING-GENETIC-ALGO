package hosting

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	
)
func ReadFromTheCsv() [][]string{
	file , err:= os.Open("C:\\Users\\Admin\\Desktop\\ClassRoom_Management\\HostingDataBaseTest\\hosting\\subjects.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	records ,new_err := reader.ReadAll()
	if new_err != nil {
		log.Fatal(err)
	}
	return records
}
func PushIntoTheHost(){
	 dbUrl:= "postgresql://postgres:dRBlRNYqWCNynfNhsKyQRLgPSwwcJDoA@autorack.proxy.rlwy.net:33642/railway"
	 all_records :=  ReadFromTheCsv()
	 conn, err := pgxpool.New(context.Background(), dbUrl)
	 if err != nil {
		 log.Fatalf("Unable to connect to database: %v\n", err)
	 }
	 defer conn.Close()
	 for _,record := range all_records[1:] {
		_,err:=conn.Exec(context.Background(),"INSERT INTO names (id,name,subject) VALUES ($1,$2,$3)",record[0],"Harsha",record[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted the rows into the database successfully ")

}
