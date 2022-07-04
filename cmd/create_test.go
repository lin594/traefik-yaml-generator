package cmd

import (
	"os"
	"testing"
)

func CompareTwoStringSlice(t *testing.T, s1 []string, s2 []string) {
	if len(s1) != len(s2) {
		t.Error("unexpected labels length")
	}
	for i, v := range s1 {
		if v != s2[i] {
			t.Error("unmatch labels!" + v + " " + s2[i])
		}
	}
}

func TestCreate(t *testing.T) {
	t.Log("test create")
	os.Mkdir("./this-just-a-test", 0755)
	defer os.RemoveAll("./this-just-a-test")
	cmd := rootCmd
	cmd.SetArgs([]string{"create", "test-service", "--service", "test", "-i", "test/test:latest", "-a", "test.com", "-o", "./this-just-a-test/test.yml"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	config, err := loadConfig("./this-just-a-test/test.yml")
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
		"traefik.http.services.test.loadbalancer.server.port=80",
		"traefik.http.routers.test.entrypoints=websecure",
		"traefik.http.routers.test.tls=true",
		"traefik.http.routers.test.middlewares=test-compress",
		"traefik.http.middlewares.test-compress.compress=true"}
	CompareTwoStringSlice(t, config.GetStringSlice("services.test-service.labels"), labels)

}
