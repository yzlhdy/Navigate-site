package repository

import (
	"navigate/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	// 查询所有用户
	FindAll(page int, limit int) []entity.User
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	DeleteUser(id int) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// 加密密码
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func (u *userRepository) FindAll(page int, limit int) []entity.User {
	var users []entity.User
	u.db.Limit(limit).Offset(page).Find(&users)
	return users
}

func (u *userRepository) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	u.db.Create(&user)
	return user
}

func (u *userRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}

	u.db.Save(&user)
	return user
}

func (u *userRepository) DeleteUser(id int) entity.User {
	var user entity.User
	u.db.Where("id = ?", id).Delete(&user)
	return user
}

func (u *userRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	u.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil
	}
	return user
}

func (u *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	tx = u.db.Where("email = ?", email).First(&entity.User{})
	return
}
