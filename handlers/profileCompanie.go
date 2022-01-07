package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/Eli15x/MovieWorkNow/service"
)

func CreateProfileCompanie(c echo.Context) error {
	/*
	iduser
	,Nome,email,
	senha,
	profilemessage,
	aniversario,
	cargo1, cargo2,
	experiencia
	*/
	//o que será colocado na parte de cadastro do usuario serão o nome,email,senha e aniversario.

	name := c.Param("name")
	email := c.Param("email")
	password := c.Param("password")
	//birthDate, err := time.Parse("YYYY-MM-DD",c.Param("birthDate"))

	/*if err != nil {
		return c.String(403,"Create Profile Error: data format")
	}*/

	if name == "" {
		return c.String(403,"Create Profile Companie Error: name not find")
	}

	if email == "" {
		return c.String(403,"Create Profile Companie Error: email not find")
	}

	if password == "" {
		return c.String(403,"Create Profile Companie Error: password not find")
	}


	//ver qual o melhor caso, deixar no models como uma string a data inserir a data como string
	//no banco e quando for manipular o dado pegar e transformar em time.Time
	//ou já pegar o valor do context e transformar em data e salvar no banco como data.
	//validate data
	/*if birthDate == "" {
		return c.String(403,"Create Profile Error: birthDate not find")
	}*/

	err := service.GetInstanceProfileCompanie().CreateNewProfileCompanie(c,name,email,password)
	if err != nil{
		return c.String(403,"Create Profile Companie error: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

func AddInformationCompanie(c echo.Context) error {

	companieId := c.Param("companieId")
	job := c.Param("job")
	message := c.Param("message")


	if companieId == "" {
		return c.String(403,"Add information Companie Error: CompanieId not find")
	}

	if job == "" {
		return c.String(403,"Add information Companie Error: job not find")
	}

	if message == "" {
		return c.String(403,"Add information Companie Error: message not find")
	}

	err := service.GetInstanceProfileCompanie().AddInformationProfileCompanie(c,companieId,job,message)
	if err != nil{
		return c.String(403,"Add information Companie: error in service")
	}

	return c.String(http.StatusOK, "Ok")
}

