// Package cmd contains the command line interface for excalibur
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "excalibur",
	Short: "A Next Generation APT Scanner via Go",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
		Greeting()
	},
}

func Greeting() {
	color.HiGreen(`
___________                            .__    .__  ___.
\_   _____/ ___  ___   ____   _____    |  |   |__| \_ |__    __ __  _______
 |    __)_  \  \/  / _/ ___\  \__  \   |  |   |  |  | __ \  |  |  \ \_  __ \
 |        \  >    <  \  \___   / __ \_ |  |__ |  |  | \_\ \ |  |  /  |  | \/
/_______  / /__/\_ \  \___  > (____  / |____/ |__|  |___  / |____/   |__|
        \/        \/      \/       \/                   \/

                   A Next Generation APT Scanner via Go`)
}

func Execute()  {
	//Greeting()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
