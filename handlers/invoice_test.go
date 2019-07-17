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

func (i InvoiceModel) Create(invoice *models.Invoice) (models.Invoice, error) {
	datetime, _ := time.Parse("2006-01-02T15:04:05-070", "2019-05-02T15:04:05-070")
	return models.Invoice{
		ID:             1,
		ReferenceMonth: 2,
		ReferenceYear:  2017,
		Document:       "0324566545",
		Description:    "Some notes",
		Amount:         38.90,
		IsActive:       1,
		CreatedAt:      datetime,
	}, nil
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

	want := `{"id":1,"reference_month":2,"reference_year":2017,"document":"0324566545","description":"Some notes","amount":38.90,"is_active":1,"created_at":"2019-05-02T15:04:05-07:00", "decativated_at": null}`
	if assert.NoError(t, h.CreateInvoice(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, want+"\n", rec.Body.String())
	}
}
