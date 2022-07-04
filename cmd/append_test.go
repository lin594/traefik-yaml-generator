package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	os.Mkdir("./this-just-a-test", 0755)
	defer os.RemoveAll("./this-just-a-test")
	ioutil.WriteFile("./this-just-a-test/test-append.yml", []byte("version: \"3\"\nservices:\n    test-service:\n"), 0644)
	InitVariables()
	cmd := rootCmd

	t.Log("test append")
	cmd.SetArgs([]string{"append", "./this-just-a-test/test-append.yml",
		"--service", "test",
		"-i", "test/test:latest",
		"-a", "test.com",
		"-p", "8080",
		"-e", "web",
		"-t=false"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	config, err := loadConfig("./this-just-a-test/test-append.yml")
	if err != nil {
		t.Fatal(err)
	}
	if config.GetString("version") != "3" {
		t.Error("bugs exist in field version")
	}
	if config.GetStringMap("services.test-service") == nil {
		t.Fatal("bugs exist in service name")
	}
	if config.GetString("services.test-service.image") != "test/test:latest" {
		t.Error("bugs exist in field image")
		t.Error(config.GetString("services.test-service.image"))
	}
	if config.GetStringMap("networks")[newService().network].(map[string]interface{})["external"] != true {
		t.Error("bugs exist in field external")
	}

	labels := []string{"traefik.http.routers.test.rule=Host(`test.com`)",
		"traefik.http.services.test.loadbalancer.server.port=8080",
		"traefik.http.routers.test.entrypoints=web",
		"traefik.http.routers.test.middlewares=test-compress",
		"traefik.http.middlewares.test-compress.compress=true"}
	CompareTwoStringSlice(t, config.GetStringSlice("services.test-service.labels"), labels)
}
