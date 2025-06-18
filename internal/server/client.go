package server

import (
	"time"

	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

type Client struct {
	id              string
	applicationType op.ApplicationType
	accessTokenType op.AccessTokenType
	authMethod      oidc.AuthMethod
	grantTypes      []oidc.GrantType
	responseTypes   []oidc.ResponseType
	idTokenLifetime time.Duration
}

func NewClient(id string) *Client {
	return &Client{
		id:              id,
		applicationType: op.ApplicationTypeWeb,
		accessTokenType: op.AccessTokenTypeBearer,
		authMethod:      oidc.AuthMethodBasic,
		grantTypes:      []oidc.GrantType{oidc.GrantTypeCode},
		responseTypes:   []oidc.ResponseType{oidc.ResponseTypeCode},
		idTokenLifetime: 7 * 24 * time.Hour, // TODO: shorten this once we have refresh token support
	}
}

func (c *Client) GetID() string {
	return c.id
}

func (c *Client) ApplicationType() op.ApplicationType {
	return c.applicationType
}

func (c *Client) AccessTokenType() op.AccessTokenType {
	todo("Client.AccessTokenType")
	return c.accessTokenType
}

func (c *Client) AuthMethod() oidc.AuthMethod {
	todo("Client.AuthMethod")
	return c.authMethod
}

func (c *Client) GrantTypes() []oidc.GrantType {
	todo("Client.GrantTypes")
	return c.grantTypes
}

func (c *Client) LoginURL(id string) string {
	todo("Client.LoginURL")
	return "/login?authRequestID=" + id
}

func (c *Client) RedirectURIs() []string {
	todo("Client.RedirectURIs")
	return []string{}
}

func (c *Client) PostLogoutRedirectURIs() []string {
	todo("Client.PostLogoutRedirectURIs")
	return []string{}
}

func (c *Client) ResponseTypes() []oidc.ResponseType {
	todo("Client.ResponseTypes")
	return c.responseTypes
}

func (c *Client) IDTokenLifetime() time.Duration {
	todo("Client.IDTokenLifetime")
	return c.idTokenLifetime
}

func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	todo("Client.IDTokenUserinfoClaimsAssertion")
	return false
}

func (c *Client) IsScopeAllowed(scope string) bool {
	todo("Client.IsScopeAllowed")
	return false
}

func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}
func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

func (c *Client) ClockSkew() time.Duration {
	todo("Client.ClockSkew")
	return 0
}

func (c *Client) DevMode() bool {
	todo("Client.DevMode")
	return false
}
