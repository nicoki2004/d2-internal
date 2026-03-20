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
	GetUser(ctx context.Context, membershipID string) (database.User, error)
	UpsertWeapon(ctx context.Context, params database.UpsertWeaponParams) error
	ClearWeaponStats(ctx context.Context, instanceID string) error
	InsertWeaponStat(ctx context.Context, params database.InsertWeaponStatParams) error
	ClearWeaponPerks(ctx context.Context, instanceID string) error
	InsertWeaponPerk(ctx context.Context, params database.InsertWeaponPerkParams) error

	GetCharactersWithStats(ctx context.Context) ([]database.GetCharactersWithStatsRow, error)

	// Save a Character FULL
	SaveCharacterFull(ctx context.Context, dataCharacter database.UpsertCharacterParams, dataStats []database.UpsertCharacterStatParams) error

	GetAllWeaponsWithPerks(ctx context.Context) ([]database.GetAllWeaponsWithPerksRow, error)
}

func (r *SQLDestinyRepository) UpsertUser(ctx context.Context, arg database.UpsertUserParams) error {
	return r.queries.UpsertUser(ctx, arg)
}

func (r *SQLDestinyRepository) GetUser(ctx context.Context, membershipID string) (database.User, error) {
	return r.queries.GetUser(ctx, membershipID)
}

func (r *SQLDestinyRepository) UpsertWeapon(ctx context.Context, params database.UpsertWeaponParams) error {
	return r.queries.UpsertWeapon(ctx, params)
}

func (r *SQLDestinyRepository) ClearWeaponStats(ctx context.Context, instanceID string) error {
	return r.queries.ClearWeaponStats(ctx, instanceID)
}

func (r *SQLDestinyRepository) InsertWeaponStat(ctx context.Context, params database.InsertWeaponStatParams) error {
	return r.queries.InsertWeaponStat(ctx, params)
}

func (r *SQLDestinyRepository) ClearWeaponPerks(ctx context.Context, instanceID string) error {
	return r.queries.ClearWeaponPerks(ctx, instanceID)
}

func (r *SQLDestinyRepository) InsertWeaponPerk(ctx context.Context, params database.InsertWeaponPerkParams) error {
	return r.queries.InsertWeaponPerk(ctx, params)
}

func (r *SQLDestinyRepository) SaveCharacterFull(
	ctx context.Context,
	data database.UpsertCharacterParams,
	stats []database.UpsertCharacterStatParams,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var titleHash int64
	if h, ok := data.TitleRecordHash.(float64); ok {
		titleHash = int64(h)
	} else if h, ok := data.TitleRecordHash.(int64); ok {
		titleHash = h
	}

	err = r.queries.WithTx(tx).UpsertCharacter(ctx, database.UpsertCharacterParams{
		CharacterID:          data.CharacterID,
		ClassType:            int64(data.ClassType),
		LightLevel:           int64(data.LightLevel),
		EmblemUrl:            sql.NullString{String: data.EmblemUrl.String, Valid: true},
		LastPlayed:           sql.NullTime{Time: data.LastPlayed.Time, Valid: true},
		EmblemBackgroundPath: sql.NullString{String: data.EmblemBackgroundPath.String, Valid: true},
		TitleRecordHash:      int64(titleHash),
		EmblemColorR:         sql.NullInt64{Int64: int64(data.EmblemColorR.Int64), Valid: true},
		EmblemColorG:         sql.NullInt64{Int64: int64(data.EmblemColorG.Int64), Valid: true},
		EmblemColorB:         sql.NullInt64{Int64: int64(data.EmblemColorB.Int64), Valid: true},
		EmblemColorA:         sql.NullInt64{Int64: int64(data.EmblemColorA.Int64), Valid: true},
	})
	if err != nil {
		return err
	}

	// 2. Upsert de cada stat
	for _, val := range stats {

		var stahsHash int64
		if h, ok := val.StatHash.(int64); ok {
			stahsHash = int64(h)
		} else {
			stahsHash = 0
		}

		err = r.queries.WithTx(tx).UpsertCharacterStat(ctx, database.UpsertCharacterStatParams{
			CharacterID: data.CharacterID,
			StatHash:    int64(stahsHash),
			Value:       int64(val.Value),
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *SQLDestinyRepository) GetCharactersWithStats(ctx context.Context) ([]database.GetCharactersWithStatsRow, error) {
	return r.queries.GetCharactersWithStats(ctx)
}

func (r *SQLDestinyRepository) GetAllWeaponsWithPerks(ctx context.Context) ([]database.GetAllWeaponsWithPerksRow, error) {
	return r.queries.GetAllWeaponsWithPerks(ctx)
}
