package repository

import (
	"encoding/json"
	"time"

	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindByID(id string) (*model.Category, error)
	FindByName(name string) (*model.Category, error)
	FindBySlug(slug string) (*model.Category, error)
	FindAll() ([]*model.Category, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) cacheGet(key string, dest interface{}) bool {
	val, err := cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), dest)
		return true
	}
	return false
}

func (r *categoryRepository) cacheSet(key string, value interface{}, ttl time.Duration) {
	b, _ := json.Marshal(value)
	cache.Client.Set(cache.Ctx, key, b, ttl)
}

func (r *categoryRepository) cacheDel(keys ...string) {
	cache.Client.Del(cache.Ctx, keys...)
}

func (r *categoryRepository) FindByID(id string) (*model.Category, error) {
	cacheKey := "category:id:" + id
	var category model.Category

	if r.cacheGet(cacheKey, &category) {
		return &category, nil
	}

	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, category, 10*time.Minute)
	return &category, nil
}

func (r *categoryRepository) FindByName(name string) (*model.Category, error) {
	cacheKey := "category:name:" + name
	var category model.Category

	if r.cacheGet(cacheKey, &category) {
		return &category, nil
	}

	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, category, 10*time.Minute)
	return &category, nil
}

func (r *categoryRepository) FindBySlug(slug string) (*model.Category, error) {
	cacheKey := "category:slug:" + slug
	var category model.Category

	if r.cacheGet(cacheKey, &category) {
		return &category, nil
	}

	if err := r.db.Where("slug = ?", slug).First(&category).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, category, 10*time.Minute)
	return &category, nil
}

func (r *categoryRepository) FindAll() ([]*model.Category, error) {
	cacheKey := "category:all"
	var categories []*model.Category

	if r.cacheGet(cacheKey, &categories) {
		return categories, nil
	}

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	r.cacheSet(cacheKey, categories, 10*time.Minute)
	return categories, nil
}

func (r *categoryRepository) Create(category *model.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	r.cacheDel("category:all")
	return nil
}

func (r *categoryRepository) Update(category *model.Category) error {
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	r.cacheDel(
		"category:all",
		"category:id:"+category.ID,
		"category:slug:"+category.Slug,
	)
	return nil
}

func (r *categoryRepository) Delete(id string) error {
	var category model.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model.Category{}, "id = ?", id).Error; err != nil {
		return err
	}

	r.cacheDel(
		"category:all",
		"category:id:"+id,
		"category:slug:"+category.Slug,
		"category:name:"+category.Name,
	)
	return nil
}
