package insteon

import (
	yaml "gopkg.in/yaml.v2"
)

type Credential struct {
	ClientID string
	Username string
	Password string
}

type Config struct {
	Creds Credential
}

//NewConfig creates a new Configuration object needed
func NewConfig() *Config {
	//config := Config{}
	return &Config{
		Creds: Credential{Username: "youremail@example.com", Password: "password", ClientID: "api_key_from_insteon"},
	}
}

//LoadConfig will load up a Config object based on configPath
func LoadConfig(data []byte) *Config {
	//config := Config{}
	var config = new(Config)
	// data, err := ioutil.ReadFile(configPath)
	// if err != nil {
	// 	panic(err.Error())
	// }

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
