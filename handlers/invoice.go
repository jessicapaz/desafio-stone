package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/labstack/echo"
	"net/http"
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
		in, err := h.InvoiceModel.ByDocument(document)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, in)
	}
	return c.JSON(http.StatusOK, i)
}