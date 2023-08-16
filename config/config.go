package config

type CommandConfig struct {
	Name        string
	Default     string
	Description string
}

var (
	Region = CommandConfig{
		Name:        "region",
		Default:     "1",
		Description: "The city that you want to retrieve the data of branches from",
	}
	Command = CommandConfig{
		Name:        "command",
		Default:     "status",
		Description: "Command name that can be executed",
	}
)
