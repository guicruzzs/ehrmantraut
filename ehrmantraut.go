package main

import(
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)
// TODO: Data handling and DB should be in their places, not with the main program
type CarData struct {
    Fields map[string]interface{} `json:"fields"`
    Tags map[string]interface{} `json:"tags"`
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

// TODO: In parse error, it shouldn't store data
func storeData(mosquittoStreaming chan []byte){
	for {
		jsonData := <- mosquittoStreaming
		carData := parseJSON(jsonData)
		fmt.Printf("Stored: %s", carData)
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
