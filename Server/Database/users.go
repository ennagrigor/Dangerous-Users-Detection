package Database

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/google/uuid"
	"log"
	"math"
	"sort"
	"time"
)

func SaveUsersTweets(usersTweets UserTweetList) error {
	if len(usersTweets) == 0 {
		return nil
	}
	batch := OpenedIndex.NewBatch()
	for _, userTweet := range usersTweets {
		userTweet.Created = time.Now()
		err := batch.Index(uuid.New().String(), userTweet)
		if err != nil {
			return err
		}
	}

	return OpenedIndex.Batch(batch)
}

func GetTweets(filter *Filter) (UserTweetList, error) {
	query := bleve.NewQueryStringQuery(filter.Query)
	searchRequest := &bleve.SearchRequest{
		Query: query,
		Size:  filter.Size,
	}
	searchRequest.Fields = []string{"*"}

	searchResult, err := OpenedIndex.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	userTwitsResult := make(UserTweetList, 0, len(searchResult.Hits))

	for _, hit := range searchResult.Hits {
		createdTime, err := time.Parse(time.RFC3339Nano, hit.Fields["created"].(string))
		if err != nil {
			log.Printf("Failed parsing time of user twit with id: %s", hit.ID)
		}
		userTwitsResult = append(userTwitsResult, &UserTweet{
			UserID:       hit.Fields["userID"].(string),
			Username:     hit.Fields["username"].(string),
			Location:     hit.Fields["location"].(string),
			ProfileImage: hit.Fields["profileImage"].(string),
			Text:         hit.Fields["text"].(string),
			Score:        int(hit.Fields["score"].(float64)),
			Threat:       hit.Fields["threat"].(string),
			Created:      createdTime,
		})
	}

	return userTwitsResult, nil
}

func GetTopDangerousUsers() (UsersInfoList, error) {
	searchRequest := &bleve.SearchRequest{
		Query: bleve.NewMatchAllQuery(),
		Size:  math.MaxInt64,
	}
	searchRequest.Fields = []string{"*"}

	searchResult, err := OpenedIndex.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	usersMapInfo := make(map[string]*UserInfo, 0)
	for _, hit := range searchResult.Hits {
		userID := hit.Fields["userID"].(string)
		if info, ok := usersMapInfo[userID]; !ok {
			usersMapInfo[userID] = &UserInfo{
				Username:     hit.Fields["username"].(string),
				Location:     hit.Fields["location"].(string),
				ProfileImage: hit.Fields["profileImage"].(string),
				Score:        int(hit.Fields["score"].(float64)),
			}
		} else {
			info.Score += int(hit.Fields["score"].(float64))
			usersMapInfo[userID] = info
		}
	}

	userInfoResult := make(UsersInfoList, 0, len(usersMapInfo))
	for _, info := range usersMapInfo {
		userInfoResult = append(userInfoResult, info)
	}

	sort.Slice(userInfoResult, func(i, j int) bool {
		return userInfoResult[i].Score > userInfoResult[j].Score
	})

	if len(userInfoResult) > 5 {
		userInfoResult = userInfoResult[:5]
	}

	return userInfoResult, nil
}
