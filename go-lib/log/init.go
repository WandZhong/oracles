package log

import (
	"errors"
	"os"

	"github.com/robert-zaremba/log15"
	"github.com/robert-zaremba/log15/rollbar"
)

var root = log15.Get("sweetbridge")

// timeFMT is the map of predefined time formats for TerminalFmt
var timeFMT = map[string]string{
	"off":   "",
	"date":  " 01-02 ",
	"sec":   " 01-02 15:04:05 ",
	"musec": " 01-02 15:04:05.0000 ",
}

// A Logger writes key/value pairs to a Handler
type Logger interface {
	// New returns a new Logger that has this logger's context plus the given context
	New(ctx ...interface{}) log15.Logger

	// Log a message at the given level with context key/value pairs
	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})

	// Fatal is a Crit log followed by panic
	Fatal(msg string, ctx ...interface{})
}

// Config represents  log config
type Config struct {
	Color   bool   `yaml:"color"`
	TimeFmt string `yaml:"timeFmt"` // one of timeFMT values
	Level   string `yaml:"level"`

	lvl log15.Lvl
}

// Check validates the config content
func (c *Config) Check() error {
	if _, ok := timeFMT[c.TimeFmt]; !ok {
		return errors.New("Wrong timeFmt value, should be one of log.timeFMT values")
	}
	var err error
	c.lvl, err = log15.LvlFromString(c.Level)
	return err
}

// New constructs logger and registers it in the log15 repository.
// If the logger `name` is already registered then it's upgraded according to given config.
func New(name string, c Config, rc rollbar.Config) (Logger, error) {
	err := c.Check()
	if err != nil {
		return nil, err
	}

	f := log15.TerminalFormat{WithColor: c.Color, TimeFmt: timeFMT[c.TimeFmt]}
	h := log15.StreamHandler(os.Stderr, f)
	h = log15.SyncHandler(h)
	stderrHandler := h
	if rc.Token != "" {
		if err = rc.Check(); err != nil {
			return nil, err
		}
		h = log15.MultiHandler(h, rollbar.MkHandler(rc))

		rollbarLogger := log15.Get("rollbar")
		rollbarLogger.SetHandler(stderrHandler)
		go rollbar.LogInternalErrors(rollbarLogger)
	}
	h = log15.CallerFileHandler(h, true)
	h = log15.LvlFilterHandler(c.lvl, h)
	l := log15.Get(name)
	l.SetHandler(h)
	if rc.Token == "" {
		l.Debug("Rollbar token not set. Disabling rollbar integration.")
	}
	return l, nil
}

// Root returns the default logger
func Root() Logger {
	return root
}
