package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tygen",
	Short: "Tygen is a tool for generating traefik yaml files.",
	Long:  `Tygen is a free and open source tool for generating traefik yaml files. You can also use it to modify your already existing traefik yaml files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("no subcommand specified")
	},
}

var service Service = *newService() // Traefik 服务配置信息
var output string = ""              // 输出文件路径

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output file")
	rootCmd.PersistentFlags().StringVarP(&service.name, "service", "s", newService().name, "service name")
	rootCmd.PersistentFlags().StringVarP(&service.network, "network", "n", newService().network, "network name")
	rootCmd.PersistentFlags().IntVarP(&service.port, "port", "p", newService().port, "internal port for the service")
	var address string = ""
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "example.com", "Traefik host address")
	if address != "" {
		service.host = fmt.Sprintf("Host(`%s`)", address)
	}
	rootCmd.PersistentFlags().StringVarP(&service.host, "rule", "r", newService().host, "Traefik host rule")
	rootCmd.PersistentFlags().StringVarP(&service.image, "image", "i", newService().image, "docker image")
	rootCmd.PersistentFlags().StringVarP(&service.entrypoints, "entrypoints", "e", newService().entrypoints, "entrypoints for the Traefik")
	rootCmd.PersistentFlags().BoolVarP(&service.tls, "tls", "t", newService().tls, "enable tls for the service in Traefik")
}

func Execute() {
	rootCmd.Execute()
}
