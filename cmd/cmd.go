package cmd

import "flag"

// Cmd struct containing the flags data
type Cmd struct {
	ConfigPath string
}

// ParseCmd parse the flags and assign them to the Cms struct
func ParseCmd() Cmd {
	var cmd Cmd

	flag.StringVar(&cmd.ConfigPath, "config_path", "config.json", "file containing the proxy configuration")
	flag.Parse()

	return cmd
}
