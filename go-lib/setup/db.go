package setup

import (
	stdlog "log"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/go-pg/pg"
	bat "github.com/robert-zaremba/go-bat"
)

// pgLogger wraps log.Logger to provide io.Writer interface
type pgLogger struct {
	logger log.Logger
}

func (l pgLogger) Write(p []byte) (int, error) {
	logger.Debug(bat.UnsafeByteArrayToStr(p))
	return len(p), nil
}

// MustPsql initiates Postgersql connection. It panics in case of error
func MustPsql() *pg.DB {
	l := pgLogger{logger.New("PG:")}
	pg.SetLogger(stdlog.New(l, "", 0))
	db := pg.Connect(&pg.Options{
		User:     "swc-queue",
		Password: "password",
		Database: "sw",
		Addr:     "localhost:5432"})

	// let's test the connection
	var x int
	_, err := db.QueryOne(&x, "SELECT 1")
	if err != nil || x != 1 {
		logger.Fatal("Can't connect to the Postgresql", err)
	}
	return db
}
