package main

import (
	"net/http"
	"praktikum/praktikum/config"
	"praktikum/praktikum/features/handler"
	"praktikum/praktikum/features/service"
	"praktikum/praktikum/routes"
)

func main() {
	config := config.InitConfig()
	key := config.OpenAPIKey

	service := service.NewRecommendation(key)
	handler := handler.NewRecommendationHandler(service)

	route := routes.Routes(handler)

	server := http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}
