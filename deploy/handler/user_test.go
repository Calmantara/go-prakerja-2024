package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Calmantara/go-prakerja-2024/deploy/mocks"
	"github.com/Calmantara/go-prakerja-2024/deploy/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserGetGorm(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// mock seakan akan ada request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(rec)
	g.Request = req

	// mock
	userRepoMock := mocks.NewUserRepo(t)
	userRepoMock.On("Get").Return([]*model.User{{ID: 1, Username: "test"}}, nil)

	usrHdl := NewUserHdl(userRepoMock)
	usrHdl.GetGorm(g)
	// check body response
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
}

func TestUserGetGormError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// mock seakan akan ada request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(rec)
	g.Request = req

	// mock
	userRepoMock := mocks.NewUserRepo(t)
	userRepoMock.On("Get").Return(nil, errors.New("some error"))

	usrHdl := NewUserHdl(userRepoMock)
	usrHdl.GetGorm(g)
	// check body response
	assert.Equal(t, http.StatusInternalServerError, rec.Result().StatusCode)
}
