package main

import (
	"log"
	"myshelf/config"
	"myshelf/internal/db"
	sqlc "myshelf/internal/db/sqlc/generated"
	"myshelf/internal/modules/post"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := config.LoadConfig()

	dbConn, err := db.MySQLConnect(cfg)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	defer dbConn.Close()

	queries := sqlc.New(dbConn)

	postHandler := post.Handler{
		Service: post.PostService{
			Repo: post.NewPostRepository(queries),
		},
	}

	e.POST("/posts", postHandler.CreatePost)

	e.Logger.Fatal(e.Start(":" + (cfg.Port)))

}
