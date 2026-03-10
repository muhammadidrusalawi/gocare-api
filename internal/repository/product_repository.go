package repository

import (
	"encoding/json"
	"time"

	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(id string) (*model.Product, error)
	FindBySlug(slug string) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) cacheGet(key string, dest interface{}) bool {
	val, err := cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), dest)
		return true
	}
	return false
}

func (r *productRepository) cacheSet(key string, value interface{}, ttl time.Duration) {
	b, _ := json.Marshal(value)
	cache.Client.Set(cache.Ctx, key, b, ttl)
}

func (r *productRepository) cacheDel(keys ...string) {
	cache.Client.Del(cache.Ctx, keys...)
}

func (r *productRepository) FindByID(id string) (*model.Product, error) {
	cacheKey := "product:id:" + id
	var product model.Product

	if r.cacheGet(cacheKey, &product) {
		return &product, nil
	}

	if err := r.db.Where("id = ?", id).
		Preload("Category").
		First(&product).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, product, 10*time.Minute)
	return &product, nil
}

func (r *productRepository) FindBySlug(slug string) (*model.Product, error) {
	cacheKey := "product:slug:" + slug
	var product model.Product

	if r.cacheGet(cacheKey, &product) {
		return &product, nil
	}

	if err := r.db.Where("slug = ?", slug).First(&product).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, product, 10*time.Minute)
	return &product, nil
}

func (r *productRepository) FindAll() ([]*model.Product, error) {
	cacheKey := "product:all"
	var products []*model.Product

	if r.cacheGet(cacheKey, &products) {
		return products, nil
	}

	if err := r.db.
		Preload("Category").
		Find(&products).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, products, 10*time.Minute)
	return products, nil
}

func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	if err := r.db.
		Preload("Category").
		First(product, "id = ?", product.ID).Error; err != nil {
		return err
	}

	r.cacheDel("product:all")
	return nil
}

func (r *productRepository) Update(product *model.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}
	r.cacheDel(
		"product:all",
		"product:id:"+product.ID,
		"product:slug:"+product.Slug,
	)
	return nil
}

func (r *productRepository) Delete(id string) error {
	var product model.Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model.Product{}, "id = ?", id).Error; err != nil {
		return err
	}

	r.cacheDel(
		"product:all",
		"product:id:"+id,
		"product:slug:"+product.Slug,
	)
	return nil
}
