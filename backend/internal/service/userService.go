package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DeS313/cloud-disk/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	salt       = "sadfsdafa32434212"
	signingKey = "jkjdskfjaks3839#83kas"
	tokenTTL   = 1 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	user_id string `jsong:"user_id"`
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *Service) GenerateToken(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not  of type *tokenClaims")
	}

	return claims.user_id, nil
}

func (s *Service) FindOne(ctx context.Context, key, value string) (models.User, error) {

	return s.storage.FindOne(ctx, value)
}

func (s *Service) FindOneByEmail(ctx context.Context, user *models.User) (models.User, error) {
	u, err := s.storage.FindOneByEmail(ctx, user.Email)
	if err != nil {
		log.Println(err)
		log.Println(mongo.ErrNoDocuments)
		return models.User{}, err
	}

	if u.Password != generatePasswordHash(user.Password) {
		return models.User{}, fmt.Errorf("неверный логин или пароль")
	}

	return u, nil
}

func (s *Service) Create(ctx context.Context, user *models.User) (string, error) {
	//  TODO придумать как проверять если ли пользователь с таким email
	fmt.Println(len(user.Password))
	if len(user.Password) < 3 && len(user.Password) > 12 {
		return "некорректный пароль", fmt.Errorf("некорректный пароль")
	}

	if !strings.Contains(user.Email, "@") {
		return "некорртный email", fmt.Errorf("некорртный email")
	}

	user.Password = generatePasswordHash(user.Password)

	id, err := s.storage.Create(ctx, *user)

	return id, err
}
