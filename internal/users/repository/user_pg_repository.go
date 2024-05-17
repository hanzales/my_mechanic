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

func (u *usersRepo) GetUserById(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}

	if err := u.db.QueryRowxContext(
		ctx,
		userById,
		id,
	).StructScan(user); err != nil {
		return nil, errors.Wrap(err, "usersRepo.GetUserById.GetContext")
	}

	return user, nil
}

func (u *usersRepo) Register(ctx context.Context, request *models.RegisterRequest) (*models.User, error) {
	user := &models.User{}

	if err := u.db.QueryRowxContext(
		ctx,
		createUser,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.About,
		&user.Avatar,
		&user.PhoneNumber,
		&user.Address,
		&user.City,
		&user.Gender,
		&user.Postcode,
		&user.Birthday,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.Register.StructScan")
	}

	return user, nil
}
