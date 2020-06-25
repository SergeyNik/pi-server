package main

import (
	"fmt"
	//"github.com/roylee0704/gron"
	//"github.com/ssimunic/gosensors"
	//"net/http"
	//"time"
	"os/exec"
)

//
//func main() {
//	task := gron.New()
//	taskPi := gron.New()
//
//	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			task.AddFunc(gron.Every(2*time.Second), getSensorsInfoLm)
//			task.Start()
//			fmt.Fprintf(w, "Task started!")
//		}
//	})
//
//	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			task.Stop()
//			fmt.Fprintf(w, "Task stopped!")
//		}
//	})
//
//	http.HandleFunc("/pi/start", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			taskPi.AddFunc(gron.Every(2*time.Second), getSensorsInfoPi)
//			taskPi.Start()
//			fmt.Fprintf(w, "Task started!")
//		}
//	})
//
//	http.HandleFunc("/pi/stop", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			taskPi.Stop()
//			fmt.Fprintf(w, "Task stopped!")
//		}
//	})
//
//	http.HandleFunc("/pi/temp", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			fmt.Fprintf(w, currentTempPi())
//		}
//	})
//
//	http.ListenAndServe(":3000", nil)
//}
//
//func getSensorsInfoLm() {
//	sensors, err := gosensors.NewFromSystem()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(sensors)
//}
//
//func getSensorsInfoPi() {
//	command := exec.Command("vcgencmd", "measure_temp")
//	if output, error := command.Output(); error != nil {
//		fmt.Println("Error: ", error)
//	} else {
//		fmt.Println("CPU temp: ", output)
//	}
//}

func main() {
	currentTempPi()
}

func currentTempPi() {
	command := exec.Command("vcgencmd", "measure_temp")
	if output, error := command.Output(); error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println(string(output))
	}
}
