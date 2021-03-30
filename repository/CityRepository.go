package repository

import "maktabkhoone/model"
// cuz some time multi repositories
type CityRepository interface {
 	FinCityById(id string) *model.City
}



