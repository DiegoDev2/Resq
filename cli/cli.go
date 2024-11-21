package cli

import (
	"fmt"

	resq "github.com/DiegoDev2/resq/src"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "resq",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "get [url]",
			Short: "Get method",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				// Mostrar el resultado de Get
				result := resq.Get(args[0])
				fmt.Println(result)
			},
		},
		&cobra.Command{
			Use:   "post [url] [data]",
			Short: "Post method",
			Args:  cobra.ExactArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				// Mostrar el resultado de Post
				result := resq.Post(args[0], args[1])
				fmt.Println(result)
			},
		},
		&cobra.Command{
			Use:   "put [url] [data]",
			Short: "Put method",
			Args:  cobra.ExactArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				// Mostrar el resultado de Put
				result := resq.Put(args[0], args[1])
				fmt.Println(result)
			},
		},
		&cobra.Command{
			Use:   "delete [url]",
			Short: "Delete method",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				// Mostrar el resultado de Delete
				result := resq.Delete(args[0])
				fmt.Println(result)
			},
		},
	)

	return rootCmd
}
