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
package alias

import (
	"fmt"
	"github.com/oliverziegert/dccmd-go/config"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:        "set ALIAS [URL] [ACCESSKEY] [SECRETKEY]",
	Short:      "set a new alias to configuration file",
	Long:       `set a new alias to configuration file`,
	Run:        runSet,
	Aliases:    []string{"add"},
	SuggestFor: []string{"s", "a"},
	Example:    "dccmd alias set myalias https://mydomain.com myaccess",
	Args:       validateSet,
}

func validateSet(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("requires at least an alias name")
	}
	return nil
}

var (
	domain       *string
	clientId     *string
	clientSecret *string
	returnFlow   *config.ReturnFlow
	bindAddress  *string
	bindPort     *uint16
)

func runSet(cmd *cobra.Command, args []string) {
	fmt.Println("set called")

	var err error
	target := args[0]

	if cmd.Flags().NFlag() == 0 {
		domain, err = GetDomain(target)
		cobra.CheckErr(err)
		clientId, err = GetClientId(target)
		cobra.CheckErr(err)
		clientSecret, err = GetClientSecret(target)
		cobra.CheckErr(err)
		returnFlow, err = GetReturnFlow(target)
		cobra.CheckErr(err)
		if *returnFlow == config.ReturnFlowBrowser {
			bindAddress, err = GetBindAddress(target)
			cobra.CheckErr(err)
			bindPort, err = GetBindPort(target)
			cobra.CheckErr(err)
		}
	}

	alias := config.NewAlias(*domain, *clientId, *clientSecret, *returnFlow, *bindAddress, *bindPort)
	err = config.AddAlias(target, alias)
	cobra.CheckErr(err)
}

func init() {
	AliasCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	domain = setCmd.Flags().StringP("domain", "d", "dracoon.team", "Domain of the target")
	clientId = setCmd.Flags().StringP("client-id", "i", "", "Client ID of the target")
	clientSecret = setCmd.Flags().StringP("client-secret", "s", "", "Client Secret of the target")
	returnFlowS := setCmd.Flags().StringP("return-flow", "r", string(config.ReturnFlowBrowser), "Return Flow of the target")
	returnFlowP := config.ReturnFlow(*returnFlowS)
	returnFlow = &returnFlowP
	bindAddress = setCmd.Flags().StringP("bind-address", "b", "127.0.0.1", "Bind Address of the target (only necessary for return flow 'browser')")
	bindPort = setCmd.Flags().Uint16P("bind-port", "p", 1337, "Bind Port of the target (only necessary for return flow 'browser')")
}
