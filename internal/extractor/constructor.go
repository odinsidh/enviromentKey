package extractor

import (
	"enviromentkey/internal/extractor/extraxtordefault"
	"errors"
)

var (
	implementationDoesNotExist = errors.New("implementation does not exist")
)

func GetExtractor(target string) (Extractor, error) {
	if r, ok := implementation[target]; ok {
		return r, nil
	} else {
		return nil, implementationDoesNotExist
	}
}

var implementation map[string]Extractor = map[string]Extractor{
	"default": extraxtordefault.GetExtractor(),
}
