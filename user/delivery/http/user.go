package delivery_http

import (
	"base-go/domain"
	"base-go/domain/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(g *gin.Engine, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}

	g.GET("/user", handler.Fetch)
	g.GET("/user/:user_id", handler.GetByID)
	g.POST("/user", handler.Create)
	g.PUT("/user/:user_id", handler.Update)
	g.DELETE("/user/:user_id", handler.Delete)
}

func (handler *UserHandler) Fetch(c *gin.Context) {
	ctx := c.Request.Context()

	user, err := handler.userUsecase.Fetch(ctx)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "Error",
			},
		)
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}

func (handler *UserHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "internal server error",
			},
		)
		return
	}

	user_id := uint(userIdP)

	user, err := handler.userUsecase.GetByID(ctx, user_id)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "Internal server error",
			},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}

func (handler *UserHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var request domain.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.ResponseError{
				Message: "bad request",
			},
		)

		return
	}

	user, err := handler.userUsecase.Create(ctx, request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "internal server error",
			},
		)
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Code:   201,
		Status: "Created",
		Data:   user,
	})
}

func (handler *UserHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "internal server error",
			},
		)
		return
	}

	user_id := uint(userIdP)

	var request domain.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.ResponseError{
				Message: "bad request",
			},
		)

		return
	}

	request.ID = user_id

	user, err := handler.userUsecase.Update(ctx, request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "internal server error",
			},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}

func (handler *UserHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "internal server error",
			},
		)
		return
	}

	user_id := uint(userIdP)

	err = handler.userUsecase.Delete(ctx, user_id)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Message: "Internal server error",
			},
		)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "OK",
	})
}
