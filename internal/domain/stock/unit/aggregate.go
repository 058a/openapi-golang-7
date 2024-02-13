package unit

import (
	"errors"
	"openapi/internal/domain/stock/item"
	"openapi/internal/domain/stock/location"
)

type (
	Aggregate struct {
		Id         Id
		Items      ValidItems
		LocationId location.Id
		Deleted    bool
	}
)

var (
	ErrZeroItems = errors.New("zero items")
)

func NewAggregate(id Id, items ValidItems, locationId location.Id) (*Aggregate, error) {
	return &Aggregate{
		Id:         id,
		Items:      items,
		LocationId: locationId,
		Deleted:    false,
	}, nil
}

func RestoreAggregate(id Id, items ValidItems, locationId location.Id, deleted bool) (*Aggregate, error) {
	return &Aggregate{
		Id:         id,
		Items:      items,
		LocationId: locationId,
		Deleted:    deleted,
	}, nil
}

func (a *Aggregate) ChangeLocation(locationId location.Id) {
	a.LocationId = locationId
}

func (o *Aggregate) RemoveItem(id item.Id) {
	itemCount := len(o.Items.Items())
	for k, v := range o.Items.Items() {
		if v.Id == id {
			v.deleted = true
			o.Items.Items()[k] = v
		}
		if v.deleted {
			itemCount--
		}
	}

	if itemCount == 0 {
		o.Deleted = true
	}
}
