package url

import (
	"gorm.io/gorm"
	"url-shortener/model"
)

type Repository interface {
	Get(id uint64) (model.Url, error)
	Create(url model.Url) (uint64, error)
	Update(id uint64, shortUrl string) error
}
type repository struct {
	DB *gorm.DB
}

func (r repository) Create(url model.Url) (uint64, error) {
	var id model.Url
	err := r.DB.Create(&url).Scan(&id)

	return id.ID, err.Error
}

func (r repository) Update(id uint64, shortUrl string) error {
	err := r.DB.Model(&model.Url{}).Where("id = ?", id).Update("short_url", shortUrl)

	return err.Error
}

func (r repository) Get(id uint64) (model.Url, error) {
	var url model.Url
	err := r.DB.First(&url, "id= ?", id)
	return url, err.Error
}

func New(DB *gorm.DB) Repository {
	return &repository{DB: DB}
}
