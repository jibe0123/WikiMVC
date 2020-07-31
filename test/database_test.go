package test

import (
	"github.com/jibe0123/WikiMVC/database"
	"strings"
	"testing"
)

func connectTest(t *testing.T) {
	tests := []struct {
		name         string
		description string
		want         string
	}{
		{
			description: "Is there errors",
			want:         "No errors in err",
		},
		{
			name:         "error",
			description: "Is there errors",
			want:         "Database configuration error",
		},
	}

	for _, oneTest := range tests {
		t.Run(oneTest.name, func(t *testing.T) {
			_, err := database.Connect()
			if err != nil {
				// Database configuration error
				isBadDatabaseConfig := strings.Contains(err.Error(), "Access denied for user")
				if isBadDatabaseConfig {
					return
				}
				t.Errorf("connectTest() = %v, want %v", err.Error(), oneTest.want)
			}
		})
	}
}
