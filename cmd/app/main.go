package main

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasramosdev/cnotes/internal/database"
	"github.com/lucasramosdev/cnotes/internal/http"
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

	connectionString := fmt.Sprintf("postgresql://%s:%s@db:5432/%s", dbUser, dbPass, db)
	conn, err := database.NewConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := gin.Default()
	router.SetHTMLTemplate(template.Must(template.ParseGlob("web/templates/*.tmpl")))
	router.Static("static", "./web/static")

	http.SetRoutes(router)
	router.Run()
}
