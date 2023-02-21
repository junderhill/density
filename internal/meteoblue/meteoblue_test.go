package meteoblue

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ApiKey = "test"
	m.Run()
	os.Exit(0)
}

func TestConvertToForecast(t *testing.T) {
	var mateoblueResponse meteoblueResponse
	err := json.Unmarshal([]byte(sampleMeteoBlueResponse), &mateoblueResponse)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	t.Run("Output times matches data point count", func(t *testing.T) {
		forecast := ConvertToForecast(&mateoblueResponse)
		if len(forecast.Times) != len(mateoblueResponse.Data1H.Time) {
			t.Errorf("Expected %v, got %v", len(mateoblueResponse.Data1H.Time), len(forecast.Times))
		}
	})

	t.Run("Each time window has a temperature", func(t *testing.T) {
		forecast := ConvertToForecast(&mateoblueResponse)
		for _, timeWindow := range forecast.Times {
			if timeWindow.TemperatureCelcius != 0.0 {
				t.Errorf("Expected temperature to be set")
			}
		}
	})

	t.Run("Each time window has a feels like temperature", func(t *testing.T) {
		forecast := ConvertToForecast(&mateoblueResponse)
		for _, timeWindow := range forecast.Times {
			if timeWindow.FeelsLikeTemperatureCelcius != 0.0 {
				t.Errorf("Expected feels like temperature to be set")
			}
		}
	})

	t.Run("Each time window has a low cloud cover percentage", func(t *testing.T) {
		forecast := ConvertToForecast(&mateoblueResponse)
		for _, timeWindow := range forecast.Times {
			if timeWindow.LowCloudCoverPercent != 0 {
				t.Errorf("Expected low cloud cover percentage to be set")
			}
		}
	})

	t.Run("Each time window has a high cloud cover percentage", func(t *testing.T) {
		forecast := ConvertToForecast(&mateoblueResponse)
		for _, timeWindow := range forecast.Times {
			if timeWindow.HighCloudCoverPercent != 0 {
				t.Errorf("Expected high cloud cover percentage to be set")
			}
		}
	})
}

var sampleMeteoBlueResponse string = `
{
	"metadata": 
	{
		"name": "", 
		"latitude": 52.79, 
		"longitude": -1.67, 
		"height": 46, 
		"timezone_abbrevation": "GMT", 
		"utc_timeoffset": 0.00, 
		"modelrun_utc": "2023-02-18 12:00", 
		"modelrun_updatetime_utc": "2023-02-18 20:07"
	}, 
	"units": 
	{
		"time": "YYYY-MM-DD hh:mm", 
		"precipitation_probability": "percent", 
		"cloudcover": "percent", 
		"sunshinetime": "minutes", 
		"pressure": "hPa", 
		"relativehumidity": "percent", 
		"visibility": "m", 
		"co": "ug/m3", 
		"precipitation": "mm", 
		"temperature": "C", 
		"windspeed": "ms-1", 
		"winddirection": "degree"
	}, 
	"data_1h": 
	{
		"time": ["2023-02-18 00:00", "2023-02-18 01:00", "2023-02-18 02:00", "2023-02-18 03:00", "2023-02-18 04:00", "2023-02-18 05:00", "2023-02-18 06:00", "2023-02-18 07:00", "2023-02-18 08:00", "2023-02-18 09:00", "2023-02-18 10:00", "2023-02-18 11:00", "2023-02-18 12:00", "2023-02-18 13:00", "2023-02-18 14:00", "2023-02-18 15:00", "2023-02-18 16:00", "2023-02-18 17:00", "2023-02-18 18:00", "2023-02-18 19:00", "2023-02-18 20:00", "2023-02-18 21:00", "2023-02-18 22:00", "2023-02-18 23:00", "2023-02-19 00:00"], 
		"sunshinetime": [null, 0, 0, 0, 0, 0, 0, 0, 16, 3, 3, 42, 19, 3, 3, 3, 3, 22, 9, 0, 0, 0, 0, 0, 0], 
		"lowclouds": [75, 60, 75, 45, 35, 0, 40, 60, 90, 90, 29, 65, 90, 90, 90, 90, 60, 60, 90, 45, 31, 60, 45, 62, 66], 
		"midclouds": [62, 60, 75, 45, 35, 0, 4, 0, 0, 0, 20, 62, 85, 81, 61, 48, 55, 60, 81, 45, 0, 0, 0, 0, 0], 
		"highclouds": [52, 48, 42, 35, 25, 14, 6, 0, 0, 0, 30, 71, 99, 97, 89, 76, 52, 24, 9, 23, 51, 69, 64, 49, 37], 
		"visibility": [10040, 10540, 10850, 10500, 10050, 10050, 10010, 10010, 10010, 9410, 10810, 13010, 14610, 16460, 17010, 17260, 16410, 15010, 13250, 11450, 10050, 10050, 10050, 10650, 10450], 
		"totalcloudcover": [75, 60, 75, 45, 35, 4, 40, 60, 90, 90, 29, 65, 90, 90, 90, 90, 60, 60, 90, 45, 31, 60, 45, 62, 66], 
		"precipitation": [0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00], 
		"snowfraction": [0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00], 
		"rainspot": ["0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000099909019999909911999", "0000000000000000000000000009000009000900990000999", "0000000000000000000000000090000000009900900999900", "0000000000000000000000000000000000000000000000009", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000090000009000000", "0000000000000000000000000000000000000000000000000", "0000000000000900000990000099000000900000000000000", "0000000000000000000000000000000000000000000000001", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000009990090000", "0000000000000000000000000000000000000000090000009", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000090000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000"], 
		"temperature": [11.67, 11.30, 10.93, 10.74, 10.56, 10.56, 10.56, 10.65, 10.92, 11.40, 12.01, 12.64, 13.03, 13.20, 13.19, 12.82, 12.26, 11.58, 10.99, 10.60, 10.30, 10.03, 9.85, 9.73, 9.56], 
		"felttemperature": [6.80, 5.75, 5.29, 5.27, 5.22, 5.37, 5.30, 5.18, 5.56, 6.37, 6.95, 7.25, 7.59, 8.19, 8.03, 7.64, 7.18, 6.72, 6.23, 5.87, 5.23, 5.35, 5.96, 5.44, 5.57], 
		"pictocode": [8, 8, 8, 8, 5, 14, 4, 7, 19, 19, 5, 8, 21, 21, 21, 20, 8, 8, 19, 8, 14, 8, 8, 8, 8], 
		"windspeed": [6.63, 7.41, 7.36, 7.08, 6.89, 6.77, 6.90, 7.22, 7.05, 6.56, 6.54, 6.87, 6.84, 6.16, 6.39, 6.36, 6.18, 5.85, 5.73, 5.78, 6.35, 5.82, 4.58, 4.96, 4.40], 
		"winddirection": [236, 241, 244, 243, 246, 247, 247, 255, 258, 259, 263, 268, 270, 269, 266, 267, 266, 261, 257, 257, 258, 264, 267, 273, 278], 
		"relativehumidity": [84, 84, 83, 83, 84, 86, 86, 86, 84, 81, 77, 72, 69, 67, 68, 68, 71, 74, 77, 80, 83, 85, 84, 82, 81], 
		"sealevelpressure": [1018.73, 1018.48, 1018.11, 1017.75, 1017.36, 1016.96, 1016.78, 1016.95, 1017.32, 1017.68, 1018.00, 1018.30, 1018.44, 1018.27, 1017.94, 1017.70, 1017.65, 1017.70, 1017.83, 1018.03, 1018.32, 1018.70, 1019.29, 1019.98, 1020.48], 
		"precipitation_probability": [9, 3, 0, 4, 4, 5, 5, 5, 5, 0, 0, 0, 0, 0, 0, 1, 1, 1, 5, 1, 1, 1, 1, 0, 0], 
		"convective_precipitation": [0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00], 
		"isdaylight": [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0], 
		"uvindex": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	}, 
	"data_day": 
	{
		"time": ["2023-02-18", "2023-02-19"], 
		"sunrise": ["07:16", "07:14"], 
		"sunset": ["17:25", "17:27"], 
		"moonrise": ["07:03", "07:33"], 
		"moonset": ["14:26", "16:04"], 
		"moonphaseangle": [333.98, 347.67], 
		"moonage": [27.40, 28.52], 
		"moonphasename": ["waning crescent", "waning crescent"], 
		"moonphasetransittime": ["---", "---"], 
		"indexto1hvalues_start": [0, 24], 
		"indexto1hvalues_end": [23, 24]
	}
}
	`
