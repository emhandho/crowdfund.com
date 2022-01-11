package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// var SECRET_KEY = []byte(os.Getenv("APP_SECRET_KEY"))

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	SECRET_KEY []byte
}

func NewJwtService(SK []byte) *jwtService {
	return &jwtService{SECRET_KEY: SK}
}

func (j *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(j.SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(j.SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
