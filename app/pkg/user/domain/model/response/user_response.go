package response

import (
	"github.com/google/uuid"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"time"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func BuildFindUserResponse(user *aggregate.User) *UserResponse {
	return &UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Address:     user.Address,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
