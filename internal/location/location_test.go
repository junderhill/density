package location

import (
	"fmt"
	"testing"
)

func TestLocationValidation(t *testing.T) {
	t.Run("Valid Location", func(t *testing.T) {
		lat, long := "52.78868", "-1.66946"

		location, err := NewLocation(lat, long)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if fmt.Sprint(location.Latitude) != lat {
			t.Errorf("Expected %v, got %v", lat, location.Latitude)
		}

		if fmt.Sprint(location.Longitude) != long {
			t.Errorf("Expected %v, got %v", long, location.Longitude)
		}
	})

	t.Run("Invalid Latitude", func(t *testing.T) {
		lat, long := "100.78868", "-1.66946"

		_, err := NewLocation(lat, long)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Invalid long", func(t *testing.T) {
		lat, long := "52.78868", "200.66946"

		_, err := NewLocation(lat, long)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Invalid string for lat", func(t *testing.T) {
		lat, long := "abc", "-1.66946"

		_, err := NewLocation(lat, long)
		if err.Error() != "latitude or longitude is missing or not a valid number" {
			t.Errorf("Expected specific error message, got %v", err)
		}
	})
}
