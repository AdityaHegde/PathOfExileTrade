package authcore

import (
	"crypto/rand"
  "crypto/rsa"
	"errors"
	"fmt"
	accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

const UserNameKey = "UserName"

type JwtAuth struct {
	key *rsa.PrivateKey
}

func (jwtAuth *JwtAuth) Init() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	jwtAuth.key = key
	return nil
}

func (jwtAuth *JwtAuth) Generate(user *accountmodel.User) (string, error) {
	token := jwt.New()

	setErr := token.Set(UserNameKey, user.Name)
	if setErr != nil {
		return "", setErr
	}

	payload, signErr := jwt.Sign(token, jwa.RS256, jwtAuth.key)
	if signErr != nil {
		return "", signErr
	}

	return string(payload), nil
}

func (jwtAuth *JwtAuth) Validate(payload string) (string, error) {
	token, parseErr := jwt.Parse([]byte(payload), jwt.WithValidate(true),
		jwt.WithVerify(jwa.RS256, &jwtAuth.key.PublicKey))
	if parseErr != nil {
		return "", parseErr
	}

	userName, exists := token.Get(UserNameKey)
	if exists {
		return fmt.Sprintf("%s", userName), nil
	}

	return "", errors.New("UserName not found in token")
}
