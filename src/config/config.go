package config

import (
	"log"
	"os"
//	"time"

	"github.com/BurntSushi/toml"
	rpc "github.com/xiphiar/secret-validator_exporter/getData/rpc"
	rest "github.com/xiphiar/secret-validator_exporter/getData/rest"
//	"github.com/spf13/viper"
)

const (
)

var (
	ConfigPath string
	Config	configType
)


type configType struct {

	Title	string	`json:"title"`

	Servers struct {
                Addr struct {
                        RPC	string `json:"rpc"`
                        REST	string `json:"rest"`
                }
        }

	Validator struct {
		OperatorAddr	string	`json:"operatorAddr"`
	}

	Options	struct {
		ListenPort	string	`json:"listenPort"`
	}
}


func Init() {

	//Config = readConfig()

	rpc.Addr = os.Getenv("RPC_URL")
	rest.Addr = os.Getenv("LCD_URL")

	rest.OperAddr = os.Getenv("OPER_ADDR")

}

func readConfig() configType {

        var config configType

//	path := viper.GetString(ConfigPath)+"/config.toml"

//        if _, err := toml.DecodeFile(path, &config); err != nil{
        if _, err := toml.DecodeFile(ConfigPath +"/config.toml", &config); err != nil{

                log.Fatal("Config file is missing: ", config)
        }

	return config

}
