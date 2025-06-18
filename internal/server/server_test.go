package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"void.dev/identity/internal/server"
)

//-------------------------------------------------------------------------------------------------

func TestNewServer(t *testing.T) {
	cfg := server.DefaultConfig()
	assert.Equal(t, "dev", cfg.Env)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 3000, cfg.Port)
	assert.Equal(t, "Hello World", cfg.Message)

	server, err := server.NewServer(cfg)
	require.Nil(t, err)
	require.NotNil(t, server)

	assert.Equal(t, cfg.Env, server.Env, "verify server config matches requested config")
	assert.Equal(t, cfg.Host, server.Host, "verify server config matches requested config")
	assert.Equal(t, cfg.Port, server.Port, "verify server config matches requested config")
	assert.Equal(t, cfg.Message, server.Message, "verify server config matches requested config")
}

//-------------------------------------------------------------------------------------------------

func TestPing(t *testing.T) {
	// rr := serve(t, get("/ping"), server.DefaultConfig())
	// response := assertex.JSONResponseOK(t, rr)
	// assert.Equal(t, map[string]interface{}{"ping": "pong"}, response)
}

//-------------------------------------------------------------------------------------------------

func TestHello(t *testing.T) {
	// rr := serve(t, get("/"), server.DefaultConfig())
	// response := assertex.JSONResponseOK(t, rr)
	// assert.Equal(t, map[string]interface{}{
	// 	"message": "Hello World",
	// 	"version": map[string]interface{}{
	// 		"env":      "dev",
	// 		"domain":   "platform-examples",
	// 		"service":  "hello-world",
	// 		"build":    "none",
	// 		"revision": "HEAD",
	// 		"version":  "latest",
	// 		"swagger":  "/swagger.json",
	// 	},
	// }, response)
}

//-------------------------------------------------------------------------------------------------

func TestHelloWithCustomConfig(t *testing.T) {
	// cfg := server.DefaultConfig()
	// cfg.Message = "this is a custom message"
	// rr := serve(t, get("/custom/path"), cfg)
	// response := assertex.JSONResponseOK(t, rr)
	// assert.Equal(t, map[string]interface{}{
	// 	"message": "this is a custom message",
	// 	"version": map[string]interface{}{
	// 		"env":      "dev",
	// 		"domain":   "platform-examples",
	// 		"service":  "hello-world",
	// 		"build":    "none",
	// 		"revision": "HEAD",
	// 		"version":  "latest",
	// 		"swagger":  "/custom/path/swagger.json",
	// 	},
	// }, response)
}

//-------------------------------------------------------------------------------------------------
// PRIVATE TEST HELPER METHODS
//-------------------------------------------------------------------------------------------------

// func get(path string) *http.Request {
// 	return httptest.NewRequest(echo.GET, path, nil)
// }
//
// func makeServer(t *testing.T, cfg *server.Config) (*server.Server, *httptest.ResponseRecorder) {
// 	rr := httptest.NewRecorder()
// 	server, err := server.NewServer(cfg)
// 	require.Nil(t, err)
// 	return server, rr
// }
//
// func serve(t *testing.T, request *http.Request, cfg *server.Config) *httptest.ResponseRecorder {
// 	server, rr := makeServer(t, cfg)
// 	server.ServeHTTP(rr, request)
// 	return rr
// }

//-------------------------------------------------------------------------------------------------
