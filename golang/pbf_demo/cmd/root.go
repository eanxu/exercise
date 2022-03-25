package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"pbf_demo/internal/config"
	"os"
	"runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	RunE:  run,
}

var serviceName string

func Execute(sName, bTime, version string) {
	serviceName = sName
	fmt.Println("Golang Env: ", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Println("UTC build time:", bTime)
	fmt.Println("Build from gogs repo version: ", version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(ConfigFunc)
}

func ConfigFunc() {
	viper.SetConfigName(serviceName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("read config file errr", err)
	}
	viper.WatchConfig()
	if err := viper.Unmarshal(&config.C); err != nil {
		panic(fmt.Errorf("format struct err, message:\n%v", err))
	}
}
