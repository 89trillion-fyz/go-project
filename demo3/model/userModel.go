package model

type User struct {
	Id         string            `json:"id"`
	ContentMap map[uint32]uint64 `json:"contentMap" bson:"contentMap"`
}
