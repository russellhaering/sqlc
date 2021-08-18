// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgtype"
)

const generateSeries = `-- name: GenerateSeries :many
SELECT ($1::inet) + i
FROM generate_series(0, $2::int) AS i
LIMIT 1
`

type GenerateSeriesParams struct {
	Column1 pgtype.Inet
	Column2 int32
}

func (q *Queries) GenerateSeries(ctx context.Context, arg GenerateSeriesParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, generateSeries, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var column_1 int32
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name
FROM users_func()
WHERE first_name != ''
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.FirstName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
