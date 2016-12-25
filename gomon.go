// https://golang.org/doc/articles/wiki/
package main

import (
    "encoding/json"
    "html/template"
    "net/http"
    "os/exec"
    "log"
    "strconv"
    "strings"
    "time"
)

type Sample struct {
    Timestamp int
    Load int
}

type ReportResponse struct {
    Timestamp int
    Samples[] Sample
}

var maxSamples int = 32 //6 * 10 * 10
var buffer = make([]Sample, maxSamples)

var ns int = 0
var ls int = 0
var numSamples = 0

var quit = make(chan struct{})

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func handler(w http.ResponseWriter, r *http.Request) {
    cmd := r.URL.Path[1:]
    if (cmd == "quit") {
        close(quit)
    } else if (cmd == "ui") {
        t, _ := template.ParseFiles("templates/ui.html")
        t.Execute(w, nil)
    } else {
        w.Header().Set("Content-Type", "application/json")
        
        now := time.Now()
        curstamp := now.Minute() * 60 + now.Second()
        timestamp := now.Hour() * 3600 + curstamp
    
        // TODO improve memory use of RelativeSample buffer
        x := make([]Sample, numSamples)
        ix := 0
    
        // buffer
        limit := min(maxSamples, (ls + numSamples))
        for i := ls; i < limit; i++ {
            x[ix] = buffer[i]
            ix++
        }
        for i := 0; i < ls; i++ {
            x[ix] = buffer[i]
            ix++
        }
    
        // New JSON stuff
        m := ReportResponse{timestamp, x}
        json.NewEncoder(w).Encode(m)
    }
}

// TODO get rid of Fatals (i.e. recover)
func sampleLoad() int {
    out, err := exec.Command("uptime").Output()
    if err != nil {
        log.Fatal(err)
    }

    n := len(out)
    s := string(out[:n])
    ss := strings.Split(s, " ")
    loadString := strings.TrimRight(ss[13], ",")
    loadFloat, err := strconv.ParseFloat(loadString, 32)
    if err != nil {
        log.Print("%v", ss)
        log.Fatal(err)
    }
    load := int(loadFloat * 100)
    return load
}

func takeSample() {
    numSamples++

    load := sampleLoad()
    
    now := time.Now()
    timestamp := now.Minute() * 60 + now.Second()

    buffer[ns] = Sample{Timestamp: timestamp, Load: load}

    // manage leading index (Next Sample)
    ns++
    ns = ns % maxSamples

    // manage trailing index (Last Sample)
    if (numSamples > maxSamples) {
        numOver := numSamples - maxSamples
        ls = ls + numOver
        ls = ls % maxSamples
        numSamples = numSamples - numOver
    }
}

func main() {
    takeSample()

    // http://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
    ticker := time.NewTicker(10 * time.Second)
    go func() {
        for {
            select {
                case <- ticker.C:
                    takeSample()

                case <- quit:
                    ticker.Stop()
                    return
            }
        }
    }()

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
