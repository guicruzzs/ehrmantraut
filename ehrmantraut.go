package main

import(
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

// TODO: Move these configs to another place, like env variables
const (
	MyDB = "ehrmantraut"
	username = "ehrmantraut"
	password = "test123"
)

// TODO: Data handling and DB should be in their places, not with the main program
type CarData struct {
    Fields map[string]interface{} `json:"fields"`
    Tags map[string]string `json:"tags"`
    Time int64 `json:"time"`
}

// TODO: In error case: it should properly be logged and the error should be returned
func parseJSON(jsonData []byte) CarData {
	var data CarData
	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println("Parse error")
	}
	return data
}

// TODO: Move To a single responsability package
func save(data CarData) {
	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://db:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Println(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		// TODO: Check this precision
		Precision: "s",
	})
	if err != nil {
		fmt.Println(err)
	}

	// TODO: Use the Car Data instead time.Now()
	// TODO: Study the app case and understand needs of measurements to database
	pt, err := client.NewPoint("phantom_car", data.Tags, data.Fields, time.Now())
	if err != nil {
		fmt.Println(err)
	}
	bp.AddPoint(pt)

	if err := influxClient.Write(bp); err != nil {
		fmt.Println(err)
	}
}

// TODO: In parse error, it shouldn't store data
func storeData(mosquittoStreaming chan []byte){
	for {
		jsonData := <- mosquittoStreaming
		carData := parseJSON(jsonData)
		fmt.Printf("Storing: %v\n", carData)
		save(carData)
	}
}

func receiveMosquittoData(mosquittoStreaming chan []byte){
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadBytes('\n')
		mosquittoStreaming <- text
	}
}

func main(){
	mosquittoStreaming := make(chan []byte)

	go storeData(mosquittoStreaming)
	receiveMosquittoData(mosquittoStreaming)
}
