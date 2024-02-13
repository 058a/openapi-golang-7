package unit

import (
	"errors"
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type (
	Item struct {
		Id      item.Id
		added   bool
		deleted bool
	}

	Items map[uuid.UUID]Item

	UnverifiedItems struct {
		items Items
	}

	ValidItems struct {
		items Items
	}
)

var (
	ErrItemsZero = errors.New("zero items")
)

func NewUnverifiedItems() UnverifiedItems {
	items := make(Items)
	return UnverifiedItems{items}
}

func (o *UnverifiedItems) Add(id item.Id) {
	o.items[id.UUID()] = Item{
		Id:      id,
		added:   true,
		deleted: false,
	}
}

func (o UnverifiedItems) Verify() (ValidItems, error) {
	if len(o.items) == 0 {
		return ValidItems{}, ErrItemsZero
	}
	items := make(Items, len(o.items))
	for k, v := range o.items {
		items[k] = v
	}
	return ValidItems{items}, nil
}

func (o UnverifiedItems) Count() int {
	return len(o.items)
}

func (o UnverifiedItems) Items() Items {
	return o.items
}

func (o ValidItems) Items() Items {
	return o.items
}
