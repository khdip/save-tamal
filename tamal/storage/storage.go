package storage

import (
	"database/sql"
	"time"
)

type User struct {
	UserID   string `db:"user_id"`
	Name     string `db:"name"`
	Batch    int32  `db:"batch"`
	Email    string `db:"email"`
	Password string `db:"password"`
	CRUDTimeDate
}

type Collection struct {
	CollectionID  int32  `db:"collection_id"`
	AccountType   string `db:"account_type"`
	AccountNumber string `db:"account_number"`
	Sender        string `db:"sender"`
	Date          string `db:"date"`
	Amount        int32  `db:"amount"`
	Currency      string `db:"currency"`
	CRUDTimeDate
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
	SortBy     string
	Order      string
	SearchTerm string
}

type Stats struct {
	Count       int32
	TotalAmount int32 `db:"sum"`
}
