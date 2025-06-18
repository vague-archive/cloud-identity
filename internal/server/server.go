package server

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/text/language"
)

//-------------------------------------------------------------------------------------------------

type CustomContext struct {
	echo.Context
	username *string
}

func customContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c, nil}
		return next(cc)
	}
}

func (c *CustomContext) isLoggedIn() bool {
	return c.username != nil
}

//-------------------------------------------------------------------------------------------------

type Server struct {
	*Config
	router *echo.Echo
}

//-------------------------------------------------------------------------------------------------

func NewServer(cfg *Config) (*Server, error) {

	server := &Server{
		Config: cfg,
		router: echo.New(),
	}

	server.router.Use(middleware.Logger())
	server.router.Use(customContext)
	server.router.HideBanner = true
	server.router.HidePort = true

	server.router.GET("/", server.Hello)
	server.router.GET("/goodbye", server.Goodbye)
	server.router.GET("/login", server.Login)
	server.router.GET("/ping", server.Ping)

	key := sha256.Sum256([]byte("test")) // TODO: need real secret key

	storage := newStorage()
	provider, err := NewProvider(cfg.Issuer, key, storage)
	if err != nil {
		log.Fatal(err)
	}

	handler := http.Handler(provider)

	oidc := server.router.Group("/oidc/*")
	oidc.Use(middleware.Rewrite(map[string]string{
		"/oidc/*": "/$1",
	}))
	oidc.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.WrapHandler(handler)
	})

	return server, nil

}

//-------------------------------------------------------------------------------------------------
// START & STOP
//-------------------------------------------------------------------------------------------------

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	log.Printf("Environment: %s", s.Env)
	log.Printf("Version: %s", s.Version)
	log.Printf("Listening on %s", address)
	log.Printf("Issuer %s", s.Config.Issuer)
	return s.router.Start(address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}

//-------------------------------------------------------------------------------------------------
// ENDPOINTS
//-------------------------------------------------------------------------------------------------

func (s *Server) Hello(c echo.Context) error {
	context := c.(*CustomContext)
	if context.isLoggedIn() {
		return context.HTML(http.StatusOK, "<h1>TODO: Welcome Screen</h1>")
	} else {
		return context.Redirect(http.StatusFound, "/login")
	}
}

func (s *Server) Goodbye(c echo.Context) error {
	context := c.(*CustomContext)
	if context.isLoggedIn() {
		return context.Redirect(http.StatusFound, "/hello")
	} else {
		return context.HTML(http.StatusOK, "<h1>Signed out successfully</h1>")
	}
}

func (s *Server) Login(c echo.Context) error {
	context := c.(*CustomContext)
	if context.isLoggedIn() {
		return context.Redirect(http.StatusFound, "/hello")
	} else {
		return context.HTML(http.StatusOK, "<h1>TODO: Login Screen</h1>")
	}
}

func (s *Server) Ping(c echo.Context) error {
	context := c.(*CustomContext)
	return context.JSON(http.StatusOK, map[string]interface{}{"ping": "pong", "username": context.username})
}

//-------------------------------------------------------------------------------------------------
// PROVIDER
//-------------------------------------------------------------------------------------------------

func NewProvider(issuer string, key [32]byte, storage op.Storage) (op.OpenIDProvider, error) {
	config := &op.Config{
		CryptoKey: key,

		// will be used if the end_session endpoint is called without a post_logout_redirect_uri
		DefaultLogoutRedirectURI: "/goodbye",

		// enables code_challenge_method S256 for PKCE (and therefore PKCE in general)
		CodeMethodS256: true,

		// enables additional client_id/client_secret authentication by form post (not only HTTP Basic Auth)
		AuthMethodPost: true,

		// enables additional authentication by using private_key_jwt
		AuthMethodPrivateKeyJWT: true,

		// enables refresh_token grant use
		GrantTypeRefreshToken: true,

		// enables use of the `request` Object parameter
		RequestObjectSupported: true,

		// this example has only static texts (in English), so we'll set the here accordingly
		SupportedUILocales: []language.Tag{language.English},

		DeviceAuthorization: op.DeviceAuthorizationConfig{
			Lifetime:     5 * time.Minute,
			PollInterval: 5 * time.Second,
			UserFormPath: "/device",
			UserCode:     op.UserCodeBase20,
		},
	}

	// op.newProvider creates a provider. The provider provides (with HttpHandler())
	// a http.Router that handles a suite of endpoints (some paths can be overridden):
	//
	//	/healthz
	//	/ready
	//	/.well-known/openid-configuration
	//	/oauth/token
	//	/oauth/introspect
	//	/callback
	//	/authorize
	//	/userinfo
	//	/revoke
	//	/end_session
	//	/keys
	//	/device_authorization
	//
	handler, err := op.NewProvider(config, storage, op.StaticIssuer(issuer),
		[]op.Option{
			//we must explicitly allow the use of the http issuer
			op.WithAllowInsecure(),
		}...,
	)
	if err != nil {
		return nil, err
	}
	return handler, nil
}
