package config

// Config
type Config struct {
	StartURL     string   `yaml:"start_url"`
	URLFilter    string   `yaml:"url_filter"` // only visit url in this pattern
	Workers      int      `yaml:"workers"`
	MaxDepth     int32    `yaml:"max_depth"`
	UserAgents   []string `yaml:"user_agents"`
	Proxies      []string `yaml:"proxies"`
	TargetScope  string   `yaml:"target_scope"`
	ParsePattern string   `yaml:"parse_pattern"`
	Storage      struct {
		DB   string `yaml:"db"`
		HTML string `yaml:"html"`
	} `yaml:"storage"`
}

// ParseConfig
func ParseConfig(cf string) (*Config, error) {
	//
	return &Config{}, nil
}
