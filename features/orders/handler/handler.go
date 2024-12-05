package handler

import (
	"floweys_app/app/helpers"
	"floweys_app/features/orders"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
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

func (handler *OrderHandler) Delete(c echo.Context) error {
	id := c.Param("order_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	err := handler.orderService.Remove(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "operation failed, data not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *OrderHandler) Get(c echo.Context) error {
	id := c.Param("order_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	result, err := handler.orderService.Read(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrDataNotFound.Error(), nil))
	}
	resultResponse := OrderCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success get data", resultResponse))
}

func (handler *OrderHandler) Update(c echo.Context) error {
	id := c.Param("order_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	var orderInput OrderRequest
	errBind := c.Bind(&orderInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	updatedOrder := OrderRequestToCore(orderInput)

	err := handler.orderService.Edit(uint(idConv), updatedOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error update data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update data", nil))
}

func New(service orders.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}
