package main

import (
	"database/sql"
	"info441-final-project/servers/todolist/handlers"
	"info441-final-project/servers/todolist/middlewares"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/stats"
	"info441-final-project/servers/todolist/models/tasks"
	"info441-final-project/servers/todolist/models/users"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
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

	SESSIONKEY := os.Getenv("SESSIONKEY")
	if len(SESSIONKEY) == 0 {
		log.Fatal("No SESSIONKEY environment variable found")
	}

	REDISADDR := os.Getenv("REDISADDR")
	if len(REDISADDR) == 0 {
		log.Fatal("No REDISADDR environment variable found")
	}

	DSN := os.Getenv("DSN")
	if len(DSN) == 0 {
		log.Fatal("No DSN environment variable found")
	}

	// Connect to redis
	log.Printf("connecting to Redis...")
	redis := redis.NewClient(&redis.Options{
		Addr:     REDISADDR,
		Password: "",
		DB:       0,
	})
	_, err := redis.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to Redis")

	// Connect to MySQL
	log.Printf("connecting to MySQL...")
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
	statsStore := &stats.MySQLStore{Client: mysql, UserStore: userStore}
	ctx := &handlers.HandlerContext{
		UserStore:    userStore,
		TaskStore:    taskStore,
		StatStore:    statsStore,
		SigningKey:   SESSIONKEY,
		SessionStore: sessions.NewRedisStore(redis, time.Hour),
	}

	// Create handlers
	r := middlewares.NewSessionMux(ctx)

	r.HandleFunc("/users", ctx.UsersHandler)
	r.HandleSessionFunc("/users/", ctx.SpecificUserHandler)

	r.HandleFunc("/sessions", ctx.SessionsHandler)
	r.HandleFunc("/sessions/", ctx.SpecificSessionHandler)

	r.HandleFunc("/stats", ctx.AllStatsHandler)
	r.HandleFunc("/stats/", ctx.PeriodicStatsHandler)

	r.HandleSessionFunc("/helloworld", ctx.TodoList)

	// Wrap with cors
	wrappedMux := &middlewares.Cors{Handler: r}

	// Serve
	log.Printf("server is listening at %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, wrappedMux))
}
