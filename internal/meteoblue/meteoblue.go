package meteoblue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/junderhill/density/internal/forecast"
	"github.com/junderhill/density/internal/location"
	"github.com/junderhill/density/internal/persist"
)

var apiKey string
var baseUrl string = "https://my.meteoblue.com/packages/basic-1h_clouds-1h_sunmoon"

func init() {
	apiKey = os.Getenv("METEOBLUE_API_KEY")
	if apiKey == "" {
		panic("METEOBLUE_API_KEY environment variable not set")
	}
}

func GetForecast(request *ForecastRequest) (*forecast.Forecast, error) {
	url := generateUrl(request.Location)

	fmt.Println("Contacting Meteoblue API...")

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		//todo: debug logging
		return nil, fmt.Errorf("meteoblue api returned status code %d", response.StatusCode)
	}

	fmt.Println("Parsing Meteoblue response...")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if request.PersistDir != "" {
		err := persist.PersistForecastResponse(request.PersistDir, body)
		if err != nil {
			//todo: write out err to stderr
			fmt.Printf("Failed to persist forecast response: %s", err)
		}
	}

	var meteoblueResponse meteoblueResponse
	err = json.Unmarshal(body, &meteoblueResponse)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", meteoblueResponse)

	return nil, nil
}

func generateUrl(location *location.Location) string {
	v := url.Values{}
	v.Set("apikey", apiKey)
	v.Set("lat", fmt.Sprint(location.Latitude))
	v.Set("lon", fmt.Sprint(location.Longitude))
	v.Set("format", "json")
	v.Set("forecast_days", "1")

	result := baseUrl + "?" + v.Encode() // put it all together
	return result
}

type meteoblueResponse struct {
	Metadata struct {
		Name                  string  `json:"name"`
		Latitude              float64 `json:"latitude"`
		Longitude             float64 `json:"longitude"`
		Height                int     `json:"height"`
		TimezoneAbbrevation   string  `json:"timezone_abbrevation"`
		UtcTimeoffset         float64 `json:"utc_timeoffset"`
		ModelrunUtc           string  `json:"modelrun_utc"`
		ModelrunUpdatetimeUtc string  `json:"modelrun_updatetime_utc"`
	} `json:"metadata"`
	Units struct {
		Time                     string `json:"time"`
		PrecipitationProbability string `json:"precipitation_probability"`
		Cloudcover               string `json:"cloudcover"`
		Sunshinetime             string `json:"sunshinetime"`
		Pressure                 string `json:"pressure"`
		Relativehumidity         string `json:"relativehumidity"`
		Visibility               string `json:"visibility"`
		Co                       string `json:"co"`
		Precipitation            string `json:"precipitation"`
		Temperature              string `json:"temperature"`
		Windspeed                string `json:"windspeed"`
		Winddirection            string `json:"winddirection"`
	} `json:"units"`
	Data1H struct {
		Time                     []string  `json:"time"`
		Sunshinetime             []int     `json:"sunshinetime"`
		Lowclouds                []int     `json:"lowclouds"`
		Midclouds                []int     `json:"midclouds"`
		Highclouds               []int     `json:"highclouds"`
		Visibility               []int     `json:"visibility"`
		Totalcloudcover          []int     `json:"totalcloudcover"`
		Precipitation            []float64 `json:"precipitation"`
		Snowfraction             []float64 `json:"snowfraction"`
		Rainspot                 []string  `json:"rainspot"`
		Temperature              []float64 `json:"temperature"`
		Felttemperature          []float64 `json:"felttemperature"`
		Pictocode                []int     `json:"pictocode"`
		Windspeed                []float64 `json:"windspeed"`
		Winddirection            []int     `json:"winddirection"`
		Relativehumidity         []int     `json:"relativehumidity"`
		Sealevelpressure         []float64 `json:"sealevelpressure"`
		PrecipitationProbability []int     `json:"precipitation_probability"`
		ConvectivePrecipitation  []float64 `json:"convective_precipitation"`
		Isdaylight               []int     `json:"isdaylight"`
		Uvindex                  []int     `json:"uvindex"`
	} `json:"data_1h"`
	DataDay struct {
		Time                 []string  `json:"time"`
		Sunrise              []string  `json:"sunrise"`
		Sunset               []string  `json:"sunset"`
		Moonrise             []string  `json:"moonrise"`
		Moonset              []string  `json:"moonset"`
		Moonphaseangle       []float64 `json:"moonphaseangle"`
		Moonage              []float64 `json:"moonage"`
		Moonphasename        []string  `json:"moonphasename"`
		Moonphasetransittime []string  `json:"moonphasetransittime"`
		Indexto1HvaluesStart []int     `json:"indexto1hvalues_start"`
		Indexto1HvaluesEnd   []int     `json:"indexto1hvalues_end"`
	} `json:"data_day"`
}
