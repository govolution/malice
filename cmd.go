package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/maliceio/malice/commands"
	er "github.com/maliceio/malice/libmalice/errors"
	"github.com/maliceio/malice/libmalice/logger"
	"github.com/maliceio/malice/plugins"
	"github.com/maliceio/malice/version"
)

func init() {
	// config.Load()
	logger.Init()
	setDebugOutputLevel()
	plugins.Load()
}

func setDebugOutputLevel() {
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			log.SetLevel(log.DebugLevel)
		}
	}

	debugEnv := os.Getenv("MALICE_DEBUG")
	if debugEnv != "" {
		showDebug, err := strconv.ParseBool(debugEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean value from MALICE_DEBUG: %s\n", err)
			os.Exit(1)
		}
		if showDebug {
			log.SetLevel(log.DebugLevel)
		}
	}
}

// Init initializes Malice
func Init() {

}

func main() {
	Init()
	// setDebugOutputLevel()
	cli.AppHelpTemplate = commands.AppHelpTemplate
	cli.CommandHelpTemplate = commands.CommandHelpTemplate
	app := cli.NewApp()

	app.Name = "malice"
	app.Author = "blacktop"
	app.Email = "https://github.com/blacktop"

	app.Commands = commands.Commands
	app.CommandNotFound = commands.CmdNotFound
	app.Usage = "Open Source Malware Analysis Framework"
	app.Version = version.FullVersion()
	// app.EnableBashCompletion = true

	log.Debug("Malice Version: ", app.Version)

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			EnvVar: "MALICE_DEBUG",
			Name:   "debug, D",
			Usage:  "Enable debug mode",
		},
	}

	err := app.Run(os.Args)
	er.CheckError(err)
}
