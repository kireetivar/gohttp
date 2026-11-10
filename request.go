package gohttp

import "io"

type Request struct {
	Method  string
	Host    string
	Path    string
	Version Version
	Headers map[string]string
	Body    io.Reader
	Client  string
}
