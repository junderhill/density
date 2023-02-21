package forecast

import "time"

type Forecast struct {
	Times []TimeWindow
}

type TimeWindow struct {
	Time                        time.Time
	TemperatureCelcius          float64
	FeelsLikeTemperatureCelcius float64
	LowCloudCoverPercent        int
	MedCloudCoverPercent        int
	HighCloudCoverPercent       int
}
