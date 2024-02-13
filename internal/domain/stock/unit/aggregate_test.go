package unit_test

import (
	"openapi/internal/domain/stock/item"
	"openapi/internal/domain/stock/location"
	"openapi/internal/domain/stock/unit"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	t.Parallel()

	// Given
	id, err := unit.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	unverifiedItems := unit.NewUnverifiedItems()
	for i := 0; i < 3; i++ {
		itemId, err := item.NewId(uuid.New())
		if err != nil {
			t.Fatal(err)
		}
		unverifiedItems.Add(itemId)
	}

	validItems, err := unverifiedItems.Verify()
	if err != nil {
		t.Fatal(err)
	}

	locationId, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	a, err := unit.RestoreAggregate(id, validItems, locationId, false)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.Id != id {
		t.Errorf("%T %+v want %+v", a.Id, a.Id, id)
	}

	if reflect.DeepEqual(a.Items.Items(), validItems.Items()) != true {
		t.Errorf("%T %+v want %+v", a.Items, a.Items, validItems)
	}

	if a.LocationId != locationId {
		t.Errorf("%T %+v want %+v", a.LocationId, a.LocationId, locationId)
	}
}

func TestRestoreAggregate(t *testing.T) {
	t.Parallel()

	// Given
	id, err := unit.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	unverifiedItems := unit.NewUnverifiedItems()
	for i := 0; i < 3; i++ {
		itemId, err := item.NewId(uuid.New())
		if err != nil {
			t.Fatal(err)
		}
		unverifiedItems.Add(itemId)
	}

	validItems, err := unverifiedItems.Verify()
	if err != nil {
		t.Fatal(err)
	}

	locationId, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	a, err := unit.RestoreAggregate(id, validItems, locationId, false)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.Id != id {
		t.Errorf("%T %+v want %+v", a.Id, a.Id, id)
	}

	if reflect.DeepEqual(a.Items.Items(), validItems.Items()) != true {
		t.Errorf("%T %+v want %+v", a.Items, a.Items, validItems)
	}

	if a.LocationId != locationId {
		t.Errorf("%T %+v want %+v", a.LocationId, a.LocationId, locationId)
	}
}

func TestChangeLocation(t *testing.T) {
	t.Parallel()

	// Given
	id, err := unit.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	unverifiedItems := unit.NewUnverifiedItems()
	for i := 0; i < 3; i++ {
		itemId, err := item.NewId(uuid.New())
		if err != nil {
			t.Fatal(err)
		}
		unverifiedItems.Add(itemId)
	}

	validItems, err := unverifiedItems.Verify()
	if err != nil {
		t.Fatal(err)
	}

	locationId, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	a, err := unit.NewAggregate(id, validItems, locationId)
	if err != nil {
		t.Fatal(err)
	}

	// When
	newLocationId, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	a.ChangeLocation(newLocationId)

	// Then
	if a.LocationId != newLocationId {
		t.Errorf("%T %+v want %+v", a.LocationId, a.LocationId, newLocationId)
	}
}

func TestRemoveItems(t *testing.T) {
	t.Parallel()

	// Given
	id, err := unit.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	unverifiedItems := unit.NewUnverifiedItems()
	for i := 0; i < 3; i++ {
		itemId, err := item.NewId(uuid.New())
		if err != nil {
			t.Fatal(err)
		}
		unverifiedItems.Add(itemId)
	}

	validItems, err := unverifiedItems.Verify()
	if err != nil {
		t.Fatal(err)
	}

	locationId, err := location.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	a, err := unit.NewAggregate(id, validItems, locationId)
	if err != nil {
		t.Fatal(err)
	}

	// When
	var count = 3
	for _, v := range a.Items.Items() {
		count--
		a.RemoveItem(v.Id)
		if count > 0 {
			if a.Deleted != false {
				t.Errorf("%T %+v want %+v", a.Deleted, a.Deleted, true)
			}
			continue
		}
		if a.Deleted != true {
			t.Errorf("%T %+v want %+v", a.Deleted, a.Deleted, false)
		}

	}
}
