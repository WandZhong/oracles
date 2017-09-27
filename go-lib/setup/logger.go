package setup

import (
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/robert-zaremba/log15/rollbar"
)

// MustLogger setups logger
func MustLogger(appname string, rollbartoken string) {
	assert(checkAppName(appname))
	appname = envName + "-" + appname
	logger.Info("Rollbar", "token", rollbartoken)
	rc := rollbar.Config{
		Version: GitVersion,
		Env:     appname,
		Token:   rollbartoken}
	var err error
	logger, err = log.New(log.RootName,
		log.Config{Color: true, TimeFmt: "sec", Level: "DEBUG"}, rc)
	assert(err)
	if strings.HasPrefix(envName, "prod") && rollbartoken == "" {
		logger.Error("Rollbar token must be set in production environment")
	}
}
