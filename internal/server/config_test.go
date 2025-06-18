package server_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"void.dev/identity/internal/server"
)

//-------------------------------------------------------------------------------------------------

func TestNewConfig(t *testing.T) {
	cfg := server.NewConfig("v1.2.3", "env", "0.0.0.0", 1234, "https://id.void.dev/", "foobar")
	assert.Equal(t, "v1.2.3", cfg.Version)
	assert.Equal(t, "env", cfg.Env)
	assert.Equal(t, "0.0.0.0", cfg.Host)
	assert.Equal(t, 1234, cfg.Port)
	assert.Equal(t, "https://id.void.dev/", cfg.Issuer)
	assert.Equal(t, "foobar", cfg.Message)
}

//-------------------------------------------------------------------------------------------------

func TestDefaultConfig(t *testing.T) {
	cfg := server.DefaultConfig()
	assert.Equal(t, "v0", cfg.Version)
	assert.Equal(t, "dev", cfg.Env)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 3000, cfg.Port)
	assert.Equal(t, "http://localhost:3000/", cfg.Issuer)
	assert.Equal(t, "Hello World", cfg.Message)
}

//-------------------------------------------------------------------------------------------------

func TestConfigAsJSON(t *testing.T) {
	expected := `{ "version": "v0", "env": "dev", "host": "localhost", "port": 3000, "issuer": "http://localhost:3000/", "message": "Hello World" }`
	cfg := server.DefaultConfig()
	actual, err := json.Marshal(cfg)
	require.Nil(t, err)
	assert.JSONEq(t, string(expected), string(actual))
}

//-------------------------------------------------------------------------------------------------
