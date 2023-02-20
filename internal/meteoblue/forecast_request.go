package meteoblue

import "github.com/junderhill/density/internal/location"

type ForecastRequest struct {
	Location   *location.Location
	PersistDir string
}
