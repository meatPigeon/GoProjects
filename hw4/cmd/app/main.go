package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"hw4/internal/handler"
	"hw4/internal/server"
	"hw4/internal/service"
	"hw4/internal/storage/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := postgres.NewRepository(conn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)
	srv := server.NewServer()

	e := echo.New()
	srv.MapRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
