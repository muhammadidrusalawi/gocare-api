package helper

import (
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

func GenerateUniqueSlug(db *gorm.DB, model interface{}, name string) (string, error) {
	baseSlug := normalizeSlug(name)
	slug := baseSlug
	counter := 1

	for {
		var count int64

		err := db.Model(model).
			Where("slug = ?", slug).
			Count(&count).Error

		if err != nil {
			return "", err
		}

		if count == 0 {
			break
		}

		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	return slug, nil
}

func normalizeSlug(name string) string {
	slug := strings.ToLower(name)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}
