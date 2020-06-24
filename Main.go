package main

import (
	"fmt"
	"github.com/ssimunic/gosensors"
	"time"
)

func main() {
	//_ = os.Mkdir("subdir", 0755) create folder
	initSensorsInfo()
}

func initSensorsInfo() {
	sensors, err := gosensors.NewFromSystem()

	if err != nil {
		panic(err)
	}

	for true {
		time.Sleep(time.Second * 2)
		fmt.Println(sensors)
		//for chip := range sensors.Chips {
			// Iterate over entries
			//for key, value := range sensors.Chips[chip] {
				// If CPU or GPU, print out
				//if key == "Core 0" || key == "GPU" {
				//	fmt.Println(key, value)
				//}
			//}
		//}
	}
}
