package gohttp

import "io"

type Response struct {
	Version    Version
	StatusCode int
	Status     string
	Body       io.ReadCloser
}
