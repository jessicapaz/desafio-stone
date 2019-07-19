package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/jessicapaz/desafio-stone/models"
)

// CreateInvoice handler
func (h *Handler) CreateInvoice(c echo.Context) error {
	invoice := &models.Invoice{}
	if err := c.Bind(invoice); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid request")
	}
	if err := c.Validate(invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	i, err := h.InvoiceModel.Create(invoice)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, i)
}

// ListInvoice handler
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

// DeactivateInvoice handler
func (h *Handler) DeactivateInvoice(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	invoice, err := h.InvoiceModel.ByID(idInt)
	if err != nil {
		return c.String(http.StatusNotFound, "Invoice not found")
	}
	i, err := h.InvoiceModel.Deactivate(&invoice)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, i)
}

// RetrieveInvoice handler
func (h *Handler) RetrieveInvoice(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	invoice, err := h.InvoiceModel.ByID(idInt)
	if err != nil {
		return c.String(http.StatusNotFound, "Invoice not found")
	}
	return c.JSON(http.StatusOK, invoice)
}

// UpdateInvoice handler
func (h *Handler) UpdateInvoice(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	invoice, err := h.InvoiceModel.ByID(idInt)
	if err != nil {
		return c.String(http.StatusNotFound, "Invoice not found")
	}
	newInvoice := &models.Invoice{}
	if err := c.Bind(newInvoice); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid request")
	}
	i, err := h.InvoiceModel.Update(&invoice, newInvoice)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, i)
}
