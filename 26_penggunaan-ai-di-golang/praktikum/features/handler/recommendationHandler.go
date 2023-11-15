package handler

import (
	"fmt"
	"net/http"
	"praktikum/praktikum/features/data"
	"praktikum/praktikum/features/service"
	"praktikum/praktikum/helper"

	"github.com/julienschmidt/httprouter"
)

type RecommendationHandler interface {
	GetRecommendation(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type RecommendationModel struct {
	service service.RecommendationInterface
}

func NewRecommendationHandler(service service.RecommendationInterface) RecommendationHandler {
	return &RecommendationModel{service: service}
}

func (handler *RecommendationModel) GetRecommendation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	recommendationRequest := data.Request{}

	if request == nil {
		helper.SetResponse(writer, data.Response{
			Status: "No Recommendations",
			Data:   nil,
		})
	}

	recommendationMessage := "i need a laptop for " + recommendationRequest.Purpose + "purpose, what's the best laptops i could buy with budget of : Rp." + fmt.Sprintf("%d", recommendationRequest.Budget)
	res, err := handler.service.GetRecommendation(recommendationMessage)
	if err != nil {
		helper.SetResponse(writer, data.Response{
			Status: err.Error(),
			Data:   nil,
		})
	}

	helper.SetResponse(writer, data.Response{
		Status: "success",
		Data:   res,
	})
}
