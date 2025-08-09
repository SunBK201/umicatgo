package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/SunBK201/umicatgo/conf"
	"github.com/SunBK201/umicatgo/log"
	"github.com/SunBK201/umicatgo/policy"
)

type Cycle struct {
	Workers   int
	Mode      Mode
	Localport int
	Policy    policy.Policy
	LogFile   string
}

func main() {
	log.SetLogConf()

	app := &cli.App{
		Name:  "umicatgo",
		Usage: "umicat go implementaion",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "conf",
				Aliases: []string{"c"},
				Usage:   "set configuration file",
				Value:   "/etc/umicatgo/umicatgo.conf",
			},
			&cli.StringFlag{
				Name:    "log",
				Aliases: []string{"l"},
				Usage:   "set log file",
				Value:   "/var/log/umicatgo/umicatgo.log",
			},
		},
		Action: func(cCtx *cli.Context) error {
			run(cCtx)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(cCtx *cli.Context) {
	var cycle Cycle

	confFilePath := cCtx.String("conf")
	conf, error := conf.ParseConfFile(confFilePath)
	if error != nil {
		logrus.Fatal(error, "failed to parse config file: ", confFilePath)
	}
	cycle.loadConf(conf)

	ServerPool(cycle)
}

func (cycle *Cycle) loadConf(conf conf.Config) {
	var err error

	if conf.Workers == "auto" {
		cycle.Workers = runtime.NumCPU()
	} else {
		cycle.Workers, err = strconv.Atoi(conf.Workers)
		if err != nil {
			logrus.Fatal(err, "failed to parse workers: ", conf.Workers)
		}
	}

	cycle.Mode, err = ParseMode(conf.Mode)
	if err != nil {
		logrus.Fatal(err, "failed to parse mode: ", conf.Mode)
	}

	cycle.Localport = conf.LocalPort

	cycle.Policy, err = policy.ParsePolicy(conf.Policy)
	if err != nil {
		logrus.Fatal(err, "failed to parse policy: ", conf.Policy)
	}

	cycle.LogFile = conf.LogFile
}

type Mode int

const (
	TCP Mode = iota
	UDP
	HTTP
)

func ParseMode(mode string) (Mode, error) {
	switch strings.ToLower(mode) {
	case "tcp":
		return TCP, nil
	case "udp":
		return UDP, nil
	case "http":
		return HTTP, nil

	}
	var m Mode
	return m, fmt.Errorf("not a valid mode: %q", mode)
}

func (mode Mode) MarshalText() (string, error) {
	switch mode {
	case TCP:
		return "tcp", nil
	case UDP:
		return "udp", nil
	case HTTP:
		return "http", nil
	}
	return "", fmt.Errorf("not a valid mode: %q", mode)
}
