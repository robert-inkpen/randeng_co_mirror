package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Test Server Config Instantiates", func(t *testing.T) {
		expectedAddre := "127.0.0.1"
		expectedPort := "8080"
		err := os.Setenv(BackendServerAddress, "127.0.0.1")
		err = os.Setenv(BackendServerPort, "8080")
		if err != nil {
			t.Error(err)
		}

		cfg, err := Load()
		if err != nil {
			t.Error(err)
		}

		if cfg.ServerConfig.Address != expectedAddre || cfg.ServerConfig.Port != expectedPort {
			t.Errorf("Expected Address and Port to be configured correctly")
		}
	})
}
