package account

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// User is extended
type User struct {
	Name     string   `jsonapi:"primary,users" gorm:"primaryKey"`
	Email    string   `jsonapi:"attr,email" gorm:"unique"`
	Password string   `jsonapi:"attr,password"`
	Role     UserRole `jsonapi:"attr,role"`
}

// UserContextKey is exported
const UserContextKey = "user"

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord(db *gorm.DB) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetRequestWithUser(req *http.Request) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), UserContextKey, user))
}
