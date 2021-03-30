package service

import (
	"maktabkhoone/model"
	"maktabkhoone/repository"
)

type CityService struct {
	repository repository.CityRepository
}
// initialization instance
func GetCityService() *CityService  {
	instance := &CityService{}
	instance.repository = repository.GetRepository()
	return instance
}
//query for find instance
func (c * CityService) FindById(id string) *model.City {
	return c.repository.FinCityById(id)
}
