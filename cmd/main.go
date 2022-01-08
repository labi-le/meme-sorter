package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"meme-sorter/internal"
	"meme-sorter/web"
	"os"
)

var (
	config string
)

func init() {
	flag.StringVar(&config, "config", "config.toml", "path to config file")
}

func main() {
	flag.Parse()
	Config := internal.Config{}

	_, err := toml.DecodeFile(config, &Config)

	if err != nil {
		log.Fatal(err)
	}

	db := internal.NewDB(Config)
	Config.DB = db

	if db.Migrate() != nil {
		log.Fatal(err)
	}

	if flag.Arg(1) == "generate_db" {
		parseImageAndAddDb(&Config)
		os.Exit(0)
	}

	err = web.Start(&Config)
	if err != nil {
		log.Fatal(err)
	}
}

func parseImageAndAddDb(Config *internal.Config) {
	dir, err := ioutil.ReadDir(Config.ItemsDir)
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		byteFile, err := ioutil.ReadFile(Config.ItemsDir + string(os.PathSeparator) + file.Name())
		if err != nil {
			panic(err)
		}

		err = Config.DB.Create(&internal.Meme{
			Model: gorm.Model{},
			Fun:   false,
			Image: byteFile,
		})

		if err != nil {
			panic(err)
		}
	}
}
