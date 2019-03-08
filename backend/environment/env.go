package environment

import (
	"os"
	"strconv"
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Migration MigrationCongig
}

type MigrationCongig struct {
	Version string
}

func MigrationVesion() int {
	var config Config
	_, err := toml.DecodeFile("conf/migration.tml", &config)
	if err != nil {
		log.Fatal("not found migration.toml")
		return 0
	}

	version, err := strconv.Atoi(config.Migration.Version)
	if err != nil {
		return 0
	}
	return version
}

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}

func OriginSiteUrl() string {
	url := os.Getenv("ORIGIN_SITE_URL")
	if url == "" {
		return "http://localhost:9000"
	}
	return url
}

func AdminSiteUrl() string {
	url := os.Getenv("ADMIN_SITE_URL")
	if url == "" {
		return "http://localhost:7000"
	}
	return url
}

func DBHost() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return ""
	}
	return host
}

func DBUser() string {
	user := os.Getenv("DB_USER")
	if user == "" {
		return ""
	}
	return user
}

func DBPassword() string {
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return ""
	}
	return pass
}