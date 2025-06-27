package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasramosdev/cnotes/internal/database"
	"github.com/lucasramosdev/cnotes/internal/web"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		err := godotenv.Load()
		if err != nil {
			panic("Error on loading .env")
		}
	}

	ctx := context.Background()

	dbUser := os.Getenv("CNOTES_DBUSER")
	dbPass := os.Getenv("CNOTES_DBPASS")
	db := os.Getenv("CNOTES_DB")

	connectionString := fmt.Sprintf("postgresql://%s:%s@192.168.0.74:5432/%s", dbUser, dbPass, db)
	fmt.Println(connectionString)
	conn, err := database.NewConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := gin.Default()
	router.Static("static", "./web/static")

	web.Configure()
	web.SetRoutes(router)

	router.Run()
}
