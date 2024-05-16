package repository

import (
	"MyMechanic/internal/comments"
	"MyMechanic/internal/models"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Comments Repository
type commentsRepo struct {
	db *sqlx.DB
}

// Comments Repository constructor
func CommentsRepository(db *sqlx.DB) comments.Repository {
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

func (r *commentsRepo) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, deleteComment, id)
	if err != nil {
		return errors.Wrap(err, "commentsRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return errors.Wrap(err, "commentsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "commentsRepo.Delete.rowsAffected")
	}

	return nil
}

func (r *commentsRepo) Create(ctx context.Context, addCommentRequest *models.AddCommentRequest) (*models.Comment, error) {
	comment := &models.Comment{}

	if err := r.db.QueryRowxContext(
		ctx,
		createComment,
		addCommentRequest.Message,
		addCommentRequest.UserId,
		addCommentRequest.DemandId,
	).StructScan(comment); err != nil {
		return nil, errors.Wrap(err, "commentsRepo.Create.StructScan")
	}

	return comment, nil
}

func (r *commentsRepo) Update(ctx context.Context, updateCommentRequest *models.UpdateCommentRequest) (*models.Comment, error) {
	comment := &models.Comment{}

	if err := r.db.QueryRowxContext(
		ctx,
		updateComment,
		updateCommentRequest.Message,
		updateCommentRequest.Id,
	).StructScan(comment); err != nil {
		return nil, errors.Wrap(err, "commentsRepo.Update.StructScan")
	}

	return comment, nil
}

func (r *commentsRepo) IncreaseLikeCount(ctx context.Context, increaseLikeRequest *models.IncreaseLikeRequest) error {
	result, err := r.db.ExecContext(ctx, increaseLikeCount, increaseLikeRequest.Id)
	if err != nil {
		return errors.Wrap(err, "commentsRepo.IncreaseLikeCount.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return errors.Wrap(err, "commentsRepo.IncreaseLikeCount.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "commentsRepo.IncreaseLikeCount.rowsAffected")
	}

	return nil
}
