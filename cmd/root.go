// Package cmd contains the command line interface for excalibur
package cmd

import "github.com/fatih/color"

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

func Init()  {
	Greeting()
}
