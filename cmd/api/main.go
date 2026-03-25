package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/GabrielBrotas/go-categories-msvc/docs"
	c_routes "github.com/GabrielBrotas/go-categories-msvc/cmd/api/routes/categories"
	infraDb "github.com/GabrielBrotas/go-categories-msvc/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Categories API
// @version         1.0
// @description     API para gerenciamento de categorias
// @host            localhost:8080
// @BasePath        /
func main() {
	err := godotenv.Load()
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	db, err := infraDb.InitDb()

	if err != nil {
		panic(err)
	}

	infraDb.MigrateModels(db)

	c_routes.CategoryRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
