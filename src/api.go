package main;

import "net/http";
import "io/ioutil";
import "fmt";

func main() {
    resp, err := http.Get("http://www.example.com");
    if err != nil {
        fmt.Printf("error occurred in api call");
    }
    defer resp.Body.Close();
    body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        fmt.Printf("error occurred in body read");
    }
    str := string(body[:]);
    fmt.Printf(str);
}