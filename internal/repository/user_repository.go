package repository

import (
	entity "klinikserver/internal/entity/authentication"
	"klinikserver/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByIdForUpdate(tx *gorm.DB, user *entity.User) error
	FindAllForPagging(tx *gorm.DB, userId uuid.UUID, pageSize int, offset int, order string, sortBy string, users *[]entity.User) (int64, error)
	FindById(tx *gorm.DB, user *entity.User) error
	FindByUsername(tx *gorm.DB, user *entity.User) error
	Login(tx *gorm.DB, user *model.LoginResult) error
	Create(tx *gorm.DB, user *entity.User) error
	Update(tx *gorm.DB, user *entity.User) error
	Delete(tx *gorm.DB, user *entity.User) error
}

type UserRepositoryImpl struct {
	Repository[entity.User]
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// FindByIdForUpdate implements UserRepository.
func (repository *UserRepositoryImpl) FindByIdForUpdate(tx *gorm.DB, user *entity.User) error {
	query := tx.Model(&entity.User{})

	return query.First(&user).Error
}

func (repository *UserRepositoryImpl) FindAllForPagging(tx *gorm.DB, userId uuid.UUID, pageSize int, offset int, order string, sortBy string, users *[]entity.User) (int64, error) {
	var total int64
	query := tx.Model(&entity.User{})

	sortColumn := "users.created_at"

	query = query.Not("users.id = ?", userId)

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
		Find(&users).Error

	return total, err
}

// FindByUsername implements UserRepository.
func (repository *UserRepositoryImpl) FindByUsername(tx *gorm.DB, user *entity.User) error {
	return tx.First(user, "username=?", user.Username).Error
}

// FindById implements UserRepository.
func (repository *UserRepositoryImpl) FindById(tx *gorm.DB, user *entity.User) error {
	return tx.First(user).Error
}

// Login implements UserRepository.
func (repository *UserRepositoryImpl) Login(tx *gorm.DB, user *model.LoginResult) error {
	var rows []model.LoginRow

	err := tx.Model(&entity.User{}).
		Select(`
		users.id,
		users.username,
		users.password_hash,
		users.full_name,
		roles.name as role_name
	`).
		Joins("JOIN user_roles ur ON ur.user_id = users.id").
		Joins("JOIN roles ON roles.id = ur.role_id").
		Where("users.username = ?", user.Username).
		Scan(&rows).Error

	if err != nil {
		return err
	}

	if len(rows) == 0 {
		return gorm.ErrRecordNotFound
	}

	user.ID = rows[0].ID
	user.Username = rows[0].Username
	user.PasswordHash = rows[0].PasswordHash
	user.FullName = rows[0].FullName

	for _, row := range rows {
		user.RoleName = append(user.RoleName, row.RoleName)
	}

	err = tx.Model(&entity.User{}).
		Where("username = ?", user.Username).
		UpdateColumn("last_login_at", time.Now()).Error
	if err != nil {
		return err
	}

	// switch user.Role {
	// case "coach":
	// 	return tx.Preload("Coach", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("user_id", "full_name")
	// 	}).First(user, user.ID).Error
	// case "student":
	// 	return tx.Preload("Student", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("user_id", "full_name")
	// 	}).First(user, user.ID).Error
	// }

	return nil
}
