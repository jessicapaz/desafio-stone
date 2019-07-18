package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/jessicapaz/desafio-stone/models"
)

func (h *Handler) CreateInvoice(c echo.Context) error {
	invoice := &models.Invoice{}
	if err := c.Bind(invoice); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, invoice)
	}
	i, err := h.InvoiceModel.Create(invoice)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, i)
}

func (h *Handler) ListInvoice(c echo.Context) error {
	i, err := h.InvoiceModel.List()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	document := c.QueryParam("document")
	if document != "" {
		i, err := h.InvoiceModel.ByDocument(document)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, i)
	}
	month := c.QueryParam("month")
	if month != "" {
		m, _ := strconv.Atoi(month)
		i, err := h.InvoiceModel.ByMonth(m)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, i)
	}
	year := c.QueryParam("year")
	if year != "" {
		y, _ := strconv.Atoi(year)
		i, err := h.InvoiceModel.ByYear(y)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, i)
}
