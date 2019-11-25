package maintenance

import (
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const maintFileForTesting = "test_maint.txt"

func handlerFuncForTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func TestIsDownForMaintenanceReturnFalse(t *testing.T) {
	mw := NewMaintMiddleware(maintFileForTesting, handlerFuncForTest)
	assert.False(t, mw.isDownForMaintenance())
}

func TestIsDownForMaintenanceReturnTrue(t *testing.T) {
	mw := NewMaintMiddleware(maintFileForTesting, handlerFuncForTest)
	os.Create(maintFileForTesting)
	assert.True(t, mw.isDownForMaintenance())
	os.Remove(maintFileForTesting)
}
