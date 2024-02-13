package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type updateRequest struct {
	Id   item.Id
	Name item.Name
}

func NewUpdateRequest(id uuid.UUID, name string) (updateRequest, error) {
	// validation
	validId, err := item.NewId(id)
	if err != nil {
		return updateRequest{}, err
	}

	validName, err := item.NewName(name)
	if err != nil {
		return updateRequest{}, err
	}

	// post processing
	return updateRequest{
		Id:   validId,
		Name: validName,
	}, nil
}

func Update(req updateRequest, r item.IRepository) error {
	id := item.Id(req.Id)
	a, err := r.Get(id)
	if err != nil {
		return err
	}

	a.Name = req.Name

	err = r.Save(a)
	if err != nil {
		return err
	}

	return nil
}
