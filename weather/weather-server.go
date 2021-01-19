package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "net/http"
  "html/template"
  "encoding/json"
)


type Weather struct {
	Current struct {
		Cloud     int64 `json:"cloud"`
		Condition struct {
			Code int64  `json:"code"`
			Icon string `json:"icon"`
			Text string `json:"text"`
		} `json:"condition"`
		FeelslikeC       int64   `json:"feelslike_c"`
		FeelslikeF       int64   `json:"feelslike_f"`
		GustKph          int64   `json:"gust_kph"`
		GustMph          float64 `json:"gust_mph"`
		Humidity         int64   `json:"humidity"`
		IsDay            int64   `json:"is_day"`
		LastUpdated      string  `json:"last_updated"`
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		PrecipIn         int64   `json:"precip_in"`
		PrecipMm         int64   `json:"precip_mm"`
		PressureIn       float64 `json:"pressure_in"`
		PressureMb       int64   `json:"pressure_mb"`
		TempC            int64   `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		Uv               int64   `json:"uv"`
		VisKm            int64   `json:"vis_km"`
		VisMiles         int64   `json:"vis_miles"`
		WindDegree       int64   `json:"wind_degree"`
		WindDir          string  `json:"wind_dir"`
		WindKph          float64 `json:"wind_kph"`
		WindMph          float64 `json:"wind_mph"`
	} `json:"current"`
	Location struct {
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Localtime      string  `json:"localtime"`
		LocaltimeEpoch int64   `json:"localtime_epoch"`
		Lon            float64 `json:"lon"`
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		TzID           string  `json:"tz_id"`
	} `json:"location"`
}

var weatherUrl = "http://api.weatherapi.com/v1/current.json?key=314e77b40bba445ebdf111850211901&q="


func home(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    t := template.Must(template.ParseFiles("index.html"))
    t.Execute(res, nil)
  }

  if req.Method == "POST" {
    place := req.FormValue("place")
    url := weatherUrl

    fmt.Println("Place: ", place)
    url += strings.Title(place)
    fmt.Println("url: ", url)

    weatherRes, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
    }
    var weatherData Weather
    data, _ := ioutil.ReadAll(weatherRes.Body)

    // convert to struct
    json.Unmarshal([]byte(data), &weatherData)
    fmt.Printf("%v\n", weatherData)

    place = ""
    url = ""

    t := template.Must(template.ParseFiles("weather.html"))
    t.Execute(res, weatherData)
  }
}

func main() {
  fmt.Println("Server is running...")
  http.HandleFunc("/", home)
  //http.HandleFunc("/weather-server", weatherHandler)
  http.ListenAndServe(":8000", nil)
}
