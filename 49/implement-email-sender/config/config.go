package config

import (
	"dashboard-ecommerce-team2/helper"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	AppName     string
	Debug       bool
	Port        string
	SecretKey   string
	MigrateUsed bool
	DBConfig    DBConfig
	RedisConfig RedisConfig
	MailerSend  MailerSend
}

type DBConfig struct {
	DBName         string
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBTimeZone     string
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxIdleTime  int
	DBMaxLifeTime  int
}

type RedisConfig struct {
	Url      string
	Password string
	Prefix   string
}

type MailerSend struct {
	ApiKey string
}

func ReadConfig() (Configuration, error) {
	err := godotenv.Load()
	if err != nil {
		return Configuration{}, err
	}
	return Configuration{
		AppName:     os.Getenv("APP_NAME"),
		Debug:       helper.StringToBool(os.Getenv("DEBUG")),
		Port:        os.Getenv("PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		MigrateUsed: helper.StringToBool(os.Getenv("MIGRATE_USED")),
		DBConfig: DBConfig{
			DBName:         os.Getenv("DB_NAME"),
			DBUsername:     os.Getenv("DB_USERNAME"),
			DBPassword:     os.Getenv("DB_PASSWORD"),
			DBHost:         os.Getenv("DB_HOST"),
			DBTimeZone:     os.Getenv("DB_TIMEZONE"),
			DBMaxIdleConns: helper.StringToInt(os.Getenv("DB_MAX_IDLE_CONNS")),
			DBMaxOpenConns: helper.StringToInt(os.Getenv("DB_MAX_OPEN_CONNS")),
			DBMaxIdleTime:  helper.StringToInt(os.Getenv("DB_MAX_IDLE_TIME")),
			DBMaxLifeTime:  helper.StringToInt(os.Getenv("DB_MAX_LIFE_TIME")),
		},
		MailerSend: MailerSend{
			ApiKey: os.Getenv("API_KEY_MAILSENDER"),
		},
	}, nil
}
