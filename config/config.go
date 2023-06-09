package config

import (
	"github.com/AbsaOSS/env-binder/env"
)

var Configutaion Config

type Config struct {
	Port       string `env:"PORT"`
	BackupDir  string `env:"BACKUP_DIR"`
	DbUserName string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
}

func Init() {
	env.Bind(&Configutaion)
}
