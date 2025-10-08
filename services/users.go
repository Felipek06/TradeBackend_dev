package services

import (
	"fmt"
	"strings"

	"github.com/Felipek06/TradeBackend_dev.git/repositories"
	"github.com/Felipek06/TradeBackend_dev.git/utils"
)

type NewUserService struct {
	Database *repositories.NewDatabase
}

func (s *NewUserService) CreateNewUser(email, password string) error {
	exists, err := s.verifyIfUserAlreadyExists(email)
	if err != nil {
		return fmt.Errorf("error on verifying user: %s", err)
	}
	if exists {
		return fmt.Errorf("user already exists")
	}

	hashedPassword, err := utils.HashAString(password)
	if err != nil {
		return fmt.Errorf("error on hash password: %s", err)
	}

	parsedUser := repositories.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.Database.Insert(&parsedUser)
	if err != nil {
		return fmt.Errorf("error on insert user: %s", err)
	}

	return nil
}

func (s *NewUserService) verifyIfUserAlreadyExists(email string) (bool, error) {
	result, err := s.Database.FindUserByField("email", email)
	if err != nil && !strings.EqualFold(err.Error(), "record not found") {
		return false, err
	}

	if result != nil {
		return true, nil
	}

	return false, nil
}
