package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
