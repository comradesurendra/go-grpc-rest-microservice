package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var db *sql.DB

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
		Ip       string `yaml:"ip", envconfig:"DB_IP"`
	} `yaml:"database"`
	Paths struct {
		Rootpath string `yaml:"rootpath", envconfig:"ROOT_PATH"`
	} `yaml:"paths"`
	Ports struct {
		RpcPort string `yaml:"rpcport", envconfig:"PRC_PORT"`
		RpcIp   string `yaml:"rpcip", envconfig:"PRC_IP"`
	} `yaml:"ports"`
}

var Cfg Config

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	configpath := "config/config.yaml"
	if os.Getenv("ENV") == "prod" {
		configpath = "config/config.yaml"
	}

	f, err := os.Open(configpath)
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func init() {

	ReadFile(&Cfg)
	// Cfg.Server.Port = os.Getenv("PROXY_PORT")
	// Cfg.Ports.RpcPort = os.Getenv("RPC_PORT")
}

func (cfg *Config) DbConnect() {
	UserName := cfg.Database.Username
	Password := cfg.Database.Password
	DBIP := cfg.Database.Ip
	fmt.Println(DBIP)
	dbobj, err := sql.Open("mysql", UserName+":"+Password+"@tcp("+DBIP+")/db_loan")
	if err != nil {
		fmt.Println(err)
	}
	db = dbobj
}

func GetDb() *sql.DB {
	return db
}

func CloseDb() {
	db.Close()
}

func GeneratePID() int {
	pid := os.Getpid()
	return pid
}
