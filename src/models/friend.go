package models

type Friend struct {
	UserId           string    `json:"UserId,omitempty" bson:"UserId,omitempty"`
	FriendIds      	 []UserId   `json:"FriendIds,omitempty" bson:"FriendIds"`
	FriendRequests   []UserId   `json:"FriendRequests,omitempty" bson:"FriendRequests"`
}
