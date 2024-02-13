package unit_test

import (
	"openapi/internal/domain/stock/item"
	"openapi/internal/domain/stock/unit"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewUnverifiedItems(t *testing.T) {
	// When
	items := unit.NewUnverifiedItems()

	// Then
	if items.Count() != 0 {
		t.Errorf("%T %+v want %+v", items.Count(), items.Count(), 0)
	}
}

func TestUnverifiedItemsAdd(t *testing.T) {
	// Given
	unverifiedItems := unit.NewUnverifiedItems()

	// When
	itemId, err := item.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	unverifiedItems.Add(itemId)
	unverifiedItems.Add(itemId)

	newItemId, err := item.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	unverifiedItems.Add(newItemId)

	unverifiedItems.Remove(itemId)

	validItems, err := unverifiedItems.Verify()
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if unverifiedItems.Count() != 2 {
		t.Errorf("%T %+v want %+v", unverifiedItems.Count(), unverifiedItems.Count(), 2)
	}

	if !reflect.DeepEqual(unverifiedItems.Items(), validItems.Items()) {
		t.Errorf("%T %+v want %+v", unverifiedItems, unverifiedItems, validItems)
	}

}
