package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	//"github.com/go-git/go-git/v5"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
	Use: 	"clone [url to clone]",
	Short: 	"Clone a repository",
	Run: doClone,
	Args: cobra.RangeArgs(1,2),
}

func doClone(cmd *cobra.Command, args []string) {
	var (
		local string
	)
	if len(args) == 2 {
		local = args[1]
	} else {
		parts := strings.Split(args[0],"/")
		local = parts[len(parts)-1]
	}

	
	fmt.Printf("Clonging %s into %s\n", args[0], local)

}