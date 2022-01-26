package repository

import (
	"context"
	"log"

	"github.com/dimgsg9/booker_proto/backend/model"
	"github.com/dimgsg9/booker_proto/backend/model/apperrors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// pgUserRepository is data/repository implementation
// of service layer UserRepository
type pgUserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository is a factory for initializing User Repositories
func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &pgUserRepository{
		DB: db,
	}
}

// Create reaches out to database SQLX api
func (r *pgUserRepository) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *"

	if err := r.DB.GetContext(ctx, u, query, u.Email, u.Password); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
			return apperrors.NewConflict("email", u.Email)
		}

		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return apperrors.NewInternal()
	}
	return nil
}

// Create reaches out to database SQLX api
func (r *pgUserRepository) PwdlessCreate(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (email, picture, name, password) VALUES ($1, $2, $3, $4) RETURNING *"

	if err := r.DB.GetContext(ctx, u, query, u.Email, u.Picture, u.Name, ""); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
			return apperrors.NewConflict("email", u.Email)
		}

		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return apperrors.NewInternal()
	}
	return nil
}

// FindByID fetches user by id
func (r *pgUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid=$1"

	// we need to actually check errors as it could be something other than not found
	if err := r.DB.GetContext(ctx, user, query, uid); err != nil {
		return user, apperrors.NewNotFound("uid", uid.String())
	}

	return user, nil
}

func (r *pgUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE email=$1"

	if err := r.DB.GetContext(ctx, user, query, email); err != nil {
		log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
		return user, apperrors.NewNotFound("email", email)
	}

	return user, nil
}
