package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/rizkiamr/x-response-time/httptimer"
)

func main() {
	h := http.NewServeMux()
	h.HandleFunc("/ruok", index)

	log.Fatal(http.ListenAndServe(":8080", httptimer.Timed(h)))
}

func index(w http.ResponseWriter, r *http.Request) {
	wait, err := time.ParseDuration(r.URL.Query().Get("wait"))
	if err != nil {
		wait = 2333 * time.Nanosecond
	}
	time.Sleep(wait)

	fmt.Printf("Requested from %s: %s %s\n",
		r.UserAgent(),
		r.Method,
		r.RequestURI)

	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "%v: %v\n", r.Method, r.URL.Path)
	//fmt.Fprintln(w, "Check X-Response-Time for the processing time of this request.")

	resp := make(map[string]string)
	resp["message"] = "ok"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
