package item_test

import (
	"testing"

	"github.com/google/uuid"

	"openapi/internal/domain/stock/item"
)

func TestNewId(t *testing.T) {
	t.Parallel()

	// When
	value := uuid.New()
	id, err := item.NewId(value)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if id.UUID() != value {
		t.Errorf("%T %+v want %+v", id.UUID(), id.UUID(), value)
	}

	if id.String() != value.String() {
		t.Errorf("%T %+v want %+v", id.String(), id.String(), value)
	}
}

func TestNewIdFail(t *testing.T) {
	t.Parallel()

	// When
	value := uuid.Nil
	id, err := item.NewId(value)

	// Then
	if err != item.ErrInvalidId {
		t.Errorf("%T %+v want %+v", err, err, item.ErrInvalidId)
	}

	if id.UUID() != uuid.Nil {
		t.Errorf("%T %+v want %+v", id.UUID(), id.UUID(), uuid.Nil)
	}
}
