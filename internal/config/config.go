package config

// HelloConfig is the configuration structure
type HelloConfig struct {
	Hello struct {
		Name string `yaml:"name" env:"HELLO_NAME" env-default:"Unknown"`
	}
}

// Args defines a configuration path
type Args struct {
	ConfigPath string
}
