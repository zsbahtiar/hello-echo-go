package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "hello-echo-go",
	Short: "welcome to hello echo go",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("welcome to hello-echo-go")
	},
}

func Init() {
	initCmd.AddCommand(serverCmd)

	if err := initCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
