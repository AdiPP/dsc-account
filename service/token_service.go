package service

import (
	"errors"
	"reflect"
	"time"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/golang-jwt/jwt/v4"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type TokenService struct{}

func NewTokenService() TokenService {
	return TokenService{}
}

var (
	jwtKey = []byte("expecto_patronum")
	JwtKey = jwtKey
)

func (ts *TokenService) IssueToken(u entity.User, crdn Credential) (Token, error) {
	if reflect.DeepEqual(u, entity.User{}) || u.Password != crdn.Password {
		return Token{}, errors.New("credential is invalid")
	}

	ExpiresAt := jwt.NewNumericDate(time.Now().Add(time.Minute * 5))

	clm := Claim{
		Username: crdn.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: ExpiresAt,
		},
	}

	jwtTkn := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)

	jwtTknStrng, err := jwtTkn.SignedString(jwtKey)

	if err != nil {
		return Token{}, err
	}

	tkn := Token{
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(ExpiresAt.Time).Seconds()),
		AccessToken: jwtTknStrng,
	}

	return tkn, nil
}

func (ts *TokenService) RefreshToken(jwtTknStr string) (Token, error) {
	clm := Claim{}

	_, err := ts.ValidateToken(jwtTknStr)

	if err != nil {
		return Token{}, err
	}

	ExpiresAt := jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
	clm.ExpiresAt = ExpiresAt

	jwtTkn := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)
	jwtTknStrng, _ := jwtTkn.SignedString(jwtKey)

	if err != nil {
		return Token{}, errors.New("internal server error")
	}

	tkn := Token{
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(ExpiresAt.Time).Seconds()),
		AccessToken: jwtTknStrng,
	}

	return tkn, nil
}

func (ts *TokenService) ValidateToken(jwtTknStr string) (*jwt.Token, error) {
	clm := Claim{}

	jwtTkn, err := jwt.ParseWithClaims(jwtTknStr, &clm, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return jwtTkn, jwt.ErrSignatureInvalid
		}

		return jwtTkn, errors.New("bad request")
	}

	if !jwtTkn.Valid {
		return jwtTkn, errors.New("unauthorized")
	}

	if time.Until(clm.ExpiresAt.Time).Seconds() < 0 {
		return jwtTkn, errors.New("bad request")
	}

	return jwtTkn, nil
}
