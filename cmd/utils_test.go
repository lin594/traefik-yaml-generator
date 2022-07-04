package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Mkdir("./this-just-a-test", 0755)
	defer os.RemoveAll("./this-just-a-test")
	ioutil.WriteFile("./this-just-a-test/test-load-config.yml", []byte("version: 3\n"), 0644)
	config, err := loadConfig("./this-just-a-test/test-load-config.yml")
	if err != nil {
		t.Fatal(err)
	}
	if config.GetString("version") != "3" {
		t.Error("config load error")
	}
}
