package converter

import (
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"log"
)

func RoleToResponse(role *entity.Role) *model.RoleResponse {
	log.Println("log from role to response")

	reponse := &model.RoleResponse{
		ID:          role.ID.String(),
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt.Format("2006-01-02"),
	}

	return reponse
}

func RoleToResponses(roles *[]entity.Role) *[]model.RoleResponse {
	var roleResponses []model.RoleResponse

	log.Println("log from role to responses")

	for _, role := range *roles {
		roleResponses = append(roleResponses, *RoleToResponse(&role))
	}

	return &roleResponses
}
