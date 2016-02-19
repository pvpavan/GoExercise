package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler (w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if(err!=nil){
		fmt.Fprint(w, "Request Body not found")
	}
	md5Hash := md5.Sum([]byte(body))
	fmt.Fprintf(w, hex.EncodeToString(md5Hash[:])+"\n")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}
