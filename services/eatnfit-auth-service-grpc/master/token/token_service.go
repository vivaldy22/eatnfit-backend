package token

import (
	"context"
	authproto "github.com/vivaldy22/eatnfit-auth-service-grpc/proto"
	"github.com/vivaldy22/eatnfit-auth-service-grpc/tools/jwttoken"
)

const (
	customKey = "keyEatNFit"
)

type Service struct {}

func (s Service) GenerateToken(ctx context.Context, credentials *authproto.TokenCredentials) (*authproto.Token, error) {
	token, err := jwttoken.JwtEncoder(credentials.UserEmail, customKey, credentials.HmacSecret)
	if err != nil {
		return nil, err
	}
	return &authproto.Token{Token: token}, err
}

func NewService() authproto.JWTTokenServer {
	return &Service{}
}