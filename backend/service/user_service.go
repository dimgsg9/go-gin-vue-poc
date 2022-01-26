package service

import (
	"context"
	"log"

	"github.com/dimgsg9/booker_proto/backend/model"
	"github.com/dimgsg9/booker_proto/backend/model/apperrors"
	"github.com/google/uuid"
)

// userService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type userService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected into this
// this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}

	// Maybe not the best idea to mutate the password here
	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	// a good place to publish user create event here

	return nil
}

func (s *userService) Signin(ctx context.Context, u *model.User) error {
	uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)

	// Will return NotAuthorized to client to omit details of why
	if err != nil {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	// verify password - we previously created this method
	match, err := comparePasswords(uFetched.Password, u.Password)

	if err != nil {
		return apperrors.NewInternal()
	}

	if !match {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	*u = *uFetched // Consider refactor as per https://dev.to/jacobsngoodwin/16-create-gin-middleware-to-extract-authorized-user-1jom

	return nil
}

func (s *userService) OauthLogin(ctx context.Context, u *model.User) error {
	uFetched, _ := s.UserRepository.FindByEmail(ctx, u.Email) //TODO: refactor into FindOrCreate()

	// uFetched is never nil, thus checking Email attribute
	if uFetched.Email == "" {
		u.UID = uuid.New()
		// create a user on the fly
		if err := s.UserRepository.PwdlessCreate(ctx, u); err != nil {
			return err
		}
	}

	u.UID = uFetched.UID

	return nil
}
