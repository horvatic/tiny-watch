package main

import (
    "net/http"
    "os/exec"
    "fmt"
)

func main() {
    http.HandleFunc("/", HealthService)
    http.ListenAndServe(":5001", nil)
}

func HealthService(w http.ResponseWriter, r *http.Request) {
    out, err := exec.Command("/bin/sh", "health.sh").Output()

    if err != nil {
	    fmt.Fprintf(w, "Can not get system info")
    }

    fmt.Fprintf(w, "%v", string(out[:]))
}
