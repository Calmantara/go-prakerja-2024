package handler

import (
	"net/http"

	"github.com/Calmantara/go-prakerja-2024/sesi7/model"
	"github.com/Calmantara/go-prakerja-2024/sesi7/repository"
	"github.com/gin-gonic/gin"
)

type OrderHdl struct {
	Repository *repository.OrderRepo
}

func (u *OrderHdl) GetGorm(ctx *gin.Context) {
	orders, err := u.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (u *OrderHdl) CreateGorm(ctx *gin.Context) {
	// bind body
	order := &model.Order{}
	if err := ctx.Bind(order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	err := u.Repository.Create(order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}
