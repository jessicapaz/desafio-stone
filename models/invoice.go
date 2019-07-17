package models

import (
	"database/sql"
	"time"
)

// Invoice model
type Invoice struct {
	ID             int        `json:"id"`
	ReferenceMonth int        `json:"reference_month"`
	ReferenceYear  int        `json:"reference_year"`
	Document       string     `json:"document"`
	Description    string     `json:"description"`
	Amount         float64    `json:"amount"`
	IsActive       int        `json:"is_active"`
	CreatedAt      time.Time  `json:"created_at"`
	DeactivatedAt  *time.Time `json:"deactivated_at"`
}

// InvoiceModelImpl describes all methods of a InvoiceModel
type InvoiceModelImpl interface {
	Create(i *Invoice) (Invoice, error)
	List() ([]Invoice, error)
	ByDocument(document string) ([]Invoice, error)
}

type InvoiceModel struct {
	db *sql.DB
}

// NewInvoiceModel creates a new InvoiceModel
func NewInvoiceModel(db *sql.DB) *InvoiceModel {
	return &InvoiceModel{
		db: db,
	}
}

// Create creates a invoice on database
func (i *InvoiceModel) Create(invoice *Invoice) (Invoice, error) {
	newInvoice := Invoice{}
	stmt := `INSERT INTO invoices (
		reference_month,
		reference_year,
		document,
		description,
		amount,
		is_active,
		created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`
	result := i.db.QueryRow(stmt, invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, 1, time.Now())
	err := result.Scan(&newInvoice.ID, &newInvoice.ReferenceMonth, &newInvoice.ReferenceYear, &newInvoice.Document, &newInvoice.Description, &newInvoice.Amount, &newInvoice.IsActive, &newInvoice.CreatedAt, &newInvoice.DeactivatedAt)
	if err != nil {
		return newInvoice, err
	}
	return newInvoice, nil
}

// List list all invoices
func (i *InvoiceModel) List() ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := `SELECT * FROM invoices;`
	result, err := i.db.Query(stmt)
	if err != nil {
		return invoices, err
	}
	defer result.Close()

	for result.Next() {
		invoice := Invoice{}
		err := result.Scan(&invoice.ID, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
		if err != nil {
			return invoices, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (i *InvoiceModel) ByDocument(document string) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := `SELECT * FROM invoices WHERE document=$1;`
	result, err := i.db.Query(stmt, document)
	if err != nil {
		return invoices, err
	}
	defer result.Close()

	for result.Next() {
		invoice := Invoice{}
		err := result.Scan(&invoice.ID, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
		if err != nil {
			return invoices, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}
