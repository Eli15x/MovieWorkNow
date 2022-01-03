package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/Eli15x/MovieWorkNow/service"
)

func CreateProfile(c echo.Context) error {
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
	birthDate := c.Param("birthDate")

	if name == "" {
		return c.String(403,"Create Profile Error: name not find")
	}

	if email == "" {
		return c.String(403,"Create Profile Error: email not find")
	}

	if password == "" {
		return c.String(403,"Create Profile Error: password not find")
	}

	if birthDate == "" {
		return c.String(403,"Create Profile Error: birthDate not find")
	}

	err := service.createNewProfile(c,name,email,password,birthDate)
	if err != nil {
		return c.String(403,"Create New Profile Error: Error in try to create new Profile")
	}

	
	return c.String(http.StatusOK, "Ok")
}
