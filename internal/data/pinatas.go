package data

import (
	"PinataService.alikhankaliyev.net/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"time"
)

type Pinata struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Color      string    `json:"color,omitempty"`
	Shape      string    `json:"shape,omitempty"`
	Contents   []string  `json:"contents,omitempty"`
	Weight     Weight    `json:"weight,omitempty,string"`
	Dimensions struct {
		Height float32 `json:"height,string"`
		Width  float32 `json:"width,string"`
		Depth  float32 `json:"depth,string"`
	} `json:"dimensions,omitempty"`
	Version int64 `json:"version"`
}

func ValidatePinata(v *validator.Validator, pinata *Pinata) {
	v.Check(pinata.Color != "", "color", "must be provided")
	v.Check(len(strings.Split(pinata.Color, " ")) > 0, "color", "there must be one word")

	v.Check(pinata.Contents != nil, "contents", "must be provided")
	v.Check(len(pinata.Contents) >= 1, "contents", "must contain at least 1 content")
	v.Check(len(pinata.Contents) <= 100, "contents", "must not contain more than 100 contents")

	v.Check(validator.Unique(pinata.Contents), "contents", "must not contain duplicate values")
}

type PinataModel struct {
	DB *sql.DB
}

func (m PinataModel) Insert(pinata *Pinata) error {
	query := `
		INSERT INTO pinatas (color, shape, contents, weight, height, width, depth)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, version`
	args := []interface{}{pinata.Color, pinata.Shape, pq.Array(pinata.Contents), pinata.Weight, pinata.Dimensions.Height,
		pinata.Dimensions.Width, pinata.Dimensions.Depth}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&pinata.ID, &pinata.CreatedAt, &pinata.Version)
}

func (m PinataModel) Get(id int64) (*Pinata, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, color, shape , contents, weight, height, width, depth, version
		FROM pinatas
		WHERE id = $1`

	var pinata Pinata

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&pinata.ID,
		&pinata.CreatedAt,
		&pinata.Color,
		&pinata.Shape,
		pq.Array(&pinata.Contents),
		&pinata.Weight,
		&pinata.Dimensions.Height,
		&pinata.Dimensions.Width,
		&pinata.Dimensions.Depth,
		&pinata.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &pinata, err
}

func (m PinataModel) Update(pinata *Pinata) error {
	fmt.Print(&pinata)
	query := `
		UPDATE pinatas
		SET color = $1, shape = $2,  contents = $3, weight = $4, height = $5, width = $6, depth = $7, version = version + 1   
		WHERE id = $8 AND version = $9
		RETURNING version`

	args := []interface{}{
		pinata.Color,
		pinata.Shape,
		pq.Array(pinata.Contents),
		pinata.Weight,
		pinata.Dimensions.Height,
		pinata.Dimensions.Width,
		pinata.Dimensions.Depth,
		pinata.ID,
		pinata.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&pinata.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m PinataModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM pinatas WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
