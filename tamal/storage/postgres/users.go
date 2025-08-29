package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"save-tamal/tamal/storage"
)

const insertUser = `
INSERT INTO users (
	name, 
	batch, 
	email,
	password,
	created_by,
	updated_by
) VALUES (
	:name, 
	:batch,
	:email,
	:password,
	:created_by,
	:updated_by
) RETURNING
	user_id;
`

func (s *Storage) CreateUser(ctx context.Context, user storage.User) (string, error) {
	stmt, err := s.db.PrepareNamed(insertUser)
	if err != nil {
		return "", err
	}

	var id string
	if err := stmt.Get(&id, user); err != nil {
		return "", err
	}

	return id, nil
}

const getUser = `
SELECT *
FROM users
WHERE (user_id = $1 OR email = $2) AND deleted_at IS NULL; 
`

func (s *Storage) GetUser(ctx context.Context, user storage.User) (*storage.User, error) {
	var res storage.User
	if err := s.db.Get(&res, getUser, user.UserID, user.Email); err != nil {
		return nil, fmt.Errorf("executing user details: %w", err)
	}
	return &res, nil
}

const deleteUser = `
UPDATE
	users
SET
	deleted_at = now(),
	deleted_by = $1
WHERE 
	user_id = $2;

`

func (s *Storage) DeleteUser(ctx context.Context, user storage.User) error {
	_, err := s.db.Exec(deleteUser, user.DeletedBy, user.UserID)
	if err != nil {
		return err
	}
	return nil
}

const updateUser = `
UPDATE
	 users
SET
	name = :name,
    batch = :batch,
	email = :email,
	password = :password,
	updated_at = now(),
	updated_by = :updated_by
WHERE 
	user_id = :user_id
RETURNING 
	updated_at;
`

func (s *Storage) UpdateUser(ctx context.Context, user storage.User) (*storage.User, error) {
	stmt, err := s.db.PrepareNamedContext(ctx, updateUser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	if err := stmt.Get(&user, user); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("executing user update: %w", err)
	}
	return &user, nil
}

const listUser = `
SELECT * FROM users WHERE deleted_at IS NULL AND name ILIKE '%%' || '$1' || '%%' ORDER BY $2 $3 LIMIT NULLIF($4, 0) OFFSET $5";
`

func (s *Storage) ListUser(ctx context.Context, f storage.Filter) ([]storage.User, error) {
	var users []storage.User
	if err := s.db.Select(&users, listUser, f.SearchTerm, f.SortBy, f.Order, f.Limit, f.Offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return users, nil
}

const userStat = `
SELECT COUNT(*) FROM users where deleted_at IS NULL AND name ILIKE '%%' || '$1' || '%%';
`

func (s *Storage) UserStats(ctx context.Context, f storage.Filter) (storage.Stats, error) {
	var stat storage.Stats
	if err := s.db.Select(&stat, userStat, f.SearchTerm); err != nil {
		if err == sql.ErrNoRows {
			return storage.Stats{}, err
		}
		return storage.Stats{}, err
	}

	return stat, nil
}
