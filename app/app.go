package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SwArch-2025-1-2A/users_ms/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DBPool  *pgxpool.Pool
	Queries *repository.Queries
	Router  *gin.Engine
	Context context.Context
}

func NewApp() *App {

	dbURL := os.Getenv("DATABASE_URL")
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("connection pool could not be created: %w", err))
	}

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatal(fmt.Errorf("could not connect to the DB, %w", err))
	}

	log.Println("Successfully connected to the Main DB")

	queries := repository.New(dbpool)

	return &App{
		DBPool:  dbpool,
		Queries: queries,
		Router:  gin.Default(),
		Context: ctx,
	}
}
