package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// NewAccount : Create account repository
func (r *account) Insert(ctx context.Context, account *object.Account) (object.AccountID, error) {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO
			account (
				username,
				password_hash,
				display_name,
				avatar,
				header,
				note
			) VALUES (
				?,
				?,
				?,
				?,
				?,
				?
			)
		`,
		account.Username,
		account.PasswordHash,
		account.DisplayName,
		account.Avatar,
		account.Header,
		account.Note,
	)
	if err != nil {
		log.Default().Printf("failed to insert account: %v", err)
		return 0, fmt.Errorf("%w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Default().Printf("failed to get last insert id: %v", err)
		return 0, fmt.Errorf("%w", err)
	}
	log.Default().Printf("id: %d", id)
	return object.AccountID(id), nil
}
