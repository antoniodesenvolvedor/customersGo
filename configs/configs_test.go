package configs

import (
	"testing"
)
//
func init() {
	LoadEnvFromFile()
}

func TestMeu(t *testing.T) {
	if configs.DbConnectionString == "" {
		t.Fatalf("Não foi possível carregar as variáveis de ambiente")
	}
}
