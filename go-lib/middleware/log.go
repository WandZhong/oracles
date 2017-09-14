package middleware

import (
	"time"

	routing "github.com/go-ozzo/ozzo-routing"
	goweb "github.com/scale-it/go-web"
	"github.com/scale-it/go-web/handlers"
)

func logRequest(s string) { logger.Debug(s) }

// Log tracks the request HTTP info and time
func Log(c *routing.Context) error {
	t := time.Now()
	w := handlers.WrapWriter(c.Response)
	c.Response = w
	// originalPath := r.URL.Path // this can be overwritten by a middleware
	err := c.Next()
	status := w.Status()
	if status == 0 {
		status = 200
	}
	if err, ok := err.(routing.HTTPError); ok {
		status = err.StatusCode()
	}
	goweb.LogRequest(logRequest, c.Request, t, status, w.BytesWritten())
	return err
}
