package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [service name]",
	Short: "Create subcommand create a docker-compose file with Traefik.",
	Long: "Create subcommand create a docker-compose file with Traefik.\n" +
		"You could specify a service name for tygen to generate docker-compose.yml." +
		" If you don't specify a service name but a service flag, tygen will use the value of the flag to continue.\n" +
		" If you don't specify a output file, tygen will generate a yaml file called docker-compose.yml in the current directory.",
	Args: cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// 创建一个 docker-compose.yml 文件
		config := newConfig()

		// 确认要添加 Traefik 的服务名称
		if len(args) == 1 {
			service.name = args[0]
		} else {
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
			output = "docker-compose.yml"
		}
		return config.WriteConfigAs(output)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
