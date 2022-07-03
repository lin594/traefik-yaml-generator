package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

// 从 file 中读取一个 Viper 对象
func loadConfig(file string) (*viper.Viper, error) {
	config := viper.New()
	file, fileName := filepath.Split(file)
	ext := filepath.Ext(fileName)
	if ext != ".yaml" && ext != ".yml" {
		return nil, fmt.Errorf("the file extension is not yaml")
	}
	config.AddConfigPath(file)     //设置读取的文件路径
	config.SetConfigFile(fileName) //设置读取的文件
	err := config.ReadInConfig()
	return config, err
}

// 检查 service 中的各个属性
func checkService(config *viper.Viper, service Service) {
	property := "services." + service.name + "."

	config.SetDefault(property+"container_name", service.name)

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
	labels = append(labels, fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.entrypoints=%s", service.name, service.entrypoints))
	if service.tls {
		labels = append(labels, fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.tls=true", service.name))
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
		config.Set(property+"network", network)
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
	networks[service.network].(map[string]interface{})["external"] = "true"
	config.Set("networks", networks)
}
