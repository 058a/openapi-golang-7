package item

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/item"
)

type (
	createRequest struct {
		Name item.Name
	}
	createResponse struct {
		Id uuid.UUID
	}
)

func NewCreateRequest(name string) (createRequest, error) {
	// validation
	validName, err := item.NewName(name)
	if err != nil {
		return createRequest{}, err
	}

	// post processing
	return createRequest{
		Name: validName,
	}, nil
}

func newCreateResponse(id item.Id, name item.Name) createResponse {
	return createResponse{
		Id: id.UUID(),
	}
}

func Create(req createRequest, r item.IRepository, newId uuid.UUID) (createResponse, error) {
	id, err := item.NewId(newId)
	if err != nil {
		return createResponse{}, err
	}

	a := item.NewAggregate(id, req.Name)

	if err := r.Save(a); err != nil {
		return createResponse{}, err
	}

	res := newCreateResponse(a.Id, a.Name)
	return res, nil
}
