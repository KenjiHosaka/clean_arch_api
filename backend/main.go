package main

import (
	"clean_arch_api/backend/db"
	"clean_arch_api/backend/environment"
	"clean_arch_api/backend/server"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

// @title Sample API List
// @version 1.0

// @host localhost:8000
// @securityDefinitions.apikey ApiKeyAuth
// @in Headers
// @name Authorization

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("database connect failed")
		return
	}

	defer database.Close()

	s := server.SetUp(database)

	go func() {
		if err := s.Start(":" + environment.Port()); err != nil {
			s.Logger.Info("Shutting down the server")
		}
	}()

	// サーバー終了
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		s.Logger.Fatal(err)
	}
}
