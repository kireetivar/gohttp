package gohttp

import (
	"errors"
	"fmt"
)

type Version int

const (
	VersoinUnknown Version = iota
	HTTP10
	HTTP11
)

var ErrUnsupportedVersion = errors.New("unsupported HTTP version")

var versionNames = [...]string{"UNKNOWN", "HTTP/1.0", "HTTP/1.1"}

func (p Version) String() string {
	if int(p) < 0 || int(p) >= len(versionNames) {
		return VersoinUnknown.String()
	}
	return versionNames[p]
}

func ParseProtocol(s string) (Version, error) {
	switch s {
	case "HTTP/1.0":
		return HTTP10, nil
	case "HTTP/1.1":
		return HTTP11, nil
	}
	return VersoinUnknown, fmt.Errorf("%w: %q", ErrUnsupportedVersion, s)
}
