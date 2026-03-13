package repository

import (
	"encoding/json"
	"time"

	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"gorm.io/gorm"
)

type AddressRepository interface {
	FindByID(userID, id string) (*model.Address, error)
	FindAll(userID string) ([]*model.Address, error)
	Create(userID string, address *model.Address) error
	Update(userID string, address *model.Address) error
	Delete(userID, id string) error
	SetDefault(userID, id string) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (r *addressRepository) cacheGet(key string, dest interface{}) bool {
	val, err := cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), dest)
		return true
	}
	return false
}

func (r *addressRepository) cacheSet(key string, value interface{}, ttl time.Duration) {
	b, _ := json.Marshal(value)
	cache.Client.Set(cache.Ctx, key, b, ttl)
}

func (r *addressRepository) cacheDel(keys ...string) {
	cache.Client.Del(cache.Ctx, keys...)
}

func (r *addressRepository) FindByID(userID, id string) (*model.Address, error) {
	cacheKey := "address:id:" + userID + ":" + id
	var address model.Address

	if r.cacheGet(cacheKey, &address) {
		return &address, nil
	}

	if err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		First(&address).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, address, 10*time.Minute)

	return &address, nil
}

func (r *addressRepository) FindAll(userID string) ([]*model.Address, error) {
	cacheKey := "address:all:" + userID
	var addresses []*model.Address

	if r.cacheGet(cacheKey, &addresses) {
		return addresses, nil
	}

	if err := r.db.
		Where("user_id = ?", userID).
		Find(&addresses).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, addresses, 10*time.Minute)

	return addresses, nil
}

func (r *addressRepository) Create(userID string, address *model.Address) error {
	address.UserID = userID

	if err := r.db.Create(address).Error; err != nil {
		return err
	}

	r.cacheDel("address:all:" + userID)

	return nil
}

func (r *addressRepository) Update(userID string, address *model.Address) error {
	if err := r.db.
		Where("id = ? AND user_id = ?", address.ID, userID).
		Updates(address).Error; err != nil {
		return err
	}

	r.cacheDel(
		"address:all:"+userID,
		"address:id:"+userID+":"+address.ID,
	)

	return nil
}

func (r *addressRepository) Delete(userID, id string) error {
	var address model.Address

	tx := r.db.Begin()

	if err := tx.
		Where("id = ? AND user_id = ?", id, userID).
		First(&address).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.
		Delete(&model.Address{}, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if address.IsDefault {
		var another model.Address

		if err := tx.
			Where("user_id = ?", userID).
			First(&another).Error; err == nil {

			if err := tx.
				Model(&another).
				Update("is_default", true).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	r.cacheDel(
		"address:all:"+userID,
		"address:id:"+userID+":"+id,
	)

	return nil
}

func (r *addressRepository) SetDefault(userID, id string) error {
	tx := r.db.Begin()

	var address model.Address
	if err := tx.
		Where("id = ? AND user_id = ?", id, userID).
		First(&address).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.
		Model(&model.Address{}).
		Where("user_id = ?", userID).
		Update("is_default", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.
		Model(&model.Address{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_default", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	r.cacheDel(
		"address:all:"+userID,
		"address:id:"+userID+":"+id,
	)

	return nil
}
