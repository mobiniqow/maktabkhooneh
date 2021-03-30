package model

type City struct {
	Name string `json:"name"`
	Id [] string `json:"id"`
}
// City constructor
func NewCity(name string,id []string) *City {
	return &City{Name: name,Id: id}
}