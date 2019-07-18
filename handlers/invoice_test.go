package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/jessicapaz/desafio-stone/models"
)

type InvoiceModel struct{}

var (
	datetime, _    = time.Parse("2006-01-02T15:04:05-070", "2019-05-02T15:04:05-070")
	invoice_data_1 = models.Invoice{
		ID:             1,
		ReferenceMonth: 2,
		ReferenceYear:  2017,
		Document:       "0324566545",
		Description:    "Some notes",
		Amount:         38.90,
		IsActive:       1,
		CreatedAt:      datetime,
	}
	invoice_data_2 = models.Invoice{
		ID:             2,
		ReferenceMonth: 3,
		ReferenceYear:  2018,
		Document:       "0324566548",
		Description:    "Some notes",
		Amount:         38.90,
		IsActive:       1,
		CreatedAt:      datetime,
	}
)

func (i InvoiceModel) Create(invoice *models.Invoice) (models.Invoice, error) {
	return invoice_data_1, nil
}

func (i InvoiceModel) List() ([]models.Invoice, error) {
	return []models.Invoice{invoice_data_1, invoice_data_2}, nil
}

func (i InvoiceModel) ByDocument(document string) ([]models.Invoice, error) {
	if document == "0324566545" {
		return []models.Invoice{invoice_data_1}, nil
	}
	return []models.Invoice{invoice_data_2}, nil
}

func (i InvoiceModel) ByMonth(month int) ([]models.Invoice, error) {
	if month == 2 {
		return []models.Invoice{invoice_data_1}, nil
	}
	return []models.Invoice{invoice_data_2}, nil
}

func (i InvoiceModel) ByYear(year int) ([]models.Invoice, error) {
	if year == 2017 {
		return []models.Invoice{invoice_data_1}, nil
	}
	return []models.Invoice{invoice_data_2}, nil
}

func TestCreateInvoice(t *testing.T) {
	e := echo.New()
	invoiceJSON := `{"reference_month":2,"reference_year":2017,"document":"0324566545",
		"description":"Some notes", "amount":38.90,"is_active":1}`
	req := httptest.NewRequest(http.MethodPost, "/invoices", strings.NewReader(invoiceJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	i := &InvoiceModel{}
	h := NewHandler(nil, nil, i)

	want := `{"id":1,"reference_month":2,"reference_year":2017,"document":"0324566545","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null}`
	if assert.NoError(t, h.CreateInvoice(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, want+"\n", rec.Body.String())
	}
}

func TestListInvoice(t *testing.T) {
	t.Run("Returns all invoices", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/invoices", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		i := &InvoiceModel{}
		h := NewHandler(nil, nil, i)

		want := `[{"id":1,"reference_month":2,"reference_year":2017,"document":"0324566545","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null},{"id":2,"reference_month":3,"reference_year":2018,"document":"0324566548","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null}]`
		if assert.NoError(t, h.ListInvoice(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})

	t.Run("Returns invoices of a given document", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/invoices?document=0324566548", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		i := &InvoiceModel{}
		h := NewHandler(nil, nil, i)

		want := `[{"id":2,"reference_month":3,"reference_year":2018,"document":"0324566548","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null}]`
		if assert.NoError(t, h.ListInvoice(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})

	t.Run("Returns invoices of a given month", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/invoices?month=2", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		i := &InvoiceModel{}
		h := NewHandler(nil, nil, i)

		want := `[{"id":1,"reference_month":2,"reference_year":2017,"document":"0324566545","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null}]`
		if assert.NoError(t, h.ListInvoice(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})

	t.Run("Returns invoices of a given year", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/invoices?year=2017", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		i := &InvoiceModel{}
		h := NewHandler(nil, nil, i)

		want := `[{"id":1,"reference_month":2,"reference_year":2017,"document":"0324566545","description":"Some notes","amount":38.9,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00","deactivated_at":null}]`
		if assert.NoError(t, h.ListInvoice(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})
}
