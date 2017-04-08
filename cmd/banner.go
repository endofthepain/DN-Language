package cmd

import (
	"os"

	"github.com/fatih/color"
)

var (
	banner = `
 			____  __.  _____      _____ ____________________________________________      __                       
			|    |/ _| /  _  \    /     \\______   \______   \_   _____/\__    ___/  \    /  \_____ _______   ____  
			|      <  /  /_\  \  /  \ /  \|     ___/|       _/|    __)_   |    |  \   \/\/   /\__  \\_  __ \_/ __ \ 
			|    |  \/    |    \/    Y    \    |    |    |   \|        \  |    |   \        /  / __ \|  | \/\  ___/ 
			|____|__ \____|__  /\____|__  /____|    |____|_  /_______  /  |____|    \__/\  /  (____  /__|    \___  >
        			\/       \/         \/                 \/        \/                  \/        \/            \/ `
	delimiter = `
    ========================================================================================================================
    `
)

// Print banner
func PrintBanner() {
	color.New(color.FgCyan).Add(color.Bold).Println(delimiter)
	color.New(color.FgRed).Add(color.Bold).Println(banner)
	color.New(color.FgCyan).Add(color.Bold).Println(delimiter)
}

// Print script usage
func Usage(errmsg string) {
	color.New(color.FgWhite).Add(color.Bold).Printf(
		"%s\n\n"+
			"usage: %s\n"+
			"       %s decrypt key\n"+
			"      Will try to decrypt your files using the given key\n",
		errmsg, os.Args[0], os.Args[0])
	os.Exit(2)
}
