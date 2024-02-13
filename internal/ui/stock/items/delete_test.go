package items_test

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stock/item"
	"openapi/internal/ui/stock/items"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestDeleteOk2(t *testing.T) {
	t.Parallel()

	// Given
	beforeReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "test",
	}

	b := NewRequest(http.MethodPost, "/stock/items", beforeReqBody)

	if err := items.Api.PostStockItem(items.Api{}, b.context); err != nil {
		t.Fatal(err)
	}
	defer b.recorder.Result().Body.Close()

	if b.recorder.Code != http.StatusCreated {
		t.Fatal(b.recorder.Code)
	}

	postReqBody, err := Response[oapicodegen.Created](b.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: "test",
	}

	a := NewRequest(http.MethodDelete, "/stock/items/"+postReqBody.Id.String(), afterReqBody)

	if err := items.Api.DeleteStockItem(items.Api{}, a.context, postReqBody.Id); err != nil {
		t.Fatal(err)
	}
	defer a.recorder.Result().Body.Close()

	// Then
	if a.recorder.Code != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, a.recorder.Code)
	}
}

func TestDeleteOk(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}

	// Given
	postRes, err := rh.Post(
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: "test",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatalf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBody, err := rch.AsCreated(postRes)
	if err != nil {
		t.Fatal(err)
	}

	// When
	deleteRes, err := rh.Delete(postResBody.Id)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	// Then
	if deleteRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, deleteRes.StatusCode)
	}
}

func TestDeleteNotFound(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}

	deleteRes, err := rh.Delete(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	// Then
	if deleteRes.StatusCode != http.StatusNotFound {
		t.Errorf("%T %d, want %d", deleteRes.StatusCode, deleteRes.StatusCode, http.StatusNotFound)
	}
}
