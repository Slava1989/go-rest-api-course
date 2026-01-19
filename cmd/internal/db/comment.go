package db

import (
	"context"
	"database/sql"

	"github.com/Slava1989/go-rest-api-course/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Commnet, error)
