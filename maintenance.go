package maintenance

import (
	"os"

	"github.com/labstack/echo/v4"
)

func (middleware *Middleware) isDownForMaintenance() bool {
	if _, err := os.Stat(middleware.filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// Middleware struct
type Middleware struct {
	filePath      string
	customHandler echo.HandlerFunc
}

// NewMaintMiddleware - Creates a *Middleware instance
// with maint file path and custom response struct
func NewMaintMiddleware(filePath string, customHandler echo.HandlerFunc) *Middleware {
	mw := new(Middleware)
	mw.filePath = filePath
	mw.customHandler = customHandler
	return mw
}

// CheckMaintenance - middleware to check if maintenance mode is on
func (middleware *Middleware) CheckMaintenance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if middleware.isDownForMaintenance() {
			return middleware.customHandler(c)
		}
		return next(c)
	}
}
