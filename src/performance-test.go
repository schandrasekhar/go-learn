package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "time"
    "os"
)



type PerformanceStat struct {
    success bool
    latency time.Duration
    statuscode  string
}


func api_get_call(c chan PerformanceStat) {
    url := "https://graviton-ncp-content-gateway.media.yahoo.com/api/v1/gql/stream_view?namespace=news&id=newsroom-trending&version=v1"
    success := false
    start := time.Now()
    response, err := http.Get(url)
    if err == nil {
        success = true
    }
    end := time.Now()
    _, err = ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }
    

    if err != nil {
        log.Fatalln(err)
    }
    latency := end.Sub(start)
    stat := PerformanceStat{success: success, latency: latency, statuscode: "statuscode"}
    c <- stat
}

func api_get_call1() {
    url := "https://www.google.com"
    success := false
    start := time.Now()
    response, err := http.Get(url)
    end := time.Now()
    if err == nil {
        success = true
    }
    

    _, err = ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }
    

    if err != nil {
        log.Fatalln(err)
    }
    latency := end.Sub(start)
    stat := PerformanceStat{success: success, latency: latency, statuscode: "statuscode"}
    log.Fatalln(stat.latency)
}

func write_to_file(c chan PerformanceStat) {
    f, _ := os.OpenFile("./results.txt", os.O_APPEND|os.O_WRONLY, 0644)
    for i := range c {
        _, _ = f.WriteString(i.latency.String() + "\n")
    }
    close(c)
    f.Close()
}

func main() {
    api_get_call1()
    // count := 1
    // c := make(chan PerformanceStat, count)
    // for i := 0; i < count; i++ {
    //     go api_get_call(c)
    // }
    // write_to_file(c)
}