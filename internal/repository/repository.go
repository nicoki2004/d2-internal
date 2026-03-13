// Package repository
package repository

import (
	"context"
	"database/sql"

	"github.com/nicoki2004/d2-internal/internal/database"
)

type SQLDestinyRepository struct {
	queries *database.Queries
	db      *sql.DB
}

func NewSQLDestinyRepository(queries *database.Queries, db *sql.DB) *SQLDestinyRepository {
	return &SQLDestinyRepository{
		queries: queries,
		db:      db,
	}
}

type DestinyRepository interface {
	UpsertUser(ctx context.Context, arg database.UpsertUserParams) error
}

func (r *SQLDestinyRepository) UpsertUser(ctx context.Context, arg database.UpsertUserParams) error {
	return r.queries.UpsertUser(ctx, arg)
}
