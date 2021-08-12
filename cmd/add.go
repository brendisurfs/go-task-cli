package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use: "add", 
	Short: "Adds a task to your task list",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}


// WHAT IS AN INIT FUNC? 

/*INIT is a function that can be run before main, almost like a precursor, rather than just having stuff inside the main .
	this is good for setting stuff up, accessing env variables, other stuff.
	Its good to know, just USE IT SPARINGLY. 
	It is bloaty and should only be used in a case like this. 
	save your package.

*/
func init() {
	RootCmd.AddCommand(&addCmd)
}