package insteon

import (
	"testing"
)

func TestNew(t *testing.T) {
	url := "https://connect.insteon.com/api/v2"
	client := New(url)
	if client.baseURL != url {
		t.Errorf("Expected %s but but got: %s", url, client.baseURL)
	}

	if client.HTTPClient == nil {
		t.Errorf("Expected HTTPClient to not be nil")
	}
}

func TestAuthenticate(t *testing.T) {
	url := "https://connect.insteon.com/api/v2"
	client := New(url)
	err := client.Authenticate("kkkk", "kjkjjk@gmail.com", "jjkjkkj")
	if err != nil {
		t.Errorf("Expected to Authenticate. Instead got error %v", err)
	}
}
