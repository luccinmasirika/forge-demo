package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const version = "v4"

func main() {
	started := time.Now().UTC().Format(time.RFC3339)
	host, _ := os.Hostname()

	var db *pgxpool.Pool
	if url := os.Getenv("DATABASE_URL"); url != "" {
		var err error
		db, err = pgxpool.New(context.Background(), url)
		if err != nil {
			log.Fatalf("database: %v", err)
		}
		if _, err := db.Exec(context.Background(),
			`create table if not exists visits (id bigserial primary key, at timestamptz not null default now())`); err != nil {
			log.Fatalf("database: %v", err)
		}
		log.Print("connected to the database")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visits := "no database"
		if db != nil {
			var n int64
			err := db.QueryRow(r.Context(), `insert into visits default values returning id`).Scan(&n)
			if err != nil {
				http.Error(w, "database error", http.StatusInternalServerError)
				log.Printf("visit: %v", err)
				return
			}
			visits = fmt.Sprint(n)
		}
		fmt.Fprintf(w, "Hello from Forge!\nversion: %s\nstarted: %s\ncontainer: %s\nvisits: %s\n", version, started, host, visits)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	log.Printf("forge-demo %s listening on :8080", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
