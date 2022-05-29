package config

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/siroj100/hikarie-islamy/pkg/errorx"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func Init() Config {
	// init default config for the maps, otherwise Viper can't found some attributes
	config := Config{
		Database: map[string]DatabaseConfig{DbIslamy: {}},
	}

	b, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	//log.Println("yaml:\n", string(b))
	defaultConfig := bytes.NewReader(b)
	viper.SetConfigType("yaml")
	if err = viper.MergeConfig(defaultConfig); err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	viper.SetConfigName("backend")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/hikarie-islamy/")
	err = viper.MergeInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fmt.Printf("keys: %+v\n\n", viper.AllKeys())
	//for key, val := range viper.AllSettings() {
	//	fmt.Printf("%s: %+v\n", key, val)
	//}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	if ok := validateConfig(&config); !ok {
		log.Fatalln("Incomplete configuration file")
	}

	if config.Server.Port < 1 {
		config.Server.Port = 8080
	}

	return config
}

func validateConfig(conf *Config) bool {
	return validateServerConfig(conf) &&
		validateDbConfig(conf.Database)
}

func validateServerConfig(conf *Config) bool {
	if conf.Server.Port < 1 {
		conf.Server.Port = 8080
	}
	return true
}

func validateDbConfig(dbConf map[string]DatabaseConfig) bool {
	dbListMap := make(map[string]bool)
	result := true
	for _, dbName := range DbList {
		dbListMap[dbName] = true
		//if _, found := dbConf[dbName]; !found {
		//	log.Printf("Database configuration not found: %s\n", dbName)
		//	result = false
		//}
	}
	for dbName, _ := range dbConf {
		if _, found := dbListMap[dbName]; !found {
			log.Printf("Unnecessary database configuration found: %s\n", dbName)
			delete(dbConf, dbName)
		}
	}
	return result
}
