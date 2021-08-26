package Configuration

import (
	"encoding/json"
	"io/ioutil"
)

// Conf ...
var Conf Configuration

// Configuration configuration file representation
type Configuration struct {
	TwitterToken      string `json:"TwitterToken"`
	DictionaryFolder  string `json:"dictionaryFolder"`
	ApplicationFolder string `json:"ApplicationFolder"`
	ClientFolder      string `json:"ClientFolder"`
	IndexPath         string `json:"IndexPath"`
}

// InitConfiguration init configuration struct
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
