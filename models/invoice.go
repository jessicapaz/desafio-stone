package models

import (
	"database/sql"
	"fmt"
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
	Document       string     `json:"document" validate:"required,document"`
	Description    string     `json:"description" validate:"required"`
	Amount         float64    `json:"amount"`
	IsActive       int        `json:"is_active"`
	CreatedAt      time.Time  `json:"created_at"`
	DeactivatedAt  *time.Time `json:"deactivated_at,omitempty"`
}

// InvoiceModelImpl describes all methods of a InvoiceModel
type InvoiceModelImpl interface {
	Create(i *Invoice) (Invoice, error)
	List(sort string, offset, limit int) ([]Invoice, error)
	ListByDocument(document, sort string, offset, limit int) ([]Invoice, error)
	ListByMonth(month int, sort string, offset, limit int) ([]Invoice, error)
	ListByYear(year int, sort string, offset, limit int) ([]Invoice, error)
	ByID(id int) (Invoice, error)
	Deactivate(invoice *Invoice) (Invoice, error)
	Update(invoice, newInvoice *Invoice) (Invoice, error)
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
func (i *InvoiceModel) List(sort string, offset, limit int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := fmt.Sprintf("SELECT * FROM invoices ORDER BY %s OFFSET %d LIMIT %d", sort, offset, limit)
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

// ListByDocument list invoices by document
func (i *InvoiceModel) ListByDocument(document, sort string, offset, limit int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := fmt.Sprintf("SELECT * FROM invoices WHERE document='%s' ORDER BY %s OFFSET %d LIMIT %d", document, sort, offset, limit)
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

// ListByMonth list invoices by month
func (i *InvoiceModel) ListByMonth(month int, sort string, offset, limit int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := fmt.Sprintf("SELECT * FROM invoices WHERE reference_month=%d ORDER BY %s OFFSET %d LIMIT %d", month, sort, offset, limit)
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

// ListByYear list invoices by year
func (i *InvoiceModel) ListByYear(year int, sort string, offset, limit int) ([]Invoice, error) {
	invoices := []Invoice{}
	stmt := fmt.Sprintf("SELECT * FROM invoices WHERE reference_year=%d ORDER BY %s OFFSET %d LIMIT %d", year, sort, offset, limit)
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

// ByID get an invoice by id
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

// Update updates a invoice on database
func (i *InvoiceModel) Update(invoice, newInvoice *Invoice) (Invoice, error) {
	stmt := `UPDATE invoices SET reference_month=$1, reference_year=$2, document=$3, description=$4, amount=$5 WHERE id=$6 RETURNING *;`
	result := i.db.QueryRow(stmt, newInvoice.ReferenceMonth, newInvoice.ReferenceYear, newInvoice.Document, newInvoice.Description, newInvoice.Amount, invoice.ID)
	err := result.Scan(&newInvoice.ID, &newInvoice.ReferenceMonth, &newInvoice.ReferenceYear, &newInvoice.Document, &newInvoice.Description, &newInvoice.Amount, &newInvoice.IsActive, &newInvoice.CreatedAt, &newInvoice.DeactivatedAt)
	if err != nil {
		return *newInvoice, err
	}
	return *newInvoice, nil
}
