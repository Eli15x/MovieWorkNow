package models

type ProfileCompanie struct {
	CompanieId            string    `json:"companieId,omitempty" bson:"companieId,omitempty"`
	Name                  string    `json:"name,omitempty" bson:"-"`
	PassWord              string    `json:"password,omitempty" bson:"-"`
    Email                 string    `json:"email,omitempty" bson:"-"`
	Job                   string    `json:"job,omitempty" bson:"-"`
	ProfileMessage        string    `json:"profileMessage,omitempty" bson:"-"`
	//cargos esperados. um array tambem.
}
