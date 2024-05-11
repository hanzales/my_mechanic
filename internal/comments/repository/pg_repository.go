package repository

import (
	"MyMechanic/internal/comments"
	"MyMechanic/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Comments Repository
type commentsRepo struct {
	db *sqlx.DB
}

// Comments Repository constructor
func NewCommentsRepository(db *sqlx.DB) comments.Repository {
	return &commentsRepo{db: db}
}

// GetByID comment
func (r *commentsRepo) GetByID(ctx context.Context, id int) (*models.Comment, error) {

	comment := &models.Comment{}
	if err := r.db.GetContext(ctx, comment, getCommentByID, id); err != nil {
		return nil, errors.Wrap(err, "commentsRepo.GetByID.GetContext")
	}
	return comment, nil
}
