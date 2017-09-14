package middleware

import (
	"github.com/go-ozzo/ozzo-routing"
)

// StdRouter defines standard, default router
func StdRouter() *routing.Router {
	r := routing.New()
	r.Use(Log)
	return r
}
