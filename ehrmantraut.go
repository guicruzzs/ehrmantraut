package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("New data arriving:")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
