package repository

import (
	"fmt"
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)


type MongoDB interface {
	Find(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error)
	FindFriend(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) (models.Friend, error)
}


func Find(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error){
 
	cursor, err := storage.GetInstance().Find(ctx, collName, query , doc)
 	if err != nil {
	 return nil,ctx.String(403,"Error Repository: Error find query in mongoDb")
 	}

 	var content []bson.M
 	if err = cursor.All(ctx.Request().Context(), &content); err != nil {
		return nil,ctx.String(403,"Error Repository: Error Get Cursor information mongoDB")
 	}

 	return content, nil
}

func FindFriend(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) (models.Friend, error){
 
	cursor, err := storage.GetInstance().Find(ctx, collName, query , doc)
 	if err != nil {
	 return models.Friend{},ctx.String(403,"Error Repository: Error find query in mongoDb")
 	}

 	var content models.Friend
 	if err = cursor.All(nil, content); err != nil {
		 fmt.Println(err)
		return models.Friend{},ctx.String(403,"Error Repository: Error Get Cursor information mongoDB")
 	}

	 // a ideia era com que essa função retornasse para mim  o retorno que seria um userId_user e uma string de ids que seriam o userId.
	 // essa ideia tem como questão o agrupamento de ids.

 	return content, nil
}


