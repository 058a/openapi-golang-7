package items

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stock/item"

	"github.com/labstack/echo/v4"
)

type Api struct {
	oapicodegen.ServerInterface
}

func RegisterHandlers(e *echo.Echo) {
	oapicodegen.RegisterHandlers(e, &Api{})
}
