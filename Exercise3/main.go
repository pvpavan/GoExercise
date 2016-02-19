package main

import (
	"fmt"
	"net/http"
	"strings"
)

var cloneRedis map[string]string

func handler(w http.ResponseWriter, r *http.Request) {

	request_key := r.URL.Path[1:]
	method := r.Method

	if method == "GET" {
		fmt.Fprintf(w, "%s", cloneRedis[request_key])
	}

	if method == "PUT" {
		values := strings.Split(request_key, "/")
		cloneRedis[values[0]] = values[1]
		fmt.Fprint(w, "request_key: "+values[0]+" value: "+values[1]+" saved.")
	}

	if method == "DELETE" {
		delete(cloneRedis, request_key)
		fmt.Fprint(w, "request_key: "+request_key+" deleted.")
	}
	if method == "COUNT" {
		count := 0
		if request_key == "" {
			fmt.Fprint(w, len(cloneRedis))
		} else {
			for key, _ := range cloneRedis {
				if strings.HasPrefix(key, request_key) {
					count++
				}
			}
			fmt.Fprint(w, count)
		}

	}

}

func main() {
	cloneRedis = make(map[string]string)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
