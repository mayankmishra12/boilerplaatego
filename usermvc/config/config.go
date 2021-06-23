//Config/Database.go
package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

const (
	APP_ENV     = "APP_ENV"
	PRODUCTION  = "production"
	DEVELOPMENT = "development"
)

// DBConfig represents db configuration
var (
	configDir      = "./../config"
	configfile     = "developement"
	configFileType = "yaml"
)
var (
	Config *AppConfig
)

type dBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
}

type AppConfig struct {
	App      *appcofig
	DbConfig *dBConfig
	Logger   *logger
}
type logger struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	LocalTime  bool   `json:"local_time"`
	Compress   bool   `json:"compress"`
}

type appcofig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig() *AppConfig {

	viper.AutomaticEnv()
	env := viper.Get(APP_ENV)
	if env == PRODUCTION {
		configfile = PRODUCTION
	}
	fmt.Println("priting env is golang", env)
	viper.SetConfigName(configfile)
	viper.AddConfigPath(configDir)
	viper.SetConfigType(configFileType)
	var configuration AppConfig
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variables using the model..")
	return &configuration
}

func NewDbConfig() *dBConfig {
	return &dBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "mayank",
		DBName:   "postgres",
		Password: "",
	}
}
