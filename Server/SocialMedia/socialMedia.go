package SocialMedia

// SocialMedia ...
type SocialMedia interface {
	ScanUsers(filter string) UsersInfo
	GetUserInfo(uid string) *UserInfo
	GetUserFeed(uid string, feedID string) *FeedInfo
	GetUserGroups(uid string) *GroupsInfo
}
