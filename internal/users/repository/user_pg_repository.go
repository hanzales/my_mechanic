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
func UsersRepository(db *sqlx.DB) users.Repository {
	return &usersRepo{db: db}
}

func (u *usersRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	if err := u.db.QueryRowxContext(
		ctx,
		getUserByEmail,
		email,
	).StructScan(user); err != nil {
		return nil, errors.Wrap(err, "usersRepo.GetUserByEmail.GetContext")
	}

	return user, nil
}
