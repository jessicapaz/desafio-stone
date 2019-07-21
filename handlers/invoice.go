package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/renderings"
)

// CreateInvoice handler
func (h *Handler) CreateInvoice(c echo.Context) error {
	invoice := &models.Invoice{}
	e := renderings.ErrorResponse{}
	if err := c.Bind(invoice); err != nil {
		e.Errors = []string{"Invalid request"}
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	if err := Validate(invoice); len(err) != 0 {
		e.Errors = err
		return c.JSON(http.StatusBadRequest, e)
	}
	i, err := h.InvoiceModel.Create(invoice)
	if err != nil {
		e.Errors = []string{err.Error()}
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusCreated, i)
}

// ListInvoice handler
func (h *Handler) ListInvoice(c echo.Context) error {
	sort := c.QueryParam("sort")
	if err := ValidateSortQueryParam(sort); err != nil {
		sort = "id"
	}
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	i, err := h.InvoiceModel.List(sort, offset, limit)
	e := renderings.ErrorResponse{}
	if err != nil {
		e.Errors = []string{err.Error()}
		return c.JSON(http.StatusInternalServerError, e)
	}
	document := c.QueryParam("document")
	if document != "" {
		i, err := h.InvoiceModel.ListByDocument(document, sort, offset, limit)
		if err != nil {
			e.Errors = []string{err.Error()}
			return c.JSON(http.StatusInternalServerError, e)
		}
		return c.JSON(http.StatusOK, i)
	}
	month, _ := strconv.Atoi(c.QueryParam("reference_month"))
	if month != 0 {
		i, err := h.InvoiceModel.ListByMonth(month, sort, offset, limit)
		if err != nil {
			e.Errors = []string{err.Error()}
			return c.JSON(http.StatusInternalServerError, e)
		}
		return c.JSON(http.StatusOK, i)
	}
	year, _ := strconv.Atoi(c.QueryParam("reference_year"))
	if year != 0 {
		i, err := h.InvoiceModel.ListByYear(year, sort, offset, limit)
		if err != nil {
			e.Errors = []string{err.Error()}
			return c.JSON(http.StatusInternalServerError, e)
		}
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, i)
}

// DeactivateInvoice handler
func (h *Handler) DeactivateInvoice(c echo.Context) error {
	e := renderings.ErrorResponse{}
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.InvoiceModel.ByID(id)
	if err != nil {
		e.Errors = []string{"Invoice not found"}
		return c.JSON(http.StatusNotFound, e)
	}
	i, err := h.InvoiceModel.Deactivate(&invoice)
	if err != nil {
		e.Errors = []string{err.Error()}
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusOK, i)
}

// RetrieveInvoice handler
func (h *Handler) RetrieveInvoice(c echo.Context) error {
	e := renderings.ErrorResponse{}
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.InvoiceModel.ByID(id)
	if err != nil {
		e.Errors = []string{"Invoice not found"}
		return c.JSON(http.StatusNotFound, e)
	}
	return c.JSON(http.StatusOK, invoice)
}

// UpdateInvoice handler
func (h *Handler) UpdateInvoice(c echo.Context) error {
	e := renderings.ErrorResponse{}
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.InvoiceModel.ByID(id)
	if err != nil {
		e.Errors = []string{"Invoice not found"}
		return c.JSON(http.StatusNotFound, e)
	}
	newInvoice := &models.Invoice{}
	if err := c.Bind(newInvoice); err != nil {
		e.Errors = []string{"Invalid request"}
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	if err := Validate(newInvoice); len(err) != 0 {
		e.Errors = err
		return c.JSON(http.StatusBadRequest, e)
	}
	i, err := h.InvoiceModel.Update(&invoice, newInvoice)
	if err != nil {
		e.Errors = []string{err.Error()}
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusOK, i)
}

// PartialUpdateInvoice handler
func (h *Handler) PartialUpdateInvoice(c echo.Context) error {
	e := renderings.ErrorResponse{}
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.InvoiceModel.ByID(id)
	if err != nil {
		e.Errors = []string{"Invoice not found"}
		return c.JSON(http.StatusNotFound, e)
	}
	newInvoice := &models.Invoice{}
	if err := c.Bind(newInvoice); err != nil {
		e.Errors = []string{"Invalid request"}
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	i, err := h.InvoiceModel.PartialUpdate(&invoice, newInvoice)
	if err != nil {
		e.Errors = []string{err.Error()}
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusOK, i)
}
