package config

import "flag"

type Config struct {
	Url     string
	Timeout int
}

func Load() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.Url, "url", "", "url to search")
	flag.IntVar(&cfg.Timeout, "timeout", 5, "timeout in seconds")
	flag.Parse()

	return cfg
}
