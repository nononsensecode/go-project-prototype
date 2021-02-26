package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	config "nononsensecode.com/my_project/internal/config"
	greetings "nononsensecode.com/my_project/pkg/greetings"
)

func main() {
	var cfg config.HelloConfig

	args := ProcessArgs(&cfg)

	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("%s\n", greetings.Hello(cfg.Hello.Name))
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) config.Args {
	var a config.Args

	f := flag.NewFlagSet("Hello", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yaml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}
