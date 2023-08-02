package url

import (
	"url-shortener/config"
	"url-shortener/model"
)

type Repository interface {
	GetLongUrl(id uint64) (model.Url, error)
	Create(url model.Url) (uint64, error)
	Update(id uint64, shortUrl string) error
}
type repository struct {
}

func (r repository) Create(url model.Url) (uint64, error) {
	var id model.Url
	err := config.DB.Create(&url).Scan(&id)

	return id.ID, err.Error
}

func (r repository) Update(id uint64, shortUrl string) error {
	err := config.DB.Model(&model.Url{}).Where("id = ?", id).Update("short_url", shortUrl)

	return err.Error
}

func (r repository) GetLongUrl(id uint64) (model.Url, error) {
	var url model.Url
	err := config.DB.First(&url, "id= ?", id)
	return url, err.Error
}

func New() Repository {
	return &repository{}
}
