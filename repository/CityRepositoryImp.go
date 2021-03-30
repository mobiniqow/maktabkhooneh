package repository

import (
	"maktabkhoone/model"
)

type CityRepositoryImp struct {
	cities []model.City
}

//test repository
func GetRepository() *CityRepositoryImp {
	instance := &CityRepositoryImp{}
	instance.cities = []model.City{
		*model.NewCity("بروجرد", []string{"412", "413"}),
		*model.NewCity("خرم اباد", []string{"406", "407"}),
		*model.NewCity("دورود", []string{"421"}),
		*model.NewCity("بابلسر", []string{"498"}),
		*model.NewCity("چالوس", []string{"483"}),
		*model.NewCity("ساری", []string{"208"}),
		*model.NewCity("رامسر", []string{"227"}),
		*model.NewCity("کلاردشت", []string{"627"})}
	return instance
}

func (c *CityRepositoryImp) FinCityById(id string) *model.City {
	for _, element := range c.cities {
		for _, inner := range element.Id {
			if inner == id[0:3] {
				return &element
			}
		}
	}
	return model.NewCity("تعریف نشده", []string{""})
}
