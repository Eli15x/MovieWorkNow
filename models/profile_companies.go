package models

type ProfileCompanie struct {
	CompanieId            string    `json:"companieId,omitempty" bson:"companieId,omitempty"`
	Name                  string    `json:"name,omitempty" bson:"-"`
	JobsId               string    `json:"cargoId,omitempty" bson:"-"`
	//cargos esperados. um array tambem.
}
