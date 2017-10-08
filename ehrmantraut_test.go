package main

import(
	"testing"
)

func TestParseJSON(t *testing.T) {
	rawJSON := []byte(`{ "fields" : { "engine_temperature": 850.4, "rpm" : 7504, "throttle" : 0.75, "maf" : 0.67, "gear" : 4, "speed" : 250.6 }, "tags" : { "car_id" : 10001, "lap" : 3  }, "time" : 1496510681952374020}`)
	carData := parseJSON(rawJSON)

	if(carData.Fields == nil) {
		t.Error("Parse error")
	}

	if(carData.Tags == nil) {
		t.Error("Parse error")
	}

	if(carData.Time == 0) {
		t.Error("Parse error")
	}
}


func TestParseJSONInvalid(t *testing.T) {
	rawJSON := []byte(`{}`)
	carData := parseJSON(rawJSON)

	if(carData.Fields == nil) {
		t.Error("Parse error")
	}

	if(carData.Tags == nil) {
		t.Error("Parse error")
	}

	if(carData.Time == 0) {
		t.Error("Parse error")
	}
}