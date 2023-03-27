package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/zsbahtiar/hello-echo-go/internal/route"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	server := echo.New()
	r := route.NewRouter(server)
	r.Register()

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
