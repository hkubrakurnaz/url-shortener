package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"url-shortener/config"
	url_controller "url-shortener/controller/url"
	url_repository "url-shortener/repository/url"
	encoder_service "url-shortener/service/encoder"
	url_service "url-shortener/service/url"
)

func init() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	config.ConnectDb(&configs)
}

func main() {
	app := fiber.New()
	urlRepo := url_repository.New()
	encoderService := encoder_service.New()
	urlService := url_service.New(encoderService, urlRepo)
	urlController := url_controller.New(urlService)
	urlController.Register(app)
	log.Fatal(app.Listen(":3000"))
}
