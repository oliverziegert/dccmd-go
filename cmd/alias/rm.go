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
	"github.com/oliverziegert/dccmd-go/utils"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:        "rm ALIAS",
	Short:      "remove an alias from configuration file",
	Long:       `remove an alias from configuration file`,
	Run:        runRm,
	Aliases:    []string{"del"},
	SuggestFor: []string{"r", "d"},
	Example:    "dccmd rm myalias",
	Args:       validateRm,
}

func validateRm(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("an alias is required as argument")
	}
	target := args[0]
	if !utils.IsValidTarget(target) {
		return fmt.Errorf("invalid target specified")
	}
	return nil
}

func runRm(cmd *cobra.Command, args []string) {
	fmt.Println("rm called")
	target := args[0]
	err := config.RemoveAlias(target)
	cobra.CheckErr(err)
}

func init() {
	AliasCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
