package Facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	sm "userdetection/SocialMedia"

	fb "github.com/huandu/facebook/v2"
)

const (
	accessTokenBaseURL = "https://graph.facebook.com/oauth/access_token"
)

type FacebookSocialMedia struct {
	AccessToken string
}

func InitFacebookClient() *FacebookSocialMedia {
	accessToken := "EAAmnXIsJsqkBABywZAGUYGZBZAEyF7noYCw5MUZAZCO6Sh0rWTPNZANVcEocJ3rEUUJ6VIuBjef10Sk1NeVg67U2xElVNDbbe8BofZByNHsBz9F1I10yPZBCEaSVnquKZAYOUydGyTa76ih9sVNBQzk3KU5WvKauq5g7XW73wwMtd9wZDZD"
	return &FacebookSocialMedia{AccessToken: accessToken}
}

func authenticate() *sm.AccessTokenResponse {
	client := http.Client{}
	request, err := http.NewRequest("GET", accessTokenBaseURL, nil)
	if err != nil {
		//todo change
		fmt.Println("Error occured ", err)
		return nil
	}
	q := request.URL.Query()
	q.Add("client_id", "2717290701894313")
	q.Add("client_secret", "ee3cd665b2e417e9b12a91aaf88de577")
	q.Add("grant_type", "client_credentials")
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		//todo change
		fmt.Println("Error occured ", err)
		return nil
	}
	defer response.Body.Close()
	var accessTokenResult sm.AccessTokenResponse
	json.NewDecoder(response.Body).Decode(&accessTokenResult)

	return &accessTokenResult
}

func (facebook *FacebookSocialMedia) ScanUsers(filter string) sm.UsersInfo {
	res, err := fb.Get("/search", fb.Params{
		"access_token": facebook.AccessToken,
		"q":            filter,
	})
	if err != nil {
		return nil
	}

	var items []fb.Result
	err = res.DecodeField("data", &items)
	if err != nil {
		return nil
	}

	for _, entry := range items {
		fmt.Println(entry)
	}
	return nil
}

func (facebook *FacebookSocialMedia) GetUserInfo(uid string) *sm.UserInfo {
	res, err := fb.Get(fmt.Sprintf("/%s", uid), fb.Params{
		"access_token": facebook.AccessToken,
		"fields":       "first_name",
	})
	if err != nil {
		return nil
	}

	fmt.Println(res)
	return nil
}

func (facebook *FacebookSocialMedia) GetUserFeed(uid string, feedID string) *sm.FeedInfo { return nil }
func (facebook *FacebookSocialMedia) GetUserGroups(uid string) *sm.GroupsInfo            { return nil }
