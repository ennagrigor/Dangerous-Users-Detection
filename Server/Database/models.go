package Database

import (
	"time"
)

type Filter struct {
	Size  int    `json:"size"`
	Query string `json:"query"`
}

type UserInfo struct {
	Username     string `json:"username"`
	Location     string `json:"location"`
	ProfileImage string `json:"profileImage"`
	Score        int    `json:"score"`
}

type UsersInfoList []*UserInfo

type UserTweet struct {
	UserID       string    `json:"userID"`
	Username     string    `json:"username"`
	Location     string    `json:"location"`
	ProfileImage string    `json:"profileImage"`
	Text         string    `json:"text"`
	Score        int       `json:"score"`
	Threat       string    `json:"threat"`
	Created      time.Time `json:"created"`
}

type UserTweetList []*UserTweet
