package Configuration

import (
	"encoding/json"
	"io/ioutil"
)

var Conf Configuration

// Conf ...
type Configuration struct {
	LanguageDetectionAPI string `json:"LanguageDetectionAPI"`
	DictionaryFolder     string `json:"DictionaryFolder"`
}

func InitConfiguration(configurationPath string) (err error) {
	b, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &Conf)
	if err != nil {
		return
	}

	return
}
