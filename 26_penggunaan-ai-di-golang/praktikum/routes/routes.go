package routes

import (
	"praktikum/praktikum/features/handler"

	"github.com/julienschmidt/httprouter"
)

func Routes(recommendationHandler handler.RecommendationHandler) *httprouter.Router {
	route := httprouter.New()
	route.POST("/recommendations", recommendationHandler.GetRecommendation)
	return route
}
