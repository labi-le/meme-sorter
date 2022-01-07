package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"meme-sorter/internal"
	"meme-sorter/internal/structures"
	"meme-sorter/web"
)

var (
	config string
)

func init() {
	flag.StringVar(&config, "config", "config.toml", "path to config file")
}

func Var_dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func main() {
	flag.Parse()
	Config := structures.Config{}

	_, err := toml.DecodeFile(config, &Config)

	if err != nil {
		log.Fatal(err)
	}

	db := internal.NewDB(Config)

	if db.Migrate() != nil {
		log.Fatal(err)
	}

	Config.DB = db

	err = web.Start(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
