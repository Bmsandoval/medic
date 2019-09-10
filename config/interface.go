package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configurations struct {
	MinikubeConfig MinikubeConfig
}

// initConfig reads in config file and ENV variables if set.
func GenerateSampleConfigYaml(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	c := Configurations{}
	cYaml, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if _, err := f.WriteString(string(cYaml)); err != nil {
		return err
	}

	err = f.Close()
	return err
}

func LoadConfigYaml(filePath string) (*Configurations, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var c Configurations
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c, nil
}

func StoreConfigInViper(config *Configurations) {
	viper.Set("configs", config)
}

func GetConfigFromViper() *Configurations {
	config := viper.Get("configs")
	return config.(*Configurations)
}


//func (c *Configurations) Configure(v viper.Viper) {
//	c.DatabaseConfig = GetDatabaseConfig(v)
//}

//const VAULT_SECRETS_PATH = "VAULT_SECRETS_PATH"
//func RegisterVaultEnvVarsOrFatal(appCtx appcontext.Context, v viper.Viper) (*Configurations, error) {
//	var secretsMap credentials
//
//	secretsPath := appCtx.ConfigurationInstance.GetString(VAULT_SECRETS_PATH)
//
//	secrets, err := ioutil.ReadFile(secretsPath)
//	if err != nil {
//		log.Fatalf("Failed to read the secrets file")
//		return nil, err
//	}
//
//	err = json.Unmarshal(secrets, &secretsMap)
//	if err != nil {
//		log.Fatalf("Failed to unmarshal the secrets file")
//		return nil, err
//	}
//
//	configs := Configurations{}
//	configs.Configure(v, secretsMap)
//
//	return &configs, nil
//}
