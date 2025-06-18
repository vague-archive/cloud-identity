package server

//-------------------------------------------------------------------------------------------------

var (
	DefaultVersion = "v0"
	DefaultEnv     = "dev"
	DefaultHost    = "localhost"
	DefaultPort    = 3000
	DefaultIssuer  = "http://localhost:3000/"
	DefaultMessage = "Hello World"
)

//-------------------------------------------------------------------------------------------------

type Config struct {
	Version string `json:"version"`
	Env     string `json:"env"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Issuer  string `json:"issuer"`
	Message string `json:"message"`
}

//-------------------------------------------------------------------------------------------------

func NewConfig(version, env, host string, port int, issuer, message string) *Config {
	return &Config{
		Version: version,
		Env:     env,
		Host:    host,
		Port:    port,
		Issuer:  issuer,
		Message: message,
	}
}

func DefaultConfig() *Config {
	return NewConfig(
		DefaultVersion,
		DefaultEnv,
		DefaultHost,
		DefaultPort,
		DefaultIssuer,
		DefaultMessage)
}

//-------------------------------------------------------------------------------------------------
