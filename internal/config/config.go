package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Cgf struct {
	App struct {
		Host string `yaml:"host" env:"APP_HOST"`
		Port string `yaml:"port" env:"APP_PORT"`
	} `yaml:"app"`

	Database struct {
		Host string `yaml:"host" env:"DB_HOSt"`
		Port string `yaml:"port" env:"DB_PORT"`
		User string `yaml:"user" env:"DB_USER"`
		Pwd  string `yaml:"pwd" env:"DB_PWD"`
		Name string `yaml:"name" env:"DB_NAME"`
	} `yaml:"database"`
}

func loadSettings() Cgf {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env:\n%v", err)
	}

	cfg := Cgf{}
	if err := cleanenv.ReadConfig("conf.yml", &cfg); err != nil {
		log.Fatalf("error loading conf.yml:\n%v", err)
	}

	return cfg
}

var Settings = loadSettings()
