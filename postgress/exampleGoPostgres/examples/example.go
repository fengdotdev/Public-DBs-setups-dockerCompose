package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Example(){
	ExampleHelloWorldConn()
	ExampleHelloWorldPool()
}

func ExampleHelloWorldConn(){

	fmt.Println("PostgresConnectTODO")
	PostgresConnectTODO(func(conn *pgx.Conn) {
		var greeting string
		err := conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(greeting)
	})
}

func ExampleHelloWorldPool(){
	fmt.Println("PostgresPoolTODO")
	PostgresPoolTODO(func(dbpool *pgxpool.Pool) {
	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)
	})
}

func PostgresConnectTODO(todo func(c *pgx.Conn)){
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	url:= "postgres://postgres:pass123@localhost:5432/cars"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	todo(conn)
}


/*
The *pgx.Conn returned by pgx.Connect() represents a single connection and is not concurrency safe. 
This is entirely appropriate for a simple command line example such as above. However, for many uses, 
such as a web application server, concurrency is required. 
To use a connection pool replace the import github.com/jackc/pgx/v5 with github.com/jackc/pgx/v5/pgxpool 
and connect with pgxpool.Connect() instead of pgx.Connect().
*/

func PostgresPoolTODO(todo func(p *pgxpool.Pool)){
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url:= "postgres://postgres:pass123@localhost:5432/cars"
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	todo(dbpool)
}





