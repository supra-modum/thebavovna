package server

import (
	"bavovna/internal/conf"
	"bavovna/internal/store"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

var (
	signer   jwt.Signer
	verifier jwt.Verifier
)

func jwtSetup(conf conf.Config) {
	var err error

	key := []byte(conf.JwtSecret)

	signer, err = jwt.NewSignerHS(jwt.HS256, key)
	checkErr(err)

	verifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	checkErr(err)

}

func generateJWT(user *store.User) string {
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}

	builder := jwt.NewBuilder(signer)

	token, err := builder.Build(claims)

	checkErr(err)

	return token.String()
}

func verifyJWT(tokenStr string) (int, error) {
	rawToken := []byte(tokenStr)
	token, err := jwt.Parse(rawToken, verifier)
	if err != nil {
		log.Error().Err(err).Str("tokenStr", tokenStr).Msg("Error parsing JWT")
		return 0, err
	}

	if err := verifier.Verify(token); err != nil {
		log.Error().Err(err).Msg("Error verifying token")
		return 0, err
	}

	var claims jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claims); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JWT claims")
		return 0, err
	}

	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("token expired")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Error().Err(err).Str("claims.ID", claims.ID).Msg("Error converting claims ID to number")
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
