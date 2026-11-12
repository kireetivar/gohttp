package gohttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Timeout time.Duration
}

func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	n, err := net.Dial("tcp", req.Host)
	if err != nil {
		fmt.Printf("error opening connection %v \n", err)
		return &Response{}, err
	}
	defer n.Close()
	n.SetReadDeadline(time.Now().Add(5 * time.Second))

	fmt.Fprintf(n, "%v %v %v\r\nHost: %v\r\nUser-Agent: curl/8.0\r\n\r\n", req.Method, req.Path, req.Version, req.Host)
	b := make([]byte, 2048)

	l, err := n.Read(b)
	if err != nil {
		fmt.Println("error reading from connection")
		return &Response{}, err
	}
	res := Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b[:l]))),
	}
	return &res, nil
}

func Get(ctx context.Context, rawUrl string) (*Response, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return &Response{}, errors.New("parsing url failed")
	}
	req := Request{
		Version: Version(2),
		Method:  "GET",
		Headers: map[string]string{
			"User-Agent": "curl/8.0",
		},
		Host: u.Host,
		Path: u.Path,
	}
	c := Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Do(ctx, &req)
	if err != nil {
		return &Response{}, errors.New("get call failed")
	}
	return res, nil
}
