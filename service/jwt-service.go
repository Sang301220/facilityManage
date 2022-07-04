package service

import (
	"fmt"
	"os"
	"strings"
	_ "strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string, userName string, userEmail string, userPhone string, UserRoleID uint64) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"name"`
	UserEmail  string `json:"email"`
	UserPhone  string `json:"phone"`
	UserRoleID uint64 `json:"role_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string, UserName string, UserEmail string, UserPhone string, UserRoleID uint64) string {
	claims := &jwtCustomClaim{
		UserID,
		UserName,
		UserEmail,
		UserPhone,
		UserRoleID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtString := strings.Split(token, "Bearer ")[1]
	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {

		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
