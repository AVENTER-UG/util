package util

import (
	"os"
	"strings"
	"time"

	"github.com/AVENTER-UG/util/vault"
)

// Getenv with default value
func Getenv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	if strings.Contains(value, "vault://") {
		VaultToken := os.Getenv("VAULT_TOKEN")
		VaultURL := Getenv("VAULT_URL", "http://127.0.0.1:8200")
		VaultTimeout, _ := time.ParseDuration(Getenv("VAULT_TIMEOUT", "10s"))

		v := vault.New(VaultToken, VaultURL, VaultTimeout)
		if VaultToken != "" {
			v.Connect()
			value = v.GetKey(value)
		}
	}

	return value
}
