package cmd

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

// 从 file 中读取一个 Viper 对象
func loadConfig(file string) (*viper.Viper, error) {
	config := viper.New()
	ext := filepath.Ext(file)
	if ext != ".yaml" && ext != ".yml" {
		return nil, fmt.Errorf("the file extension is not yaml")
	}
	config.SetConfigFile(file) //设置读取的文件
	err := config.ReadInConfig()
	return config, err
}

// 创建一个 yaml 格式的 Viper 对象
func newConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("docker-compose.yml")
	config.Set("version", "3")
	return config
}

// 检查 service 中的各个属性
func checkService(config *viper.Viper, name string, service Service) {
	// 填充默认变量
	if service.name == newService().name {
		service.name = name
	}
	if address != "" {
		if service.host == newService().host {
			service.host = fmt.Sprintf("Host(`%s`)", address)
		} else {
			fmt.Println("[Warning] Flag rule is already exist, flag address will be ignored.")
		}
	}

	// 获取 service 中的属性时所用前缀
	property := "services." + name + "."

	// container_name 字段
	config.SetDefault(property+"container_name", service.name)
	if service.name != newService().name {
		config.Set(property+"container_name", service.name)
	}

	// image 字段
	config.SetDefault(property+"image", service.image)
	if service.image != newService().image {
		config.Set(property+"image", service.image)
	}

	// label 字段
	labels := config.GetStringSlice(property + "labels")
	if labels == nil {
		labels = []string{}
	}
	labels = append(labels, fmt.Sprintf("traefik.http.routers.%s.rule=%s", service.name, service.host))
	labels = append(labels, fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.port=%d", service.name, service.port))
	labels = append(labels, fmt.Sprintf("traefik.http.routers.%s.entrypoints=%s", service.name, service.entrypoints))
	if service.tls {
		labels = append(labels, fmt.Sprintf("traefik.http.routers.%s.tls=true", service.name))
	}
	labels = append(labels, fmt.Sprintf("traefik.http.routers.%s.middlewares=%s-compress", service.name, service.name))
	labels = append(labels, fmt.Sprintf("traefik.http.middlewares.%s-compress.compress=true", service.name))
	config.Set(property+"labels", labels)

	// networks 字段
	network := config.GetStringSlice(property + "network")
	flag := false
	for _, n := range network {
		if n == service.network {
			flag = true
			break
		}
	}
	if !flag {
		network = append(network, service.network)
		config.Set(property+"networks", network)
	}

}

// 检查 networks 中的各个属性
func checkNetworks(config *viper.Viper, service Service) {
	networks := config.GetStringMap("networks")
	if networks == nil {
		networks = make(map[string]interface{})
	}
	if _, ok := networks[service.network]; !ok {
		networks[service.network] = make(map[string]interface{})
	}
	networks[service.network].(map[string]interface{})["external"] = true
	config.Set("networks", networks)
}

// 比较两个字符串数组是否相等
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
