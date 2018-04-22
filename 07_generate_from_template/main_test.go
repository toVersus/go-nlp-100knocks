package main

import (
	"bytes"
	"html/template"
	"reflect"
	"testing"
)

var templatingTests = []struct {
	name string
	time int
	str  string
	temp float64
	want string
}{
	{
		name: "should generate a weather report from template",
		time: 12,
		str:  "気温",
		temp: 22.4,
		want: "12時の気温は22.4",
	},
}

func TestSimpleTemplateReporting(t *testing.T) {
	for _, testcase := range templatingTests {
		t.Log(testcase.name)

		result := newWeatherReport(testcase.time, testcase.str, testcase.temp)
		if !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.want)
		}
	}
}

func TestTemplateReporting(t *testing.T) {
	for _, testcase := range templatingTests {
		t.Log(testcase.name)

		weather := NewWeather(testcase.time, testcase.str, testcase.temp)
		tmpl := template.Must(template.New("Report").Parse(Report))

		var buf bytes.Buffer
		err := tmpl.Execute(&buf, weather)
		if err != nil {
			t.Error(err)
		}

		if result := buf.String(); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.want)
		}
	}
}
