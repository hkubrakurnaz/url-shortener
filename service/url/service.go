package url

import (
	"time"
	"url-shortener/model"
	"url-shortener/repository/url"
	"url-shortener/service/base62"
)

type Service interface {
	CreateShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

type service struct {
	encoderService base62.Service
	urlRepository  url.Repository
}

func (s service) CreateShortUrl(longUrl string) (shortUrl string, err error) {
	urlModel := model.Url{
		LongUrl:   longUrl,
		CreatedAt: time.Time{},
	}

	id, err := s.urlRepository.Create(urlModel)
	if err != nil {
		return
	}
	shortUrl = s.encoderService.Encode(id)

	err = s.urlRepository.Update(id, shortUrl)

	return
}

func (s service) GetLongUrl(shortUrl string) (string, error) {
	id, err := s.encoderService.Decode(shortUrl)
	if err != nil {
		return "", err
	}

	response, err := s.urlRepository.Get(id)
	return response.LongUrl, err
}

func New(encoderService base62.Service, urlRepository url.Repository) Service {
	return &service{encoderService: encoderService, urlRepository: urlRepository}
}
