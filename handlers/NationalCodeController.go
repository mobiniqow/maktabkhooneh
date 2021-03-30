package handlers

import (
	"github.com/gin-gonic/gin"
	"maktabkhoone/model"
	"maktabkhoone/service"
	"net/http"
)

// IsValidNationalCode godoc
// @Summary Get NationalCode Information by Id
// @Description Get Person Information by Name
// @Tags NationalCode
// @ID get-city-by-id
// @Accept  json
// @Produce  html
// @Param id path string true "NationalCode ID"
// @Header 200 {string} Token "qwerty"
// @Router /validation/national_code/{id} [get]

func  IsValidNationalCode(ctx *gin.Context) {
	cityService := service.GetCityService()
	input := ctx.Query("id")
	natinalCode  := model.NewNatinalCode(input)
	if !natinalCode.IsValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "شماره ملی صحیح نیست",
		})
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"result":*cityService.FindById(input),
		})
	}
}