package forecast

import "time"

type Forecast struct {
	Times []TimeWindow
}

type TimeWindow struct {
	Time time.Time
}
