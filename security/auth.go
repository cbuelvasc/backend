package security

import (
	"context"

	"github.com/cbuelvasc/backend/model"
	"github.com/cbuelvasc/backend/repository"
	"github.com/cbuelvasc/backend/util"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

func (authValidator *AuthValidator) ValidateCredentials(ctx context.Context, username, password string) (*model.User, bool) {
	user, err := authValidator.userRepository.FindByEmail(ctx, username)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}
