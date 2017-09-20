package main

import(
	"fmt"
	"bufio"
	"os"
)

func receiveMosquittoData(mosquittoStreaming chan string){
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		mosquittoStreaming <- text
	}
}

func storeData(mosquittoStreaming chan string){
	for {
		fmt.Print("Stored: ", <- mosquittoStreaming)
	}
}

func main(){
	mosquittoStreaming := make(chan string)

	go storeData(mosquittoStreaming)
	receiveMosquittoData(mosquittoStreaming)
}
