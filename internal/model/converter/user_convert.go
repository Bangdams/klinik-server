package converter

import (
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"log"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	log.Println("log from user to response")

	reponse := &model.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Phone:     user.Phone,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
	}

	return reponse
}

func UserToResponseForUpdate(user *entity.User) *model.UserResponse {
	log.Println("log from UserToResponseForUpdate")

	return &model.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Phone:     user.Phone,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
	}
}

func LoginUserToResponse(user *entity.User) *model.LoginResponse {
	log.Println("log from login user to response")

	loginResponse := &model.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		// Role:     user.Role,
	}

	return loginResponse
}

func UserToResponses(users *[]entity.User) *[]model.UserResponse {
	var userResponses []model.UserResponse

	log.Println("log from user to responses")

	for _, user := range *users {
		userResponses = append(userResponses, *UserToResponse(&user))
	}

	return &userResponses
}
