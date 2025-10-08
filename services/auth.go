package services

import (
	"fmt"
	"strings"

	"github.com/Felipek06/TradeBackend_dev.git/repositories"
	"github.com/Felipek06/TradeBackend_dev.git/utils"
	"github.com/golang-jwt/jwt"
)

type NewAuthService struct {
	Database *repositories.NewDatabase
}

func (a *NewAuthService) Login(email, password string) (string, error) {
	findUserInfo, err := a.Database.FindUserByField("email", email)
	if err != nil && !strings.EqualFold(err.Error(), "record not found") {
		return "", fmt.Errorf("error while finding user on db: %s", err)
	}
	if findUserInfo == nil {
		return "", fmt.Errorf("user no exists")
	}

	verifyPassword, err := utils.HashAndCompareTwoStrings(password, findUserInfo.Password)
	if !verifyPassword {
		return "", fmt.Errorf("invalid password")
	}
	if err != nil {
		return "", fmt.Errorf("error while comparing passwords: %s", err)
	}

	jwtToken, err := utils.CreateJWTToken(jwt.MapClaims{"ID": findUserInfo.ID})
	if err != nil {
		return "", fmt.Errorf("error while create JWT token: %s", err)
	}

	return jwtToken, nil
}
