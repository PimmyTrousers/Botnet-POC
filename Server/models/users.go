package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Pimmytrousers/GoSwarm/server/hash"
	"github.com/Pimmytrousers/GoSwarm/server/rand"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	userPwPepper  = "DoompyPoomb1234567890"
	hmacSecretKey = "Raya-4-1-18"
)

var (
	// ErrNotFound is returned when a resource cannot be found // in the database.
	ErrNotFound        = errors.New("models: resource not found")
	ErrInvalidID       = errors.New("models: invalid ID given")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

type UserDB interface {
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	Close() error

	AutoMigrate() error
	DestructiveReset() error
}

type UserService interface {
	Authenticate(email, password string) (*User, error)
	UserDB
}

// Our user model in which are user info is kept

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

// The various service layers

type userService struct {
	UserDB
}

type userValidator struct {
	UserDB
	hmac hash.HMAC
}

type userGorm struct {
	db *gorm.DB
}

type userValFn func(*User) error

func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}

	hmac := hash.NewHMAC(hmacSecretKey)
	uv := &userValidator{
		hmac:   hmac,
		UserDB: ug,
	}

	return &userService{
		UserDB: uv,
	}, nil
}

func (us *userService) Authenticate(email string, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+userPwPepper))

	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}
}

// Define out services here!
var _ UserDB = &userGorm{}
var _ UserService = &userService{}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &userGorm{
		db: db,
	}, nil
}

func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// ByRemember looks up a user with the given remember token
// and returns that user. This method expects the remember
// token to already be hashed.
func (ug *userGorm) ByRemember(rememberHash string) (*User, error) {
	var user User
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uv *userValidator) ByRemember(token string) (*User, error) {
	RememberHash := uv.hmac.Hash(token)
	return uv.UserDB.ByRemember(RememberHash)
}

func (uv *userValidator) Create(user *User) error {
	if err := runUserValFns(user,
		uv.bcryptPassword); err != nil {
		return err
	}
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = uv.hmac.Hash(user.Remember)
	return uv.UserDB.Create(user)
}

func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error
}

func (uv *userValidator) Update(user *User) error {
	if user.Password == "" {
		return nil
	}
	if err := runUserValFns(user,
		uv.bcryptPassword); err != nil {
		return err
	}
	if user.Remember != "" {
		user.RememberHash = uv.hmac.Hash(user.Remember)
	}
	return uv.UserDB.Update(user)
}

func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}

func (uv *userValidator) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	return uv.UserDB.Delete(id)
}

func (ug *userGorm) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (ug *userGorm) DestructiveReset() error {
	err := ug.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return ug.AutoMigrate()
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

func (uv *userValidator) bcryptPassword(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes,
		bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return nil
}

func runUserValFns(user *User, fns ...userValFn) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil

}
