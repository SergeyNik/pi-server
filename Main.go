package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"github.com/ssimunic/gosensors"
	"net/http"
	"time"
)

func main() {
	task := gron.New()

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			task.AddFunc(gron.Every(2*time.Second), getSensorsInfo)
			task.Start()
			fmt.Fprintf(w, "Task started!")
		}
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			task.Stop()
			fmt.Fprintf(w, "Task stopped!")
		}
	})

	http.ListenAndServe(":3000", nil)
}

func getSensorsInfo() {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		panic(err)
	}
	//fmt.Println(sensors)
	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == "Core 0" || key == "GPU" {
				fmt.Println(key, value)
			}
		}
	}
}
