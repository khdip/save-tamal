package storage

import (
	"database/sql"
	"time"
)

type User struct {
	UserID          string `db:"user_id"`
	Name            string `db:"name"`
	Batch           int    `db:"batch"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	ConfirmPassword string
	CRUDTimeDate
	Filter
}

type Collection struct {
	CollectionID  int       `db:"collection_id"`
	AccountType   string    `db:"account_type"`
	AccountNumber string    `db:"account_Number"`
	Sender        string    `db:"sender"`
	Date          time.Time `db:"date"`
	Amount        int       `db:"amount"`
	CRUDTimeDate
	Filter
}

type CRUDTimeDate struct {
	CreatedAt time.Time      `db:"created_at,omitempty"`
	CreatedBy string         `db:"created_by"`
	UpdatedAt time.Time      `db:"updated_at,omitempty"`
	UpdatedBy string         `db:"updated_by,omitempty"`
	DeletedAt sql.NullTime   `db:"deleted_at,omitempty"`
	DeletedBy sql.NullString `db:"deleted_by,omitempty"`
}

type Filter struct {
	Offset     int32
	Limit      int32
	Count      int
	SortBy     string
	Order      string
	SearchTerm string
}
