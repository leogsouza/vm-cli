/*
Copyright © 2022 Leonardo Souza <leogsouza@gmail.com>
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		localArgs := []string{"list"}
		run, err := cmd.Flags().GetBool("run")
		if err != nil {
			return err
		}

		if !run {
			localArgs = append(localArgs, "vms")
		} else {
			localArgs = append(localArgs, "runningvms")
		}

		out, err := exec.Command("vboxmanage", localArgs...).Output()
		if err != nil {
			return err
		}
		// TODO: Parse the output
		fmt.Println("OUTPUT: ", string(out))
		strFields := strings.Fields(string(out))
		for _, str := range strFields {
			s := strings.Split(str, " ")
			fmt.Println(s[0])
		}
		return nil
	},
}

var run bool

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&run, "run", "r", false, "List only the running vms")
}
