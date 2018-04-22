package main

import (
	"fmt"
)

func main() {
	var (
		x = 12
		y = "気温"
		z = 22.4
	)
	fmt.Println(newWeatherReport(x, y, z))
}

// newWeatherReport generates the weather information from formatter.
func newWeatherReport(time interface{}, str interface{}, temp interface{}) string {
	return fmt.Sprintf("%d時の%sは%.1f", time, str, temp)
}

// Report represents format for templating
const Report = `{{.Time}}時の{{.Text}}は{{.Temperature}}`

// Weather represents the weather information
type Weather struct {
	Text        string
	Time        int
	Temperature float64
}

// NewWeather creates new instance of weather information for reporting
func NewWeather(time int, text string, temperature float64) *Weather {
	return &Weather{Text: text, Time: time, Temperature: temperature}
}
