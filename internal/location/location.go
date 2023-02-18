package location

import (
	"errors"
	"strconv"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

var ErrInvalidLongLat = errors.New("invalid latitude or longitude (must be between -90 and 90 for latitude and -180 and 180 for longitude)")
var ErrInvalidLongLatFormat = errors.New("latitude or longitude is missing or not a valid number")

func NewLocation(lat, long string) (*Location, error) {
	fLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, ErrInvalidLongLatFormat
	}
	fLong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return nil, ErrInvalidLongLatFormat
	}

	if fLat < -90 || fLat > 90 {
		return nil, ErrInvalidLongLat
	}

	if fLong < -180 || fLong > 180 {
		return nil, ErrInvalidLongLat
	}

	return &Location{
		Latitude:  fLat,
		Longitude: fLong,
	}, nil
}
