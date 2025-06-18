package main

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"void.dev/identity/internal/server"
)

//-------------------------------------------------------------------------------------------------

const (
	name            = "void-identity"
	summary         = "The Void Identity Service."
	usageText       = "server --env ENV --host HOST --port PORT"
	envShortArg     = "e"
	envArg          = "env"
	envEnvVar       = "ENV"
	envUsage        = "environment (dev | test | stage | prod)"
	hostArg         = "host"
	hostEnvVar      = "HOST"
	hostUsage       = "server host"
	portShortArg    = "p"
	portArg         = "port"
	portEnvVar      = "PORT"
	portUsage       = "server port"
	issuerShortArg  = "i"
	issuerArg       = "issuer"
	issuerEnvVar    = "ISSUER"
	issuerUsage     = "issuer url"
	messageShortArg = "m"
	messageArg      = "message"
	messageEnvVar   = "MESSAGE"
	messageUsage    = "provide a custom message"
)

//-------------------------------------------------------------------------------------------------

func main() {

	app := &cli.App{
		Name:      name,
		Usage:     summary,
		UsageText: usageText,
		Version:   "1.0",
		Action:    RunServer,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    strings.Join([]string{envArg, envShortArg}, ","),
				Value:   "dev",
				Usage:   envUsage,
				EnvVars: []string{envEnvVar},
			},
			&cli.StringFlag{
				Name:    strings.Join([]string{hostArg}, ","),
				Value:   "localhost",
				Usage:   hostUsage,
				EnvVars: []string{hostEnvVar},
			},
			&cli.IntFlag{
				Name:    strings.Join([]string{portArg, portShortArg}, ","),
				Value:   3000,
				Usage:   portUsage,
				EnvVars: []string{portEnvVar},
			},
			&cli.StringFlag{
				Name:    strings.Join([]string{messageArg, messageShortArg}, ","),
				Value:   "hello world",
				Usage:   messageUsage,
				EnvVars: []string{messageEnvVar},
			},
			&cli.StringFlag{
				Name:    strings.Join([]string{issuerArg, issuerShortArg}, ","),
				Value:   "http://localhost:3000/oidc/",
				Usage:   issuerUsage,
				EnvVars: []string{issuerEnvVar},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

//-------------------------------------------------------------------------------------------------

func RunServer(c *cli.Context) error {

	cfg := server.NewConfig(
		c.App.Version,
		c.String(envArg),
		c.String(hostArg),
		c.Int(portArg),
		c.String(issuerArg),
		c.String(messageArg))

	server, err := server.NewServer(cfg)
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	return nil

}

//-------------------------------------------------------------------------------------------------
