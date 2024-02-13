package unit

import (
	"errors"
	"openapi/internal/domain/stock/location"
)

type (
	Aggregate struct {
		Id         Id
		Items      ValidItems
		LocationId location.Id
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
	}, nil
}

func RestoreAggregate(id Id, items ValidItems, locationId location.Id) (*Aggregate, error) {
	return &Aggregate{
		Id:         id,
		Items:      items,
		LocationId: locationId,
	}, nil
}

func (a *Aggregate) ChangeLocation(locationId location.Id) {
	a.LocationId = locationId
}
