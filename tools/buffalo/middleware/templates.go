package middleware

var middlewareGo = `
// middleware package is intended to host the middlewares used
// across the app.
package middleware

import (
	"{{.Module}}/app/models"

	tx "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

var (
	// Transaction middleware wraps the request with a pop
	// transaction that is committed on success and rolled 
	// back when errors happen.
	Transaction = tx.Transaction(models.DB())

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)
`
