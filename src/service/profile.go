package service

import (
	"sync"
	"fmt"
	"time"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/Eli15x/MovieWorkNow/src/repository"
	"github.com/Eli15x/MovieWorkNow/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/fatih/structs"
)

var (
	instanceProfile CommandProfile
	onceProfile     sync.Once
)

type CommandProfile interface {
	CreateNewProfile(ctx echo.Context, name string,email string,password string) error
	AddInformationProfile(ctx echo.Context,id string,job []string, message string) error
	GetInformationProfile(ctx echo.Context,id string) ([]bson.M, error)
	CheckInformationValid(ctx echo.Context,id string) error
	AddRelationFriendProfile(ctx echo.Context,UserId_user string,UserId_value string, friend *models.Friend) error
	AddRequestFriend(ctx echo.Context,UserId string,FriendId string, friendUser *models.Friend) error
	DeleteFriendRequest(ctx echo.Context,UserId string, FriendId string, friendUser *models.Friend) error
	AddContent(ctx echo.Context,id string,content string) error
}

type profile struct{}

func GetInstanceProfile() CommandProfile {
	onceProfile.Do(func() {
		instanceProfile = &profile{}
	})
	return instanceProfile
}

func (p *profile)CreateNewProfile(ctx echo.Context,name string, email string, password string) error {
	
	var userId = utils.CreateCodeId()
	profile := &models.Profile {
		UserId: userId,
		Name : name,
		Email: email,
		PassWord: password,
	}

	profileInsert := structs.Map(profile)
	
	_, err := storage.GetInstance().Insert(ctx,"profile",profileInsert)
	if err != nil {
		return errors.New("Create New Profile: problem to insert into MongoDB")
	}

	err = CreateFriendTable(ctx,userId)
	if err != nil {
		return err
	}

	return  nil
}

func (p *profile)AddInformationProfile(ctx echo.Context,id string,job []string, message string) error {
	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile",userId)
	if mgoErr != nil {
		return errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	profileUpdate := map[string]interface{} {
		"Job": job,
		"ProfileMessage":message,
	}

	fmt.Println(profileUpdate)

	change := bson.M{"$set": profileUpdate}

	_, err := storage.GetInstance().UpdateOne(ctx,"profile",userId,change)
	if err != nil {
		return errors.New("Create New Profile: problem to update into MongoDB")
	}

	return  nil
}

func (p *profile)GetInformationProfile(ctx echo.Context,id string) ([]bson.M, error){
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}

	result, err := repository.Find(ctx, "profile",userId, &profile)
	if err != nil {
		return nil, errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (p *profile)CheckInformationValid(ctx echo.Context,email string, password string) error {

	userInfo := map[string]interface{}{"Email": email , "Password": password}


	result, err := repository.Find(ctx, "profile",userInfo, &profile)
	if err != nil {
		return nil, errors.New("Check Information: problem to Find user login into MongoDB")
	}

	return nil
}

func (p *profile)AddRelationFriendProfile(ctx echo.Context,UserId string,FriendId string, friendUser *models.Friend) error{

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId",userId) 
	
    err := result.Decode(friendUser)
    if err != nil {
		fmt.Println(err)
		return errors.New("Error Decode Friend") 
    }

	var UsersIds []models.UserId
	for _, friend := range friendUser.FriendIds {
		if friend.UserId == FriendId {
			return errors.New("Error User is already friend")
		}
		newUserId := models.UserId{
			UserId: friend.UserId,
		}
		UsersIds = append(UsersIds,newUserId)
	}

	newUserId := models.UserId{
		UserId: FriendId,
	}
	UsersIds = append(UsersIds, newUserId)


	newFriend := &models.Friend {
		UserId: UserId,
		FriendIds: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	err = p.DeleteFriendRequest(ctx,UserId,FriendId,friendUser)
	if err != nil {
		return err
	}

	_, err = storage.GetInstance().UpdateOne(ctx,"friend",userId,change)
	if err != nil {
		return errors.New("Add Friend Relation: problem to update into MongoDB")
	}

	return nil
}

func (p *profile)AddRequestFriend(ctx echo.Context,UserId string,FriendId string, friendUser *models.Friend) error{

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId",userId) 
	
    err := result.Decode(friendUser)
    if err != nil {
		fmt.Println(err)
		return errors.New("Error Decode Friend") 
    }

	var UsersIds []models.UserId
	for _, friend := range friendUser.FriendRequests {
		if friend.UserId == FriendId {
			return errors.New("Error User already have a request")
		}
		newUserId := models.UserId{
			UserId: friend.UserId,
		}
		UsersIds = append(UsersIds,newUserId)
	}

	newUserId := models.UserId{
		UserId: FriendId,
	}
	UsersIds = append(UsersIds, newUserId)


	newFriend := &models.Friend {
		UserId: UserId,
		FriendRequests: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	_, err = storage.GetInstance().UpdateOne(ctx,"friend",userId,change)
	if err != nil {
		return errors.New("Add Friend Relation: problem to update into MongoDB")
	}

	return nil
}

func (p *profile)DeleteFriendRequest(ctx echo.Context,UserId string, FriendId string, friendUser *models.Friend) error {

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId",userId) 
	
    err := result.Decode(friendUser)
    if err != nil {
		fmt.Println(err)
		return errors.New("Error Decode Friend") 
    }

	var UsersIds []models.UserId
	for _, friend := range friendUser.FriendRequests {
		if friend.UserId != FriendId {
		  newUserId := models.UserId{
	      UserId: friend.UserId,
		}
		  UsersIds = append(UsersIds,newUserId)
		}
	}

	newFriend := &models.Friend {
		UserId: UserId,
		FriendRequests: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	_, err = storage.GetInstance().UpdateOne(ctx,"friend",userId,change)
	if err != nil {
		return errors.New("Delete Requestion Relation: problem to update into MongoDB")
	}

	return nil
}

func CreateFriendTable(ctx echo.Context,userId string) error {

	var UsersIds []models.UserId

	friend := &models.Friend {
		UserId: userId,
		FriendIds : UsersIds,
	}

	FriendInsert := structs.Map(friend)
	

	_, err := storage.GetInstance().Insert(ctx,"friend",FriendInsert)
	if err != nil {
		return errors.New("Create Friend Table: problem to insert into MongoDB")
	}

	return  nil
}

func (p *profile)AddContent(ctx echo.Context,id string,content string) error {
	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile",userId)
	if mgoErr != nil {
		return errors.New("Add Content: problem to Find Id into MongoDB")
	}

	newContent := &models.Content {
		ContentId: utils.CreateCodeId(),
		UserId : id,
		Content: content,
		Data: time.Now(),
	}

	newContentInsert := structs.Map(newContent)
	

	_, err := storage.GetInstance().Insert(ctx,"content",newContentInsert)
	if err != nil {
		return errors.New("Add Content: problem to update into MongoDB")
	}

	return  nil
}


