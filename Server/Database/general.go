package Database

import (
	"github.com/blevesearch/bleve/v2"
	"os"
	"userdetection/Configuration"
)

const (
	IndexName = "index.bleve"
)

// OpenedIndex ...
var OpenedIndex bleve.Index

// InitIndex ...
func InitIndex() error {
	if _, indexNotExists := os.Stat(Configuration.Conf.IndexPath); os.IsNotExist(indexNotExists) {
		mapping := bleve.NewIndexMapping()
		bleveIndex, err := bleve.New(Configuration.Conf.IndexPath, mapping)
		if err != nil {
			return err
		}
		OpenedIndex = bleveIndex
	} else {
		bleveIndex, err := bleve.Open(Configuration.Conf.IndexPath)
		if err != nil {
			return err
		}
		OpenedIndex = bleveIndex
	}
	return nil
}
