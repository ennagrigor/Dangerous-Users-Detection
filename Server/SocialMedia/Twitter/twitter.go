package Twitter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/g8rswimmer/go-twitter"
	"net/http"
	"strings"
	"userdetection/Configuration"
)

const (
	twitterEndpoint = "https://api.twitter.com"
)

var targetChars = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func parseTwitterError(tweetErr *twitter.TweetErrorResponse) error {
	enc, err := json.MarshalIndent(tweetErr, "", "    ")
	if err != nil {
		return err
	}
	return fmt.Errorf(string(enc))
}

func ScanTweets() (*twitter.TweetRecentSearch, error) {
	tweet := &twitter.Tweet{
		Authorizer: authorize{
			Token: Configuration.Conf.TwitterToken,
		},
		Client: http.DefaultClient,
		Host:   twitterEndpoint,
	}

	fieldOpts := twitter.TweetFieldOptions{TweetFields: []twitter.TweetField{twitter.TweetFieldAuthorID, twitter.TweetFieldGeo, twitter.TweetFieldLanguage}}
	searchOpts := twitter.TweetRecentSearchOptions{MaxResult: 100}
	targetChars := strings.Join(targetChars, " OR ")
	query := fmt.Sprintf("(%v) lang:en", targetChars)

	recentSearch, err := tweet.RecentSearch(context.Background(), query, searchOpts, fieldOpts)
	var tweetErr *twitter.TweetErrorResponse
	switch {
	case errors.As(err, &tweetErr):
		return nil, parseTwitterError(tweetErr)
	case err != nil:
		return nil, err
	}

	return recentSearch, nil
}

func FindUserByID(userID string) (*twitter.UserObj, error) {
	user := &twitter.User{
		Authorizer: authorize{
			Token: Configuration.Conf.TwitterToken,
		},
		Client: http.DefaultClient,
		Host:   twitterEndpoint,
	}
	fieldOpts := twitter.UserFieldOptions{
		UserFields: []twitter.UserField{twitter.UserFieldProfileImageURL, twitter.UserFieldLocation},
	}

	lookups, err := user.Lookup(context.Background(), []string{userID}, fieldOpts)
	var tweetErr *twitter.TweetErrorResponse
	switch {
	case errors.As(err, &tweetErr):
		return nil, parseTwitterError(tweetErr)
	case err != nil:
		return nil, err
	}

	if len(lookups) == 0 {
		return nil, fmt.Errorf("twitter user info not found")
	}
	userInfo := lookups[userID].User

	return &userInfo, nil
}
