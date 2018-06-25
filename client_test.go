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

func TestFullCommand(t *testing.T) {
	url := "https://connect.insteon.com/api/v2"
	client := New(url)
	err := client.Authenticate("44444444-d5b1-42cb-8810-xxxxxxxxx", "foo@gmail.com", "pwd")
	if err != nil {
		t.Errorf("Expected to Authenticate. Instead got error %v", err)
	}

	status, err := client.SendCommand("get_sensor_status", 1248516)
	if err != nil {
		t.Errorf("didn't expect error but got one %v", err)
	}

	resp, err := client.CommandStatus(status.ID)
	if err != nil {
		t.Errorf("didn't expect error but got one %v", err)
	}
	t.Logf("Status %v", resp.Status)
}
