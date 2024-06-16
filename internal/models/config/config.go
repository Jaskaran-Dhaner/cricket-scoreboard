package config

import "go.mongodb.org/mongo-driver/mongo"

type Config struct {
	Server      Server      `yaml:"server"`
	Database    Database    `yaml:"database"`
	Collections Collections `yaml:"collections"`
	Version     string      `yaml:"version"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Client   *mongo.Client   //`yaml:"client"`
	DB       *mongo.Database //`yaml:"db"`
	Host     string          `yaml:"host"`
	Port     string          `yaml:"port"`
	User     string          `yaml:"user"`
	MongoURI string          `yaml:"mongouri"`
	Password string          `yaml:"password"`
	DBName   string          `yaml:"dbname"`
	SSLMode  string          `yaml:"sslmode"`
}

type Collections struct {
	Users      string `yaml:"users"`
	Teams      string `yaml:"teams"`
	Match      string `yaml:"match"`
	Tournament string `yaml:"tournament"`
	Roles      string `yaml:"roles"`
}
