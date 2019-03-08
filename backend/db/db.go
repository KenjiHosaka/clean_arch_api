package db

import (
	"clean_arch_api/backend/environment"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Database DbConfig
	Redis    RedisConfig
}

func (config Config) DB() (string, string) {
	return config.Database.Driver, config.Database.DSN()
}

type DbConfig struct {
	Driver    string
	Server    string
	User      string
	Password  string
	Database  string
	Charset   string
	ParseTime string
	TimeZone  string
}

type RedisConfig struct {
	Server    string
	Password  string
	Database  int
}

func (dbConfig DbConfig) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=true&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Server, dbConfig.Database, dbConfig.Charset)
}

func Connect() (*gorm.DB, error) {
	config := readConfig()
	database, err := gorm.Open(config.DB())
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	database.LogMode(true)
	return database, err
}

func readConfig() Config {
	var config Config
	_, err := toml.DecodeFile("conf/database.tml", &config)
	if err != nil {
		log.Fatal("not found database.toml")
	}
	if host := environment.DBHost(); host != "" {
		config.Database.Server = host
	}
	if user := environment.DBUser(); user != "" {
		config.Database.User = user
	}
	if password := environment.DBPassword(); password != "" {
		config.Database.Password = password
	}
	return config
}