package cmd

import "github.com/spf13/cobra"

/*
In this file we will center all of our cli commands
add
list
remove
*/
func NewRootCmd() *cobra.Command {
	/*the use field represents what well type in the terminal like task-tracker
	or it could be tt*/
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "CLI for managing tasks",
	}

	cmd.AddCommand(NewAddTaskCmd())
	return cmd
}
