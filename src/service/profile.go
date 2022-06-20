package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/Eli15x/MovieWorkNow/src/repository"
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"github.com/Eli15x/MovieWorkNow/utils"
	"github.com/fatih/structs"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceProfile CommandProfile
	onceProfile     sync.Once
)

type CommandProfile interface {
	CreateNewProfile(ctx context.Context, name string, email string, password string) error
	AddInformationProfile(ctx context.Context, id string, job []string, message string) error
	GetInformationProfile(ctx context.Context, id string) ([]bson.M, error)
	CheckInformationValid(ctx context.Context, email string, password string, profile *models.Profile) (string, error)
	AddRelationFriendProfile(ctx context.Context, UserId_user string, UserId_value string, friend *models.Friend) error
	AddRequestFriend(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error
	DeleteFriendRequest(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error
	AddContent(ctx context.Context, id string, content string) error
}

type profile struct{}

func GetInstanceProfile() CommandProfile {
	onceProfile.Do(func() {
		instanceProfile = &profile{}
	})
	return instanceProfile
}

func (p *profile) CreateNewProfile(ctx context.Context, name string, email string, password string) error {

	var userId = utils.CreateCodeId()
	profile := &models.Profile{
		UserId:   userId,
		Name:     name,
		Email:    email,
		PassWord: password,
	}

	profileInsert := structs.Map(profile)

	_, err := storage.GetInstance().Insert(ctx, "profile", profileInsert)
	if err != nil {
		return errors.New("Create New Profile: problem to insert into MongoDB")
	}

	err = CreateFriendTable(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (p *profile) AddInformationProfile(ctx context.Context, id string, job []string, message string) error {
	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile", userId)
	if mgoErr != nil {
		return errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	profileUpdate := map[string]interface{}{
		"Job":            job,
		"ProfileMessage": message,
	}

	fmt.Println(profileUpdate)

	change := bson.M{"$set": profileUpdate}

	_, err := storage.GetInstance().UpdateOne(ctx, "profile", userId, change)
	if err != nil {
		return errors.New("Create New Profile: problem to update into MongoDB")
	}

	return nil
}

func (p *profile) GetInformationProfile(ctx context.Context, id string) ([]bson.M, error) {
	var profile models.Profile

	userId := map[string]interface{}{"UserId": id}

	result, err := repository.Find(ctx, "profile", userId, &profile)
	if err != nil {
		return nil, errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (p *profile) CheckInformationValid(ctx context.Context, email string, password string, profile *models.Profile) (string, error) {

	filter := map[string]interface{}{"Email": email, "PassWord": password}
	result := storage.GetInstance().FindOne(ctx, "profile", filter)

	if result == nil {
		return "", errors.New("Check Information: user not find")
	}

	err := result.Decode(profile)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error Decode Profile")
	}

	log.Infof("[CheckInformationValid] Object : %s \n", profile, "")

	return profile.UserId, nil
}

func (p *profile) AddRelationFriendProfile(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error {

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId", userId)

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
		UsersIds = append(UsersIds, newUserId)
	}

	newUserId := models.UserId{
		UserId: FriendId,
	}
	UsersIds = append(UsersIds, newUserId)

	newFriend := &models.Friend{
		UserId:    UserId,
		FriendIds: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	err = p.DeleteFriendRequest(ctx, UserId, FriendId, friendUser)
	if err != nil {
		return err
	}

	_, err = storage.GetInstance().UpdateOne(ctx, "friend", userId, change)
	if err != nil {
		return errors.New("Add Friend Relation: problem to update into MongoDB")
	}

	return nil
}

func (p *profile) AddRequestFriend(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error {

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId", userId)

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
		UsersIds = append(UsersIds, newUserId)
	}

	newUserId := models.UserId{
		UserId: FriendId,
	}
	UsersIds = append(UsersIds, newUserId)

	newFriend := &models.Friend{
		UserId:         UserId,
		FriendRequests: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	_, err = storage.GetInstance().UpdateOne(ctx, "friend", userId, change)
	if err != nil {
		return errors.New("Add Friend Relation: problem to update into MongoDB")
	}

	return nil
}

func (p *profile) DeleteFriendRequest(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error {

	userId := map[string]interface{}{"UserId": UserId}
	result := storage.GetInstance().FindOne(ctx, "FriendId", userId)

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
			UsersIds = append(UsersIds, newUserId)
		}
	}

	newFriend := &models.Friend{
		UserId:         UserId,
		FriendRequests: UsersIds,
	}

	FriendUpdate := structs.Map(newFriend)

	change := bson.M{"$set": FriendUpdate}

	_, err = storage.GetInstance().UpdateOne(ctx, "friend", userId, change)
	if err != nil {
		return errors.New("Delete Requestion Relation: problem to update into MongoDB")
	}

	return nil
}

func CreateFriendTable(ctx context.Context, userId string) error {

	var UsersIds []models.UserId

	friend := &models.Friend{
		UserId:    userId,
		FriendIds: UsersIds,
	}

	FriendInsert := structs.Map(friend)

	_, err := storage.GetInstance().Insert(ctx, "friend", FriendInsert)
	if err != nil {
		return errors.New("Create Friend Table: problem to insert into MongoDB")
	}

	return nil
}

func (p *profile) AddContent(ctx context.Context, id string, content string) error {
	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile", userId)
	if mgoErr != nil {
		return errors.New("Add Content: problem to Find Id into MongoDB")
	}

	newContent := &models.Content{
		ContentId: utils.CreateCodeId(),
		UserId:    id,
		Content:   content,
		Data:      time.Now(),
	}

	newContentInsert := structs.Map(newContent)

	_, err := storage.GetInstance().Insert(ctx, "content", newContentInsert)
	if err != nil {
		return errors.New("Add Content: problem to update into MongoDB")
	}

	return nil
}
