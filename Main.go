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
	//TODO: добавить вызов vcgencmd measure_temp для raspbrry pi
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		panic(err)
	}
	fmt.Println(sensors)
}
