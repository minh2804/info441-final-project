package main

import (
	"database/sql"
	"info441-final-project/servers/app/handlers"
	"info441-final-project/servers/app/models/tasks"
	"info441-final-project/servers/app/models/users"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Entry point for the server
func main() {
	// Get environment variables
	// ADDR := os.Getenv("ADDR")
	// if len(ADDR) == 0 {
	// 	ADDR = ":443"
	// }
	ADDR := os.Getenv("ADDR")
	if len(ADDR) == 0 {
		ADDR = ":80"
	}

	// TLSCERT := os.Getenv("TLSCERT")
	// if len(TLSCERT) == 0 {
	// 	log.Fatal("No TLSCERT environment variable found")
	// }

	// TLSKEY := os.Getenv("TLSKEY")
	// if len(TLSKEY) == 0 {
	// 	log.Fatal("No TLSKEY environment variable found")
	// }

	DSN := os.Getenv("DSN")
	if len(DSN) == 0 {
		log.Fatal("No DSN environment variable found")
	}

	// Connect to MySQL
	mysql, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()
	if err := mysql.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to MySQL")

	// Create context
	userStore := &users.MySQLStore{Client: mysql}
	taskStore := &tasks.MySQLStore{Client: mysql, UserStore: userStore}
	ctx := &handlers.HandlerContext{
		UserStore: userStore,
		TaskStore: taskStore,
	}

	// Create handlers
	r := mux.NewRouter()
	r.HandleFunc("/helloworld", ctx.TodoList)

	// Wrap with cors
	wrappedMux := &handlers.Cors{Handler: r}

	// Serve
	log.Printf("server is listening at %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, wrappedMux))
}
