package TextDetection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"userdetection/Configuration"

	"github.com/detectlanguage/detectlanguage-go"
)

var languageTranslationMap = map[string]string{
	"iw": "hebrew",
	"en": "english",
	"ar": "arabic",
}

func detectLanguage(text string) (string, error) {
	client := detectlanguage.New(Configuration.Conf.LanguageDetectionAPI)
	detections, err := client.Detect(text)

	if err != nil {
		return "", err
	}

	languageShortcut := detections[0].Language
	languageResult, ok := languageTranslationMap[languageShortcut]
	if !ok {
		return "", fmt.Errorf("Language was not detected")
	}

	return languageResult, nil
}

/*
Twitter -> search on specific user -> got twitt text ->

"Hello, I'm going to bomb something! Tired of Jews"
*/

func ClassifyText(text string) error {
	language, err := detectLanguage(text)
	if err != nil {
		return err
	}

	dictionaryPath := filepath.Join(Configuration.Conf.DictionaryFolder, fmt.Sprintf("%s.json", language))
	b, err := ioutil.ReadFile(dictionaryPath)
	if err != nil {
		return err
	}

	var dictionary map[string]bool
	err = json.Unmarshal(b, &dictionary)
	if err != nil {
		return err
	}

	return nil
}
