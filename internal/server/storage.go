package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

type Storage struct {
	signatureAlgorithm jose.SignatureAlgorithm
}

func newStorage() *Storage {
	return &Storage{
		signatureAlgorithm: jose.RS256,
	}
}

func (s *Storage) AuthRequestByCode(ctx context.Context, code string) (op.AuthRequest, error) {
	todo("AuthRequestByCode")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) AuthRequestByID(ctx context.Context, id string) (op.AuthRequest, error) {
	todo("AuthRequestByID")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) AuthorizeClientIDSecret(ctx context.Context, clientID, clientSecret string) error {
	todo("AuthorizeClientIDSecret")
	return fmt.Errorf("TODO")
}

func (s *Storage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	todo("CreateAccessAndRefreshTokens")
	return "", "", time.Now(), fmt.Errorf("TODO")
}

func (s *Storage) CreateAccessToken(ctx context.Context, request op.TokenRequest) (string, time.Time, error) {
	todo("CreateAccessToken")
	return "", time.Now(), fmt.Errorf("TODO")
}

func (s *Storage) CreateAuthRequest(ctx context.Context, authReq *oidc.AuthRequest, userID string) (op.AuthRequest, error) {
	todo("CreateAuthRequest")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) DeleteAuthRequest(ctx context.Context, id string) error {
	todo("DeleteAuthRequest")
	return fmt.Errorf("TODO")
}

func (s *Storage) GetClientByClientID(ctx context.Context, clientID string) (op.Client, error) {
	return NewClient(clientID), nil
}

func (s *Storage) GetKeyByIDAndClientID(ctx context.Context, keyID, clientID string) (*jose.JSONWebKey, error) {
	todo("GetKeyByIDAndClientID")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) GetPrivateClaimsFromScopes(ctx context.Context, userID, clientID string, scopes []string) (claims map[string]any, err error) {
	todo("GetPrivateClaimsFromScopes")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) GetRefreshTokenInfo(ctx context.Context, clientID string, token string) (userID string, tokenID string, err error) {
	todo("GetRefreshTokenInfo")
	return "", "", fmt.Errorf("TODO")
}

func (s *Storage) Health(ctx context.Context) error {
	todo("Health")
	return fmt.Errorf("TODO")
}

func (s *Storage) KeySet(ctx context.Context) ([]op.Key, error) {
	todo("KeySet")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) RevokeToken(ctx context.Context, tokenIDOrToken string, userID string, clientID string) *oidc.Error {
	todo("RevokeToken")
	return nil
}

func (s *Storage) SaveAuthCode(ctx context.Context, id string, code string) error {
	todo("SaveAuthCode")
	return fmt.Errorf("TODO")
}

func (s *Storage) SetIntrospectionFromToken(ctx context.Context, introspection *oidc.IntrospectionResponse, tokenID, subject, clientID string) error {
	todo("SetIntrospectionFromToken")
	return fmt.Errorf("TODO")
}

func (s *Storage) SetUserinfoFromScopes(ctx context.Context, userinfo *oidc.UserInfo, userID, clientID string, scopes []string) error {
	todo("SetUserinfoFromScopes")
	return fmt.Errorf("TODO")
}

func (s *Storage) SetUserinfoFromToken(ctx context.Context, userinfo *oidc.UserInfo, tokenID, subject, origin string) error {
	todo("SetUserinfoFromToken")
	return fmt.Errorf("TODO")
}

func (s *Storage) SignatureAlgorithms(context.Context) ([]jose.SignatureAlgorithm, error) {
	return []jose.SignatureAlgorithm{s.signatureAlgorithm}, nil
}

func (s *Storage) SigningKey(ctx context.Context) (op.SigningKey, error) {
	todo("SigningKey")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) TerminateSession(ctx context.Context, userID string, clientID string) error {
	todo("TerminateSession")
	return fmt.Errorf("TODO")
}

func (s *Storage) TokenRequestByRefreshToken(ctx context.Context, refreshToken string) (op.RefreshTokenRequest, error) {
	todo("TokenRequestByRefreshToken")
	return nil, fmt.Errorf("TODO")
}

func (s *Storage) ValidateJWTProfileScopes(ctx context.Context, userID string, scopes []string) ([]string, error) {
	todo("ValidateJWTProfileScopes")
	return nil, fmt.Errorf("TODO")
}

func todo(name string) {
	log.Printf("TODO: %s", name)
}
