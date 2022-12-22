package example

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}



func Example() {

	user:= "root"
	pass:= "qwerty123"
	url:= "127.0.0.1"
	port:= "3306"
    databasename:= "exampleCarDB" 
	uriNODB := fmt.Sprintf("%s:%s@tcp(%s:%s)/",user,pass,url,port)
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,pass,url,port,databasename)
if false{
 db, err := sql.Open("mysql",uriNODB)
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

     _, err = db.Query("CREATE DATABASE exampleCarDB ;")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}
   

    log.Println("done...")
     db1, err := sql.Open("mysql",uri)
    if err != nil {
        log.Print(err.Error())
    }
    defer db1.Close()

    exampleQueryDatabase(db1)

}



type cars struct{

}


func exampleQueryDatabase(db *sql.DB){
 if false{
    _, err := db.Query("CREATE TABLE  Cars (CarId int, Brand varchar(255), Model varchar(255));")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
 }
    
    _, err := db.Query("INSERT INTO  Cars VALUES  ('1','lada', '$606k');")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    result, err := db.Query("SELECT * FROM Cars")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    result.Close()

}
/*


exampleQueryCreateTable(){
    
}


func exampleCreateTable(db *sql.DB){
 // Execute the query
    results, err := db.Query("SELECT id, name FROM tags")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }
}

*/