package items

import (
	"net/http"

	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/item"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/item"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func (Api) DeleteStockItem(ctx echo.Context, stockItemId openapi_types.UUID) error {
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

	// Validation
	reqDto, err := app.NewDeleteRequest(stockItemId)
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
	if err := app.Delete(reqDto, repo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postcondition
	return ctx.JSON(http.StatusOK, nil)
}
