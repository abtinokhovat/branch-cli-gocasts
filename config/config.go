package config

type CommandConfig struct {
	Name        string
	Default     string
	Description string
}

var Region = CommandConfig{
	Name:        "region",
	Default:     "Tehran",
	Description: "The city that you want to retrieve the data of branches from",
}

var Command = CommandConfig{
	Name:        "command",
	Default:     "exit",
	Description: "Command name that can be executed",
}
