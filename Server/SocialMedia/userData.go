package SocialMedia

type UserInfo struct {
	Username string
	FullName string
	Age      int
	Location string
	Gender   string
	Married  bool
}

type UsersInfo []*UserInfo

type FeedInfo struct {
	FeedID        string
	Text          string
	NumberOfLikes int
}

//todo
type GroupInfo struct {
}

//todo
type GroupsInfo []GroupInfo

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
