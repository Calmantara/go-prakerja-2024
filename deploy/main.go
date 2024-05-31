package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Calmantara/go-prakerja-2024/deploy/handler"
	"github.com/Calmantara/go-prakerja-2024/deploy/middleware"
	"github.com/Calmantara/go-prakerja-2024/deploy/repository"
)

func main() {
	engine := gin.New()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "I am happy because I am new Daddy",
		},
		)
	})
	// init db
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	// middleware gin gonic
	engine.Use(
		gin.Recovery(),
		// sentry / newrelic / grafana
		// authentication
	)

	// repo
	userRepo := repository.NewUserRepo(db)
	orderRepo := &repository.OrderRepo{DB: db}
	// CICD pipeline sendiri untuk migrate
	userRepo.Migrate()
	orderRepo.Migrate()

	// handler
	userHdl := handler.NewUserHdl(userRepo)
	orderHdl := &handler.OrderHdl{Repository: orderRepo}
	userGroup := engine.Group("/users")
	{
		// Get all user
		userGroup.GET("", userHdl.GetGorm)
		// Create user
		userGroup.Use(middleware.BasicAuthorization())
		userGroup.POST("/registrations", userHdl.CreateGorm)
		userGroup.POST("/logins", userHdl.Login)
		// Update user
		// userGroup.PUT("/:id", userHdl.Update)
	}

	orderGroup := engine.Group("/orders")
	{
		orderGroup.Use(middleware.BearerAuthorization())
		// Get all user
		orderGroup.GET("", orderHdl.GetGorm)
		// Create user
		orderGroup.POST("", orderHdl.CreateGorm)
		// Update user
		// userGroup.PUT("/:id", userHdl.Update)
	}
	appport := os.Getenv("PORT")
	err = engine.Run(":" + appport)
	if err != nil {
		log.Fatal(err)
	}
}
