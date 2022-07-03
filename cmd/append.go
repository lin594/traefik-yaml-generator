package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
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

var appendCmd = &cobra.Command{
	Use:   "append file [service name]",
	Short: "Append subcommand appends labels which traefik needs to the yaml file.",
	Long: "Append subcommand appends labels which traefik needs to the yaml file.\n" +
		"You should specify a yaml file and a service name which tygen will modify for enabling Traefik." +
		" If you don't specify a service name, tygen will choose any one of these services to continue.",
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {

		// config 是可以读取的配置文件
		config, err := loadConfig(args[0])
		if err != nil {
			return err
		}

		// 确认要添加 Traefik 的服务名称
		if len(args) == 2 {
			service.name = args[1]
			if config.Get(service.name) == nil {
				return fmt.Errorf("service %s not found", service.name)
			}
		} else {
			for name := range config.GetStringMap("services") {
				service.name = name
				break
			}
			if service.name == newService().name {
				return fmt.Errorf("no service name specified")
			}
		}

		// 检查 service 中的属性
		checkService(config, service)

		// 检查 networks
		checkNetworks(config, service)

		// 写回配置文件
		if output == "" {
			output = args[0]
		}
		return config.WriteConfigAs(output)
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
}
