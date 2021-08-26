package TextDetection

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode"
	"userdetection/Configuration"
	"userdetection/Database"
	"userdetection/SocialMedia/Twitter"
)

var languageTranslationMap = map[string]string{
	"iw": "hebrew",
	"en": "english",
}

type ThreatClassification string

const (
	NoRisk     ThreatClassification = "NoRisk"
	LowRisk    ThreatClassification = "LowRisk"
	MediumRisk ThreatClassification = "MediumRisk"
	HighRisk   ThreatClassification = "HighRisk"
)

type TextSeverity struct {
	Score  int                  `json:"score"`
	Threat ThreatClassification `json:"threat"`
}

func translateLanguageShortcut(languageShortcut string) (string, error) {
	languageResult, ok := languageTranslationMap[languageShortcut]
	if !ok {
		return "", fmt.Errorf("language was not detected")
	}

	return languageResult, nil
}

func classifyText(text, lang string) (*TextSeverity, error) {
	language, err := translateLanguageShortcut(lang)
	if err != nil {
		return nil, err
	}

	dictionaryPath := filepath.Join(Configuration.Conf.DictionaryFolder, fmt.Sprintf("%s.json", language))
	b, err := ioutil.ReadFile(dictionaryPath)
	if err != nil {
		return nil, err
	}

	var dictionary map[string]int
	err = json.Unmarshal(b, &dictionary)
	if err != nil {
		return nil, err
	}

	splitFunc := func(input rune) bool {
		return !unicode.IsLetter(input)
	}

	words := strings.FieldsFunc(strings.Trim(strings.ToLower(text), ""), splitFunc)
	severity := 0

	for i := 0; i < len(words); i++ {
		for j := i + 1; j <= len(words); j++ {
			currentWord := strings.Join(words[i:j], " ")
			if currentSeverity, ok := dictionary[currentWord]; ok {
				severity += currentSeverity
			}
		}
	}

	severityResult := &TextSeverity{
		Score: severity,
	}

	if severity == 0 {
		severityResult.Threat = NoRisk
	} else if severity > 0 && severity < 5 {
		severityResult.Threat = LowRisk
	} else if severity >= 5 && severity < 10 {
		severityResult.Threat = MediumRisk
	} else {
		severityResult.Threat = HighRisk
	}

	return severityResult, nil
}

func DetectDangerousUserTweet(_ context.Context) {
	lockups, err := Twitter.ScanTweets()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Fetched %d tweets from twitter\n", len(lockups.LookUps))

	userTweets := make(Database.UserTweetList, 0, lockups.Meta.ResultCount)
	for _, lockup := range lockups.LookUps {
		textSeverity, err := classifyText(lockup.Tweet.Text, lockup.Tweet.Language)
		if err != nil {
			log.Println(err)
			return
		}

		if textSeverity.Threat != NoRisk { // found tweet with threat
			fmt.Println("Detected dangerous tweet with id: ", lockup.Tweet.ID)
			userInfo, err := Twitter.FindUserByID(lockup.Tweet.AuthorID)
			if err != nil {
				log.Println(err)
				continue
			}
			userTweets = append(userTweets, &Database.UserTweet{
				UserID:       userInfo.ID,
				Username:     userInfo.UserName,
				Location:     userInfo.Location,
				ProfileImage: userInfo.ProfileImageURL,
				Text:         lockup.Tweet.Text,
				Score:        textSeverity.Score,
				Threat:       string(textSeverity.Threat),
			})
		}
	}

	err = Database.SaveUsersTweets(userTweets)
	if err != nil {
		log.Println("Failed saving dangerous tweets: ", err)
	}
}
