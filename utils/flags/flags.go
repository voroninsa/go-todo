package flags

import "flag"

type Flags struct {
	ConfigPath string
}

func ParseFlags() *Flags {
	config := flag.String("config", "../config.json", "enter destination of project configuration file")
	flag.Parse()

	flags := &Flags{
		ConfigPath: *config,
	}

	return flags
}
