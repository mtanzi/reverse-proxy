package cmd

import "flag"

type Cmd struct {
	SSL string
}

func ParseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.SSL, "ssl", "true", "enable SSL connection")
	flag.Parse()
	return cmd
}
