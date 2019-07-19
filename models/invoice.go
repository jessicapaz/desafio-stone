package models

import (
	"database/sql"
	"time"
)

const (
	active      = 1
	deactivated = 0
)

// Invoice model
type Invoice struct {
	ID             int        `json:"id"`
	ReferenceMonth int        `json:"reference_month" validate:"required,min=1,max=12"`
	ReferenceYear  int        `json:"reference_year" validate:"required,min=1900,max=2030"`
	Document       string     `json:"document" validate:"required,len=11|len=14"`
	Description    string     `json:"description" validate:"required"`
	Amount         float64    `json:"amount" validate:"required,gt=0"`
	IsActive       int        `json:"is_active"`
	CreatedAt      time.Time  `json:"created_at"`
	DeactivatedAt  *time.Time `json:"deactivated_at,omitempty"`
}

// InvoiceModelImpl describes all methods of a InvoiceModel
type InvoiceModelImpl interface {
	Create(i *Invoice) (Invoice, error)
	List() ([]Invoice, error)
	ByDocument(document string) ([]Invoice, error)
	ByMonth(month int) ([]Invoice, error)
	ByYear(year int) ([]Invoice, error)
	ByID(id int) (Invoice, error)
	Deactivate(invoice *Invoice) (Invoice, error)
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
	result := i.db.QueryRow(stmt, invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, active, time.Now())
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

// ByDocument list invoices by document
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

// ByMonth list invoices by month
func (i *InvoiceModel) ByMonth(month int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := `SELECT * FROM invoices WHERE reference_month=$1;`
	result, err := i.db.Query(stmt, month)
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

// ByYear list invoices by year
func (i *InvoiceModel) ByYear(year int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := `SELECT * FROM invoices WHERE reference_year=$1;`
	result, err := i.db.Query(stmt, year)
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

// ByID get an invoices by id
func (i *InvoiceModel) ByID(id int) (Invoice, error) {
	invoice := Invoice{}
	stmt := `SELECT * FROM invoices WHERE id=$1;`
	result := i.db.QueryRow(stmt, id)
	err := result.Scan(&invoice.ID, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

// Deactivate change invoice status from activated to deactivated
func (i *InvoiceModel) Deactivate(invoice *Invoice) (Invoice, error) {
	returnInvoice := Invoice{}
	stmt := `UPDATE invoices SET is_active=$1, deactivated_at=$2 WHERE id=$3 RETURNING *;`
	result := i.db.QueryRow(stmt, deactivated, time.Now(), invoice.ID)
	err := result.Scan(&returnInvoice.ID, &returnInvoice.ReferenceMonth, &returnInvoice.ReferenceYear, &returnInvoice.Document, &returnInvoice.Description, &returnInvoice.Amount, &returnInvoice.IsActive, &returnInvoice.CreatedAt, &returnInvoice.DeactivatedAt)
	if err != nil {
		return returnInvoice, err
	}
	return returnInvoice, nil
}
