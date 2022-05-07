package service

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/AdiPP/dsc-account/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
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
)

func (ts *TokenService) AuthUser(jwtTknStr string) (entity.User, error) {
	jwtTkn, err := validateToken(jwtTknStr)

	if err != nil {
		return entity.User{}, err
	}

	clm, err := covertJwtTokenToClaim(jwtTkn)

	if err != nil {
		return entity.User{}, err
	}

	u, err := userRepository.FindByUsernameOrFail(clm.Username)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (ts *TokenService) IssueToken(crdn Credential) (Token, error) {
	u, err := userRepository.FindByUsernameOrFail(crdn.Username)

	if err != nil {
		return Token{}, err
	}

	if reflect.DeepEqual(u, entity.User{}) {
		return Token{}, errors.New("user does not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(crdn.Password))

	if err != nil {
		return Token{}, errors.New("invalid credential")
	}

	ExpiresAt := jwt.NewNumericDate(time.Now().Add(time.Minute * 5))

	clm := Claim{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
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
	jwtTkn, err := validateToken(jwtTknStr)

	if err != nil {
		return Token{}, err
	}

	clm, err := covertJwtTokenToClaim(jwtTkn)

	if err != nil {
		return Token{}, err
	}

	clm.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))

	jwtTkn = jwt.NewWithClaims(jwt.SigningMethodHS256, clm)

	jwtTknStrng, _ := jwtTkn.SignedString(jwtKey)

	if err != nil {
		return Token{}, errors.New("internal server error")
	}

	tkn := Token{
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(clm.ExpiresAt.Time).Seconds()),
		AccessToken: jwtTknStrng,
	}

	return tkn, nil
}

func (ts *TokenService) ValidateToken(jwtTknStr string) (*jwt.Token, error) {
	return validateToken(jwtTknStr)
}

func validateToken(jwtTknStr string) (*jwt.Token, error) {
	jwtTkn, err := jwt.Parse(jwtTknStr, func(t *jwt.Token) (interface{}, error) {
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

	return jwtTkn, nil
}

func covertJwtTokenToClaim(jwtTkn *jwt.Token) (Claim, error) {
	clmsMap, ok := jwtTkn.Claims.(jwt.MapClaims)

	if !ok {
		return Claim{}, errors.New("")
	}

	var expUnix int64

	switch exp := clmsMap["exp"].(type) {
	case float64:
		expUnix = int64(exp)
	case json.Number:
		v, _ := exp.Int64()
		expUnix = int64(v)
	}

	ExpiresAt := jwt.NewNumericDate(time.Unix(expUnix, 0))

	clm := Claim{
		Username: clmsMap["username"].(string),
		Email:    clmsMap["email"].(string),
		Name:     clmsMap["name"].(string),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: ExpiresAt,
		},
	}

	return clm, nil
}
