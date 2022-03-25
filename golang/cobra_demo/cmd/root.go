/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	remote "github.com/yoyofxteam/nacos-viper-remote"
	"os"
	"time"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra_demo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("12323123213213")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initConfigFromNacos)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra_demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra_demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra_demo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initConfigFromNacos() {
	config_viper := viper.New()
	runtime_viper := config_viper
	runtime_viper.SetConfigFile("./example_config.yaml")
	_ = runtime_viper.ReadInConfig()
	var option *remote.Option
	_ = runtime_viper.Sub("yoyogo.cloud.discovery.metadata").Unmarshal(&option)

	remote.SetOptions(option)
	remote.SetOptions(&remote.Option{
		Url:         "nacos.nacos.svc.cluster.local",
		Port:        8848,
		NamespaceId: "数管",
		GroupName:   "DEFAULT_GROUP",
		Config: 	 remote.Config{ DataId: "rscb-district" },
		Auth:        &remote.Auth{
			User:      "bjsh",
			Password:  "pwd123",
		},
	})
	//localSetting := runtime_viper.AllSettings()
	remote_viper := viper.New()
	err := remote_viper.AddRemoteProvider("nacos", "localhost", "")
	remote_viper.SetConfigType("yaml")
	err = remote_viper.ReadRemoteConfig()

	if err == nil {
		config_viper = remote_viper
		fmt.Println("used remote viper")
		provider := remote.NewRemoteProvider("yaml")
		respChan := provider.WatchRemoteConfigOnChannel(config_viper)

		go func(rc <-chan bool) {
			for {
				<-rc
				fmt.Printf("remote async: %s", config_viper.GetString("yoyogo.application.name"))
			}
		}(respChan)

	}

	appName := config_viper.GetString("yoyogo.application.name")

	fmt.Println(appName)

	go func() {
		for {
			time.Sleep(time.Second * 30) // delay after each request
			appName = config_viper.GetString("yoyogo.application.name")
			fmt.Println("sync:" + appName)
		}
	}()

	onExit()
}
}
