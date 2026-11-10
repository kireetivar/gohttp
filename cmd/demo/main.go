package main

import (
	"fmt"
	"net"
	"time"

	"github.com/kireetivar/gohttp"
)

func main() {
	req := gohttp.Request{
		Version: gohttp.Version(2),
		Method:  "GET",
		Headers: map[string]string{
			"User-Agent": "curl/8.0",
		},
		Host: "info.cern.ch:80",
		Path: "/",
	}
	res, err := Get(req)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(res)
}

func Get(req gohttp.Request) (string, error) {
	n, err := net.Dial("tcp", req.Host)
	if err != nil {
		fmt.Printf("ERROR opening connection %v \n", err)
		return "", err
	}
	defer n.Close()
	n.SetReadDeadline(time.Now().Add(5 * time.Second))

	fmt.Fprintf(n, "%v %v %v\r\nHost: %v\r\nUser-Agent: curl/8.0\r\n\r\n", req.Method, req.Path, req.Version, req.Host)
	b := make([]byte, 2048)

	l, err := n.Read(b)
	if err != nil {
		fmt.Println("ERROR reading connection")
		return "", err
	}
	return string(b[:l]), nil
}
