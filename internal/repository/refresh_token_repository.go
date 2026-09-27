package repository

import (
	entity "klinikserver/internal/entity/authentication"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(tx *gorm.DB, refreshToken *entity.RefreshToken) error
	Update(tx *gorm.DB, refreshToken *entity.RefreshToken) error
	FindById(tx *gorm.DB, userId uuid.UUID) error
	CheckStatusLogout(tx *gorm.DB, userId uuid.UUID) error
}

type RefreshTokenRepositoryImpl struct {
	Repository[entity.RefreshToken]
}

func NewRefreshTokenRepository() RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{}
}

// CheckStatusLogout implements UserRepository.
func (repository *RefreshTokenRepositoryImpl) CheckStatusLogout(tx *gorm.DB, userId uuid.UUID) error {
	return tx.
		First(&entity.RefreshToken{}, "user_id = ? AND status_logout = ?", userId, 0).
		Error
}

// FindById implements UserRepository.
func (repository *RefreshTokenRepositoryImpl) FindById(tx *gorm.DB, userId uuid.UUID) error {
	return tx.
		First(&entity.RefreshToken{}, "user_id=?", userId).
		Error
}
