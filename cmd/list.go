/*
Copyright © 2022 Leonardo Souza <leogsouza@gmail.com>
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the VMs created locally",
	Long: `This command lists the VMs. If pass the flag -r
	will list only the running VMs`,
	RunE: func(cmd *cobra.Command, args []string) error {

		localArgs := []string{"list"}

		if !run {
			localArgs = append(localArgs, "vms")
		} else {
			localArgs = append(localArgs, "runningvms")
		}

		out, err := exec.Command("vboxmanage", localArgs...).Output()
		if err != nil {
			return err
		}

		r := regexp.MustCompile(`\S+`)
		outStr := string(out)
		arrStr := r.FindAllString(outStr, -1)
		vms := []string{}
		for _, str := range arrStr {
			if str[0] == '"' {
				vms = append(vms, str)
			}
		}
		// TODO: Display data with more details
		fmt.Println("VMS", vms)

		return nil
	},
}

var run bool

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&run, "run", "r", false, "List only the running vms")
}
