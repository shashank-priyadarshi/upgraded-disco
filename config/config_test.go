package config

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFetchConfig(t *testing.T) {
	// Set up test environment
	os.Setenv("DB_NAME", "test")
	os.Setenv("SQL_URI", "test")
	os.Setenv("MONGO_URI", "test")
	os.Setenv("SERVER_PORT", "test")
	os.Setenv("GITHUB_TOKEN", "test")
	os.Setenv("TODO_API_PORT", "test")
	os.Setenv("ALLOWED_ORIGIN", "test")
	os.Setenv("GITHUB_USERNAME", "test")
	os.Setenv("GH_INTEGRATION_ORIGIN", "test")
	os.Setenv("SERVER_ORIGIN", "test")
	os.Setenv("SECRET_KEY", "MTIzNDU2Nzg5MA==") // Base64-encoded "1234567890"
	os.Setenv("BIO", "test")
	os.Setenv("GITHUB", "test")
	os.Setenv("TODOS", "test")
	os.Setenv("GRAPH", "test")
	os.Setenv("SCHEDULE", "test")

	expectedValue := Configuration{
		DBNAME:            "test",
		SQLURI:            "test",
		MONGOURI:          "test",
		SERVERPORT:        "test",
		TODOAPIPORT:       "test",
		SERVERORIGIN:      "test",
		GHINTEGRATIONPORT: "test",
		GITHUBTOKEN:       "test",
		GITHUBUSERNAME:    "test",
		ALLOWEDORIGIN:     "test",
		SECRETKEY:         []byte("1234567890"),
		Collections: Collections{
			BIODATA:    "test",
			GITHUBDATA: "test",
			TODOS:      "test",
			GRAPHDATA:  "test",
			SCHEDULE:   "test",
		},
	}

	actualValue := FetchConfig()

	if !strings.EqualFold(strings.TrimSpace(string(actualValue.SECRETKEY)), strings.TrimSpace(string(expectedValue.SECRETKEY))) {
		t.Errorf("fetchSecretKey() returned %v, expected %v", string(actualValue.SECRETKEY), string(expectedValue.SECRETKEY))
	}

	os.Setenv("SECRET_KEY", "896876==") // Base64-encoded "1234567890"
	actualValue = FetchConfig()

	if reflect.DeepEqual(expectedValue.SECRETKEY, actualValue.SECRETKEY) {
		t.Errorf("unexpectedly equal: fetchSecretKey() returned %v, expected %v", actualValue.SECRETKEY, expectedValue.SECRETKEY)
	}

	// Clean up test environment
	os.Unsetenv("DB_NAME")
	os.Unsetenv("SQL_URI")
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("GITHUB_TOKEN")
	os.Unsetenv("TODO_API_PORT")
	os.Unsetenv("ALLOWED_ORIGIN")
	os.Unsetenv("GITHUB_USERNAME")
	os.Unsetenv("GH_INTEGRATION_ORIGIN")
	os.Unsetenv("SERVER_ORIGIN")
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("BIO")
	os.Unsetenv("GITHUB")
	os.Unsetenv("TODOS")
	os.Unsetenv("GRAPH")
	os.Unsetenv("SCHEDULE")
}

func TestFetchSecretKey(t *testing.T) {
	// Set up test environment
	os.Setenv("SECRET_KEY", "MTIzNDU2Nzg5MA==") // Base64-encoded "1234567890"

	// Call the function being tested
	key := fetchSecretKey()

	// Verify the results
	expectedKey := []byte("1234567890")
	if !bytes.Equal(key, expectedKey) {
		t.Errorf("fetchSecretKey() returned %v, expected %v", key, expectedKey)
	}

	// Clean up test environment
	os.Unsetenv("SECRET_KEY")
}
