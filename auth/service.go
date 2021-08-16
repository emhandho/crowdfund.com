package auth

import "github.com/dgrijalva/jwt-go"

var SECRET_KEY = []byte("MYCROWDFUNDAPPS_s3cr3tK3y")

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService{
	return &jwtService{}
}

func (j *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}