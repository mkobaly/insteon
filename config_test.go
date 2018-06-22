package insteon

import (
	"io/ioutil"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	if cfg.Creds.Password != "password" {
		t.Errorf("Expected 'password' but but got: %d", cfg.Creds.Password)
	}
}

func TestLoadConfigFromFile(t *testing.T) {
	content, _ := ioutil.ReadFile("config.yaml")
	cfg := LoadConfig(content)
	if cfg.Creds.Password != "aaaaaaaaa" {
		t.Errorf("Expected 'aaaaaaaaa' but but got: %d", cfg.Creds.Password)
	}
}

func TestLoadConfigFromString(t *testing.T) {
	yaml := `creds:
  clientID: clientidsecret
  username: username1
  password: password1`

	content := []byte(yaml)
	cfg := LoadConfig(content)
	if cfg.Creds.Password != "password1" {
		t.Errorf("Expected 'password1' but but got: %d", cfg.Creds.Password)
	}
}
