package repository

import (
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Comments Repository
type usersRepo struct {
	db *sqlx.DB
}

// Comments Repository constructor
func NewUsersRepository(db *sqlx.DB) users.Repository {
	return &usersRepo{db: db}
}

func (u usersRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	if err := u.db.GetContext(ctx, user, getUserByID, id); err != nil {
		return nil, errors.Wrap(err, "usersRepo.GetByID.GetContext")
	}
	return user, nil
}

func (u usersRepo) Login(ctx context.Context, request models.LoginRequest) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
