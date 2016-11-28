package main

import (
	"./jsonReader"
	"fmt"
)

func main() {
	config := jsonReader.ReadJson("config.json")
	fmt.Println(config.CmdDelay)
}
