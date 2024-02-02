package aggregate

import (
	"github.com/google/uuid"
	"github.com/noldwidjaja/slate/arango"
	"integration-test/app/pkg/user/domain/model/request"
)

type User struct {
	arango.DocumentModel
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

func ConvertToUser(request request.UserRequest) *User {
	return &User{
		Id:          uuid.New(),
		Name:        request.Name,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
}

func ConvertToUserUpdate(request request.UserRequest, id uuid.UUID) *User {
	return &User{
		Id:          id,
		Name:        request.Name,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
}
