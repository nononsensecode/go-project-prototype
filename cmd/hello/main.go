package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	config "nononsensecode.com/my_project/config"
	greetings "nononsensecode.com/my_project/pkg/greetings"
)

func main() {
	var cfg config.HelloConfig
	err := cleanenv.ReadConfig("../../config/config.yaml", &cfg)
	if err != nil {
		panic("Configuration not found")
	}

	fmt.Printf("%s\n", greetings.Hello(cfg.Hello.Name))
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) config.Args {
	var a config.Args

	f := flag.NewFlagSet("Example server", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yaml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprint(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}
