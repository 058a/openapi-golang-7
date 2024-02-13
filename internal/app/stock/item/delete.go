package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type deleteRequest struct {
	Id item.Id
}

func NewDeleteRequest(id uuid.UUID) (deleteRequest, error) {
	// validation
	validId, err := item.NewId(id)
	if err != nil {
		return deleteRequest{}, err
	}

	// post processing
	return deleteRequest{
		Id: validId,
	}, nil
}

func Delete(req deleteRequest, r item.IRepository) error {
	a, err := r.Get(req.Id)
	if err != nil {
		return err
	}

	a.Delete()

	err = r.Save(a)
	if err != nil {
		return err
	}

	return nil
}
