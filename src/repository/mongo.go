package repository

import (
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)


type MongoDB interface {
	Find(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error)
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


