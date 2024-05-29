package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Calmantara/go-prakerja-2024/sesi6/handler"
	"github.com/Calmantara/go-prakerja-2024/sesi6/repository"
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
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=35432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	// repo
	userRepo := &repository.UserRepo{DB: db}
	orderRepo := &repository.OrderRepo{DB: db}
	// CICD pipeline sendiri untuk migrate
	userRepo.Migrate()
	orderRepo.Migrate()

	// handler
	userHdl := &handler.UserHdl{Repository: userRepo}
	orderHdl := &handler.OrderHdl{Repository: orderRepo}
	userGroup := engine.Group("/users")
	{
		// Get all user
		userGroup.GET("", userHdl.GetGorm)
		// Create user
		userGroup.POST("", userHdl.CreateGorm)
		// Update user
		// userGroup.PUT("/:id", userHdl.Update)
	}

	orderGroup := engine.Group("/orders")
	{
		// Get all user
		orderGroup.GET("", orderHdl.GetGorm)
		// Create user
		orderGroup.POST("", orderHdl.CreateGorm)
		// Update user
		// userGroup.PUT("/:id", userHdl.Update)
	}
	err = engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
