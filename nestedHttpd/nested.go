// nested handlers to restrict IP range served - Jesse McNelis

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("/tmp"))

	checkip := func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")
		if ip[0] == "127.0.0.1" {
			fs.ServeHTTP(w, r)
		} else {
			fmt.Println("unauthorised access attempt from:", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
	http.HandleFunc("/", checkip)
	err := http.ListenAndServe("127.0.0.1:6061", nil)
	if err != nil {
		fmt.Println(err)
	}
}
