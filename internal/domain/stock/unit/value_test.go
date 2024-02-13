package unit_test

import (
	"errors"
	"openapi/internal/domain/stock/item"
	"openapi/internal/domain/stock/unit"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewUnverifiedItems(t *testing.T) {
	t.Parallel()

	// When
	items := unit.NewUnverifiedItems()

	// Then
	if items.Count() != 0 {
		t.Errorf("%T %+v want %+v", items.Count(), items.Count(), 0)
	}
}

func TestUnverifiedItemsAdd(t *testing.T) {
	t.Parallel()

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

func TestUnverifiedItemsFailVerify(t *testing.T) {
	t.Parallel()

	// Given
	unverifiedItems := unit.NewUnverifiedItems()

	// When
	_, err := unverifiedItems.Verify()

	// Then
	if !errors.Is(err, unit.ErrItemsZero) {
		t.Errorf("%T %+v want %+v", err, err, unit.ErrItemsZero)
	}
}
