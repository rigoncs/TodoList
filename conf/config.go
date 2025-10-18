package conf

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Port      string `yaml:"port"`
	Version   string `yaml:"version"`
	JwtSecret string `yaml:"jwtSecret"`
}

type Mysql struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	DriverName string `yaml:"driverName"`
	Database   string `yaml:"database"`
	Charset    string `yaml:"charset"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Network  string `yaml:"network""`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
