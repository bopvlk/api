package main

import (
	"flag"
	"log"
	"os"

	"github.com/BO/6.ServerAndDB/internal/app/api"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

var (
	configPath      string = "configs/api.toml"
	correctSettings string = ".toml"
)

func init() {
	// розкажемо що наш додаток буде полуучати шлях до конфіг файла із іншої папки
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
	//флаг для визначення звідки ми будемо брати настройочні файли
	flag.StringVar(&correctSettings, "format", "configs/api.toml", "format to change settings point")
}

func SettingsPoint(config *api.Config) {
	switch correctSettings {
	case ".toml":
		_, err := toml.DecodeFile(configPath, config) // десереалізаціся вміст .toml файлу
		if err != nil {
			log.Println("cen not find configs file. using default vaules:", err)
		}
		log.Println("find port in a .toml file")
	case ".env":
		err := godotenv.Load(configPath)
		if err != nil {
			log.Println("could not find .env file:", err)
		}
		log.Println("find port in a .env file")
		config.BindAddr, config.LoggerLevel = os.Getenv("app_port"), os.Getenv("logger_level")
	default:
		_, err := toml.DecodeFile(configPath, config) // десереалізаціся вміст .toml файлу
		if err != nil {
			log.Println("cen not find configs file. using default vaules:", err)
		}
		log.Println("find port in a .toml file")
	}
}

func main() {
	// в цей момен виконується ініціалізація зімінної  configPath значенням
	flag.Parse()
	log.Println("it`s works")
	//server instance initialization
	config := api.NewConfig()

	SettingsPoint(config)

	server := api.New(config)

	// api server start
	log.Fatal(server.Start())
	// \/ це ще один варінант написання /\ 61 рядка
	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }

}
