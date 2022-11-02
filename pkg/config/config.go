package config

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	StartURL   string
	URLFilter  string
	Workers    int
	MaxDepth   int32
	UserAgents []string
	Proxies    []string
	RateLimit  struct {
		Rate     int
		Interval time.Duration
	}
	Storage struct {
		DB   string
		HTML string
	}
}

// ParseConfig
func ParseConfig(cf string) (*Config, error) {
	f, err := ioutil.ReadFile(cf)
	if err != nil {
		return nil, err
	}
	var cnf = struct {
		StartURL   string   `yaml:"start_url"`
		URLFilter  string   `yaml:"url_filter"` // only visit url in this pattern
		Workers    int      `yaml:"workers"`
		MaxDepth   int32    `yaml:"max_depth"`
		UserAgents []string `yaml:"user_agents"`
		Proxies    []string `yaml:"proxies"`
		RateLimit  string   `yaml:"rate_limit"`
		Storage    struct {
			DB   string `yaml:"db"`
			HTML string `yaml:"html"`
		} `yaml:"storage"`
	}{}

	if err := yaml.Unmarshal(f, &cnf); err != nil {
		return nil, err
	}

	rate, interval := parseRateLimit(cnf.RateLimit)
	//
	return &Config{
		StartURL:   cnf.StartURL,
		URLFilter:  cnf.URLFilter,
		Workers:    cnf.Workers,
		MaxDepth:   cnf.MaxDepth,
		UserAgents: cnf.UserAgents,
		Proxies:    cnf.Proxies,
		RateLimit: struct {
			Rate     int
			Interval time.Duration
		}{
			Rate:     rate,
			Interval: interval,
		},
		Storage: struct {
			DB   string
			HTML string
		}{
			DB:   cnf.Storage.DB,
			HTML: cnf.Storage.HTML,
		},
	}, nil
}

// Validate
func (c *Config) Validate() error {
	// TODO: make validation
	return nil
}

// parseRateLimit format 10/1s  => 10 request per second
// or fallback to default
func parseRateLimit(s string) (int, time.Duration) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return 5, 1 * time.Second
	}

	i, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		return 5, 1 * time.Second
	}

	d, err := time.ParseDuration(parts[1])
	if err != nil {
		return 5, 1 * time.Second
	}

	return int(i), d
}
