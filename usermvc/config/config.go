//Config/Database.go
package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// DBConfig represents db configuration
var (
	configDir      = ""
	configfile     = ""
	configFileType = ""
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
}

type appcofig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig() *AppConfig {
	viper.SetConfigName("conf")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration AppConfig

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	//fmt.Println("Database is\t", configuration.DbConfig.Host)
	//fmt.Println("Port is\t\t", configuration.DbConfig.Port)
	//fmt.Println("Port is\t\t", configuration.DbConfig.User)

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

//Db should be access through scret in which taken from env
// logger should be middleware
//
