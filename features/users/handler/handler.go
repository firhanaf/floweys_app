package handler

import (
	"floweys_app/app/helpers"
	"floweys_app/features/users"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrLogin.Error(), nil))
	}
	dataLogin, token, err := handler.userService.Login(userInput.Username, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrLogin.Error(), nil))
		}
	}
	var response = LoginResponse{
		ID:       dataLogin.ID,
		Username: dataLogin.Username,
		Token:    token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) Register(c echo.Context) error {
	var userInput RegisterRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBindData.Error(), nil))
	}
	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	userCore := RegisterRequestToCore(userInput)
	err := handler.userService.Add(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadGateway, helpers.ErrBadRequest.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
		}
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success register", nil))
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
