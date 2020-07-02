package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = fmt.Fprintln(w, getSensorsInfoPi())
		}
	})

	_ = http.ListenAndServe(":3000", nil)
}

func getSensorsInfoPi() string {
	command := exec.Command("vcgencmd", "measure_temp")
	output, err := command.Output()
	if err != nil {
		panic(err)
	}
	return string(output)
}
