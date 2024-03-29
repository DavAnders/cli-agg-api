package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/DavAnders/cli-agg-api/internal/config"
	"github.com/DavAnders/cli-agg-api/internal/database"
	"github.com/DavAnders/cli-agg-api/internal/handler"
	"github.com/DavAnders/cli-agg-api/internal/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	apiCfg := handler.ApiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/articles", apiCfg.HandlerCreateArticle)
	mux.HandleFunc("/articles/query", apiCfg.HandlerListArticleByQuery)
	mux.HandleFunc("/articles/id", apiCfg.HandlerGetArticleByID)

	wrappedmux := middleware.CorsMiddleware(mux)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", wrappedmux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
