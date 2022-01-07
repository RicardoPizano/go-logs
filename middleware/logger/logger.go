package logger

import (
	"github.com/RicardoPizano/go-logs/logger"
	"time"

	"github.com/labstack/echo"
)

// Logger :
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {

		start := time.Now()
		_ = next(c)
		defer logger.Request(c.Request().Method, c.Response().Status, c.Request().RequestURI, start)

		return
	}
}
