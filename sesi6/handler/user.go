package handler

import (
	"fmt"
	"net/http"

	"github.com/Calmantara/go-prakerja-2024/sesi6/model"
	"github.com/Calmantara/go-prakerja-2024/sesi6/repository"
	"github.com/gin-gonic/gin"
)

type UserHdl struct {
	Repository *repository.UserRepo
}

func (u *UserHdl) GetGorm(ctx *gin.Context) {
	users, err := u.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserHdl) CreateGorm(ctx *gin.Context) {
	// bind body
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	// validate payload
	if err := user.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": err,
		})
		return
	}
	err := u.Repository.Create(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserHdl) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Users)
}

func (u *UserHdl) Create(ctx *gin.Context) {
	// bind body
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	// validate payload
	if err := user.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": err,
		})
		return
	}
	user.ID = uint64(len(model.Users) + 1)
	model.Users = append(model.Users, user)
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserHdl) Update(ctx *gin.Context) {
	// getid from param
	userID := ctx.Param("id")
	if userID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid id",
		})
		return
	}

	// bind body
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	// validate payload
	if err := user.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": err,
		})
		return
	}
	// check user exist
	isExist, userIdx := false, 0
	for idx, usr := range model.Users {
		if fmt.Sprintf("%v", usr.ID) == userID {
			isExist = true
			userIdx = idx
			user.ID = usr.ID
			break
		}
	}
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"message": "user with id not found",
		})
		return
	}
	// update user
	model.Users[userIdx] = user
}
