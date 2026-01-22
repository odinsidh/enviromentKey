package extractor

import (
	"enviromentkey/internal/dto"
)

type Extractor interface {
	Handle(requestPath string) ([]*dto.Location, error)
}
