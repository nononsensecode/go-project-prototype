package config

// HelloConfig is the configuration structure
type HelloConfig struct {
	Hello struct {
		Name string `yaml:"name" env:"HELLO_NAME" env-default:"Unknown" env-description:"Name to be mentioned in the hello greeting"`
	} `yaml: hello`
}

// Args defines a configuration path
type Args struct {
	ConfigPath string
}
