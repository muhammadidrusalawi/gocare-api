package repository

import (
	"encoding/json"
	"time"

	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Upsert(user *model.User) error
	UpdatePassword(id string, password string) (*model.User, error)
	UpdateUser(id string, fields map[string]interface{}) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) cacheGet(key string, dest interface{}) bool {
	val, err := cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), dest)
		return true
	}
	return false
}

func (r *userRepository) cacheSet(key string, value interface{}, ttl time.Duration) {
	b, _ := json.Marshal(value)
	cache.Client.Set(cache.Ctx, key, b, ttl)
}

func (r *userRepository) cacheDel(keys ...string) {
	cache.Client.Del(cache.Ctx, keys...)
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	cacheKey := "user:id:" + id
	var user model.User

	if r.cacheGet(cacheKey, &user) {
		return &user, nil
	}

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, user, 10*time.Minute)
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	cacheKey := "user:email:" + email
	var user model.User

	if r.cacheGet(cacheKey, &user) {
		return &user, nil
	}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, user, 10*time.Minute)
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	r.cacheDel(
		"user:id:"+user.ID,
		"user:email:"+user.Email,
	)
	return nil
}

func (r *userRepository) Upsert(user *model.User) error {
	if err := r.db.
		Where("email = ?", user.Email).
		FirstOrCreate(user).Error; err != nil {
		return err
	}

	r.cacheDel(
		"user:id:"+user.ID,
		"user:email:"+user.Email,
	)

	return nil
}

func (r *userRepository) UpdatePassword(id string, password string) (*model.User, error) {
	var user model.User

	if err := r.db.
		Model(&model.User{}).
		Where("id = ?", id).
		Update("password", password).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	r.cacheDel(
		"user:id:"+user.ID,
		"user:email:"+user.Email,
	)

	return &user, nil
}

func (r *userRepository) UpdateUser(id string, fields map[string]interface{}) (*model.User, error) {
	var user model.User

	if err := r.db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	r.cacheDel(
		"user:id:"+user.ID,
		"user:email:"+user.Email,
	)

	return &user, nil
}
