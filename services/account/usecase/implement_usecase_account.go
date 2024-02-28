package usecase

import (
	"context"
	"errors"
	"tinder-cloning/models"
	"tinder-cloning/pkg/util"
	"tinder-cloning/services/account/repository"
	"tinder-cloning/services/account/schema"
	membershipUseCase "tinder-cloning/services/membership/usecase"
)

type implementAccountUseCase struct {
	repo              repository.Repository
	membershipUseCase membershipUseCase.MembershipUseCase
	time              util.Time
}

func NewAccountUseCase(repo repository.Repository, membershipUseCase membershipUseCase.MembershipUseCase) AccountUseCase {
	return &implementAccountUseCase{repo: repo, membershipUseCase: membershipUseCase, time: util.ProvideNewTimesCustom()}
}

func (i *implementAccountUseCase) validateAccount(payload *schema.RequestLogin) error {
	if payload.Email == "" {
		return errors.New("email is required")
	}
	if !util.IsValidEmail(payload.Email) {
		return errors.New("email is not valid format")
	}
	if payload.Password == "" {
		return errors.New("password is required")
	}
	if !util.IsValidMinLength(payload.Password, DefaultMaxPasswdLength) {
		return errors.New("password minimum length is " + string(rune(DefaultMaxPasswdLength)) + " characters")
	}
	return nil
}

func (i *implementAccountUseCase) validateSignUp(payload *schema.RequestRegister) error {
	if err := i.validateAccount(&schema.RequestLogin{
		Email:    payload.Email,
		Password: payload.Password,
	}); err != nil {
		return err
	}

	if payload.Gender == "" {
		return errors.New("gender is required")
	}

	if payload.Username == "" {
		return errors.New("username is required")
	}
	if util.HasSpace(payload.Username) {
		return errors.New("username cannot contain space")
	}
	if util.IsValidUsername(payload.Username) && util.IsValidMinLength(payload.Username, DefaultMinUsernameLength) {
		return errors.New("username is not valid format (only alphanumeric, underscore, minimum 3 characters)")
	}

	if !util.IsValidLocation(payload.Location) {
		return errors.New("location is not valid")
	}

	// check if date of birth is valid
	if payload.DateOfBirth != "" {
		if !util.IsValidDateOfBirth(payload.DateOfBirth) {
			return errors.New("date of birth is not valid")
		}
	}

	return nil
}

func (i *implementAccountUseCase) SignUp(ctx context.Context, payload *schema.RequestRegister) error {
	if err := i.validateSignUp(payload); err != nil {
		return err
	}

	// Check if username already exists
	ac, err := i.repo.FindOne(ctx, schema.AccountFilter{
		Email:    payload.Email,
		Username: payload.Username,
		ID:       nil,
	})
	if err != nil && errors.Is(err, errors.New("sql: no rows in result set")) {
		return err
	}
	if ac != nil {
		return errors.New("username or email already exists")
	}
	// has password
	payload.Password, err = util.Hash(payload.Password)
	if err != nil {
		return err
	}

	dateOfBirth, err := i.time.GenerateDateOfBirthFromString(payload.DateOfBirth)
	if err != nil {
		return err
	}
	// Create account
	account := &models.Account{
		Email:        payload.Email,
		Username:     payload.Username,
		PasswordHash: payload.Password,
		Gender:       payload.Gender,
		CreatedAt:    i.time.Now(nil),
		IsVerified:   false,
		Location:     &payload.Location,
		Bio:          &payload.Bio,
		Avatar:       &payload.Avatar,
		DateOfBirth:  &dateOfBirth,
	}

	accountID, tx, errCreate := i.repo.CreateOne(ctx, nil, account)
	if errCreate != nil {
		return errCreate
	}

	//// Create membership
	membership := &models.Membership{
		AccountID:      accountID,
		MembershipType: membershipUseCase.LevelFree,
		StartDate:      nil,
		EndDate:        nil,
		PaymentMethod:  "",
	}

	txFinal, errCreateMembership := i.membershipUseCase.CreateOne(ctx, tx, membership)
	if errCreateMembership != nil {
		return errCreateMembership
	}

	err = txFinal.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (i *implementAccountUseCase) SingIn(ctx context.Context, payload *schema.RequestLogin) (string, error) {
	if err := i.validateAccount(payload); err != nil {
		return "", err
	}
	account, err := i.repo.FindOne(ctx, schema.AccountFilter{
		Email: payload.Email,
		ID:    nil,
	})
	if err != nil {
		return "", err
	}
	// compare password
	err = util.CompareHashAndPassword(account.PasswordHash, payload.Password)
	if err != nil {
		return "", err
	}
	// generate token
	token, err := util.GenerateTokenJWT(account.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (i *implementAccountUseCase) GetOne(ctx context.Context, filter schema.AccountFilter) (*models.AccountAsProfile, error) {
	account, err := i.repo.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}
	return &models.AccountAsProfile{
		ID:          account.ID,
		Email:       account.Email,
		Username:    account.Username,
		Gender:      account.Gender,
		Bio:         account.Bio,
		Avatar:      account.Avatar,
		DateOfBirth: account.DateOfBirth,
		Location:    account.Location,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}, nil
}
