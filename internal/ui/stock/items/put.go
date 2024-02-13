package items

import (
	"net/http"

	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/item"
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/item"
	infra "openapi/internal/infra/repository/sqlboiler/stock/item"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func (Api) PutStockItem(ctx echo.Context, stockItemId openapi_types.UUID) error {
	// Precondition
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo, err := infra.NewRepository(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Binding
	req := &oapicodegen.PutStockItemJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validation
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqDto, err := app.NewUpdateRequest(stockItemId, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	found, err := repo.Find(reqDto.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock item not found")
	}

	// Main
	err = app.Update(reqDto, repo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postcondition
	return ctx.JSON(http.StatusOK, nil)
}
