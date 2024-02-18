/*
Copyright © 2024 Oliver Ziegert <dccmd@pc-ziegert.dev>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/oliverziegert/dccmd-go/cmd/alias"
	"github.com/oliverziegert/dccmd-go/cmd/version"
	"github.com/oliverziegert/dccmd-go/config"
	"github.com/oliverziegert/dccmd-go/constant"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dccmd",
	Short: fmt.Sprintf("%s (%s) version %s\n", constant.LongName, constant.ShortName, constant.Version),
	Long:  ``,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(alias.AliasCmd)
	rootCmd.AddCommand(version.VersionCmd)
}

func init() {
	cobra.OnInitialize(initConfig)
	var verbose bool

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s/%s.%s)", constant.ConfigSubPath, constant.ConfigName, constant.ConfigType))
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "toggle verbose output")
	config.SetDebug(verbose)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add all subcommand pallets
	addSubcommandPalettes()
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

		// Search config in home directory with name ".dccmd-go" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/%s", home, constant.ConfigSubPath))
		viper.AddConfigPath(".")
		viper.SetConfigType(constant.ConfigType)
		viper.SetConfigName(constant.ConfigName)
	}

	viper.SetEnvPrefix(constant.EnvPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			fmt.Fprintln(os.Stderr, "config file not found, create one")
			cobra.CheckErr(config.CreateDefaultConfigFile())
		}
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
