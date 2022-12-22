package example

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func TODO() {
	err := CreateDB("exampleDB134")
	if err != nil {
		log.Println(err)
	}
}

const (
	USER string = "root"
	PASS string = "qwerty123"
	URL  string = "127.0.0.1"
	PORT string = "3306"
)

func DBcom(databasename string, cmd string) (*sql.Rows, error) {
	user := USER
	pass := PASS
	url := URL
	port := PORT

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, url, port, databasename)
	db, err := sql.Open("mysql", uri)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer db.Close()

	result, err := db.Query(cmd)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DBConn(databasename string, do func(db *sql.DB) (*sql.Rows, error)) (*sql.Rows, error) {

	user := USER
	pass := PASS
	url := URL
	port := PORT

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, url, port, databasename)

	db, err := sql.Open("mysql", uri)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer db.Close()

	result, err := do(db)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func command(databasename string, cmd string) (*sql.Rows, error) {
	result, err := DBConn(databasename, func(db *sql.DB) (*sql.Rows, error) {
		result, err := db.Query(cmd)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateDB(databasename string) error {
	cmd := fmt.Sprintf("CREATE DATABASE  %s;", databasename)
	_, err := DBcom("", cmd)
	if err != nil {
		return err
	}
	return nil
}

func CreateTable(databasename string, tablename string, structure string) error {
	cmd := fmt.Sprintf("CREATE TABLE [IF NOT EXISTS] %s(%s);", tablename, structure)
	_, err := DBcom(databasename, cmd)
	if err != nil {
		return err
	}
	return nil
}
