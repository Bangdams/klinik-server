package repository

import (
	entity "klinikserver/internal/entity/authentication"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll(tx *gorm.DB, name string, pageSize int, offset int, order string, roles *[]entity.Role) (int64, error)
	FindByName(tx *gorm.DB, role *entity.Role) error
	FindById(tx *gorm.DB, role *entity.Role) error
	Create(tx *gorm.DB, role *entity.Role) error
	Update(tx *gorm.DB, role *entity.Role) error
	Delete(tx *gorm.DB, role *entity.Role) error
}

type RoleRepositoryImpl struct {
	Repository[entity.Role]
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

// FindAll implements [RoleRepository].
func (repository *RoleRepositoryImpl) FindAll(tx *gorm.DB, name string, pageSize int, offset int, order string, roles *[]entity.Role) (int64, error) {
	var total int64
	query := tx.Model(&entity.Role{})

	sortColumn := "roles.created_at"

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	validOrder := "DESC"
	if order == "ASC" {
		validOrder = "ASC"
	}

	err := query.Order(sortColumn + " " + validOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&roles).Error

	return total, err
}

// FindByName implements [RoleRepository].
func (repository *RoleRepositoryImpl) FindByName(tx *gorm.DB, role *entity.Role) error {
	return tx.
		Model(&entity.Role{}).
		Where("name = ?", role.Name).
		First(role).
		Error
}

// FindById implements [RoleRepository].
func (repository *RoleRepositoryImpl) FindById(tx *gorm.DB, role *entity.Role) error {
	return tx.
		Model(&entity.Role{}).
		Where("id = ?", role.ID).
		First(role).
		Error
}
