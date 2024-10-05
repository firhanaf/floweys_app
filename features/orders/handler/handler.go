package handler

import (
	"floweys_app/app/helpers"
	"floweys_app/features/orders"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderHandler struct {
	orderService orders.OrderServiceInterface
}

func (handler *OrderHandler) Create(c echo.Context) error {
	var userInput OrderRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	orderCore := OrderRequestToCore(userInput)
	err := handler.orderService.Add(orderCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success create order", nil))
}

func New(service orders.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}
