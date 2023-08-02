package url

import (
	"time"
	"url-shortener/model"
	"url-shortener/repository/url"
	"url-shortener/service/encoder"
)

type Service interface {
	CreateShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

type service struct {
	encoderService encoder.Service
	urlRepository  url.Repository
}

func (s service) CreateShortUrl(longUrl string) (string, error) {
	urlModel := model.Url{
		LongUrl:   longUrl,
		CreatedAt: time.Time{},
	}

	id, err := s.urlRepository.Create(urlModel)
	shortUrl := s.encoderService.Encode(id)

	err = s.urlRepository.Update(id, shortUrl)

	return shortUrl, err
}

func (s service) GetLongUrl(shortUrl string) (string, error) {
	id, _ := s.encoderService.Decode(shortUrl)
	response, err := s.urlRepository.GetLongUrl(id)

	return response.LongUrl, err
}

func New(encoderService encoder.Service, urlRepository url.Repository) Service {
	return &service{encoderService: encoderService, urlRepository: urlRepository}
}
