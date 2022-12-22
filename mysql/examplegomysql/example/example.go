package example

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Example() {

	user:= "root"
	pass:= "qwerty123"
	url:= "127.0.0.1"
	port:= "3306"
	databasename:= "db"
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,pass,url,port,databasename)

    // Open up our database connection.
    db, err := sql.Open("mysql",uri)

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

   CreateDB("exampleDB",db)
}


type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}




func command (cmd string,db *sql.DB)(*sql.Rows,error){
results, err := db.Query(cmd)
    if err != nil {
		return nil, err
    }
	return results , nil
}
func CreateDB(databasename string,db *sql.DB){
	cmd:=  fmt.Sprintf("CREATE DATABASE %s;",databasename)
	_, err:= command(cmd, db)
	log.Println(err)
}


func CreateTable(tablename string,structure string, db *sql.DB){

	cmd:= fmt.Sprintf("CREATE TABLE %s(%s);",tablename,structure)
	result, _:= command(cmd, db)
	log.Println(result)
}




func example(db *sql.DB){
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