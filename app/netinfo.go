package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	hostname, error := os.Hostname()
	if error != nil {
		panic(error)
	}

	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)
	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr).IP.String()

	io.WriteString(w, ("Hostname: " + hostname + "\nIpAddress: " + ipAddress + "\n"))

}
