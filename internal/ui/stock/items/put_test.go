package items_test

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stock/item"
	"openapi/internal/ui/stock/items"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestPutOk(t *testing.T) {
	t.Parallel()

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "test",
	}
	postReq := NewRequest(http.MethodPost, "/stock/items", postReqBody)
	err := items.Api.PostStockItem(items.Api{}, postReq.context)
	if err != nil {
		t.Fatal(err)
	}
	defer postReq.recorder.Result().Body.Close()

	postResBody, err := Response[oapicodegen.Created](postReq.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: "newTest",
	}
	putReq := NewRequest(http.MethodPut, "/stock/items", putReqBody)
	err = items.Api.PutStockItem(items.Api{}, putReq.context, postResBody.Id)

	// Then
	if err != nil {
		t.Fatal(err)
	}
	defer putReq.recorder.Result().Body.Close()

	if putReq.recorder.Code != http.StatusOK {
		t.Errorf("%T %d want %d", putReq.recorder.Code, putReq.recorder.Code, http.StatusOK)
	}
}

func TestPutNotFound(t *testing.T) {
	t.Parallel()

	putReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "newTest",
	}
	putReq := NewRequest(http.MethodPut, "/stock/items", putReqBody)
	err := items.Api.PutStockItem(items.Api{}, putReq.context, uuid.New())

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	}
	defer putReq.recorder.Result().Body.Close()

	if err.(*echo.HTTPError).Code != http.StatusNotFound {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusNotFound)
	}
}

func TestPutBadRequestNameEmpty(t *testing.T) {
	t.Parallel()

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "test",
	}
	postReq := NewRequest(http.MethodPost, "/stock/items", postReqBody)
	err := items.Api.PostStockItem(items.Api{}, postReq.context)
	if err != nil {
		t.Fatal(err)
	}
	defer postReq.recorder.Result().Body.Close()

	postResBody, err := Response[oapicodegen.Created](postReq.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: "",
	}
	req := NewRequest(http.MethodPut, "/stock/items", putReqBody)
	err = items.Api.PutStockItem(items.Api{}, req.context, postResBody.Id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	}

	if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}

func TestPutBadRequestNameMaxLengthOver(t *testing.T) {
	t.Parallel()

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "test",
	}
	postReq := NewRequest(http.MethodPost, "/stock/items", postReqBody)
	err := items.Api.PostStockItem(items.Api{}, postReq.context)
	if err != nil {
		t.Fatal(err)
	}
	defer postReq.recorder.Result().Body.Close()

	postResBody, err := Response[oapicodegen.Created](postReq.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: strings.Repeat("a", 101),
	}
	req := NewRequest(http.MethodPut, "/stock/items", putReqBody)
	err = items.Api.PutStockItem(items.Api{}, req.context, postResBody.Id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	}

	if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}
