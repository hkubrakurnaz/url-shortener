package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"url-shortener/config"
	url_controller "url-shortener/controller/url"
	url_repository "url-shortener/repository/url"
	base62_service "url-shortener/service/base62"
	url_service "url-shortener/service/url"
)

var (
	urlRepository url_repository.Repository
	base62Service base62_service.Service
	urlService    url_service.Service
	db            *gorm.DB
	app           *fiber.App
)

func init() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	db = config.ConnectDb(&configs)
}

func main() {
	log.Println("Application started!")

	app = fiber.New()
	LoadDependency()

	log.Fatal(app.Listen(":3000"))
}

func LoadDependency() {
	urlRepository = url_repository.New(db)
	base62Service = base62_service.New()
	urlService = url_service.New(base62Service, urlRepository)
	url_controller.New(app, urlService)
}
