package config

// Config is the struct for the
type Config struct {
	DefaultURL  string
	DefaultPort string
	Rules       []Rule
}

// Rule define the struct for a single rule
type Rule struct {
	Matcher        string
	DownstreamPort string
}

// InitConfig initialise the configuration
func InitConfig() Config {
	defaultURL := "localhost"
	defaultPort := "1333"

	rules := []Rule{
		Rule{
			Matcher:        "/marco",
			DownstreamPort: "1331",
		},
		Rule{
			Matcher:        "/catarina",
			DownstreamPort: "1332",
		},
	}

	return Config{
		DefaultURL:  defaultURL,
		DefaultPort: defaultPort,
		Rules:       rules,
	}
}
