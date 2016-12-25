package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os/exec"
    )

func hello(w http.ResponseWriter, r *http.Request) {
    log.Fatal("I AM ALIVE")
	io.WriteString(w, "Hello world!")
}

func main() {
    fmt.Println("Uptime Monitor")
    out, err := exec.Command("uptime").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", out)

    //mux := http.NewServeMux()
    http.HandleFunc("/", hello)
	for i := 0; i < 10; i-- {
    	http.ListenAndServe(":80", nil)
	}
}
