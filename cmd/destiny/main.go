// Package main ...
package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/subosito/gotenv"

	"github.com/nicoki2004/d2-internal/internal/apicfg"
	"github.com/nicoki2004/d2-internal/internal/config"
	"github.com/nicoki2004/d2-internal/internal/database"
	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/models"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

func main() {
	if err := gotenv.Load(); err != nil {
		logger.GetLogger().Debug("No .env file found")
	}

	log := logger.GetLogger()
	cfg := config.Get()

	dbPath := os.Getenv("DB_NAME")

	dbConn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Error de configuración de DB: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatal("No se pudo conectar a la DB: %v", err)
	}

	dbQueries := database.New(dbConn)
	repo := repository.NewSQLDestinyRepository(dbQueries, dbConn)
	client := destiny.NewClient(cfg, &models.Token{})

	apiCfg := &apicfg.APIConfig{
		DB:     dbQueries,
		Log:    log,
		Repo:   repo,
		Client: *client,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", apiCfg.HandlerHome)
	mux.HandleFunc("GET /login", apiCfg.HandlerLogin)
	mux.HandleFunc("GET /refresh", apiCfg.HandlerGetProfile)
	mux.HandleFunc("GET /characters", apiCfg.HandlerGetCharacters)
	mux.HandleFunc("GET /characters/{id}", apiCfg.HandlerGetCharactersByID)

	server := &http.Server{
		Addr:    ":4200",
		Handler: mux,
	}

	log.Info("🚀 Servidor iniciado en https://localhost:4200")

	if err := server.ListenAndServeTLS("localhost.pem", "localhost-key.pem"); err != nil {
		log.Fatal("Error en el servidor: %v", err)
	}
}
