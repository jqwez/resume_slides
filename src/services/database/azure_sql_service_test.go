package database

import (
	"os"
	"testing"
)

func _TestService() *AzureSQLService {
	config := MustAzureSQLConfigFromEnv()
	service := NewAzureSQLService(config)
	return service
}

func TestTestServer(t *testing.T) {
	service := _TestService()
	if service.conn == nil {
		t.Fatal("Test Server Failed to Load, check .env")
	}
}

func TestMustAzureSQLConfigFromEnv(t *testing.T) {
	config := MustAzureSQLConfigFromEnv()
	if config == nil {
		t.Fatal("Default config failed to load")
	}
	os.Setenv("AZURE_DATABASE", "")
}

func TestGetConnection(t *testing.T) {
	service := _TestService()
	_ = service
	t.Log("test connection")
	service.conn.Close()
	err := service.conn.Ping()
	if err == nil {
		t.Error(err)
	}
	service.GetConnection()
	err = service.conn.Ping()
	if err != nil {
		t.Error(err)
	}
}
