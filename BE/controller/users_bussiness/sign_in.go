package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model"
)

type SignInService interface {
	SignIn(ctx context.Context, data *model.Users) error
}

type SignInController struct {
	s SignInService
}

func NewSignInController(s SignInService) *SignInController {
	return &SignInController{s: s}
}
func (si *SignInController) NewSignIn(ctx context.Context, data *model.Users) error {
	if data.Email == "" {
		return errors.New("SignIn error")
	}
	if err := si.s.SignIn(ctx, data); err != nil {
		return errors.New("SignIn failed")
	}
	return nil
}
