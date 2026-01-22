package extraxtordefault

import (
	"enviromentkey/internal/dto"
	"io"
	"os"
	"regexp"
	"strconv"
)

func (self *extractor) Handle(requestPath string) ([]*dto.Location, error) {
	byteContainer, err := self.filePreparation(requestPath)
	if err != nil {
		return nil, err
	}

	loc, err := self.regexByte(byteContainer)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (self *extractor) filePreparation(requestPath string) ([]byte, error) {
	// openfile
	file, err := os.OpenFile(requestPath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read all
	byteContainer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return byteContainer, nil
}

func (self *extractor) regexByte(input []byte) ([]*dto.Location, error) {
	var output []*dto.Location = []*dto.Location{}
	regexInstance := regexp.MustCompile(`(?m)(^.*?)(?:\s-.*?)(\d{2}\.\d*)(?:[\D]*)(\d{2}\.\d*)`)
	regexContainer := regexInstance.FindAllStringSubmatch(string(input), -1)
	for _, value := range regexContainer {
		// extract
		name := value[1]

		// cast
		x, err := strconv.ParseFloat(value[2], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(value[3], 64)
		if err != nil {
			return nil, err
		}

		// append
		output = append(output, &dto.Location{
			Name: name,
			X:    x,
			Y:    y,
		})
	}
	return output, nil
}
