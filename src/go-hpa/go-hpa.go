package main

import (
    "fmt"
    "math"
    "log"
    "net/http"
)

func LoopHole(bolded string) string {
    x := 0.000001
    for i := 0; i <= 10000000; i++ {
        x += math.Sqrt(x)
    }
	return fmt.Sprintf("<b>%s</b>", bolded)
}

func handler(writer http.ResponseWriter, reader *http.Request) {
    fmt.Fprintf(writer, LoopHole("Code.Education Rocks!"))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":80", nil))
}