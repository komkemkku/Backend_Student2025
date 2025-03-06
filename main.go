package main

import (
	"Beckend_Student2025/cmd"
	config "Beckend_Student2025/configs"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func main() {
	config.Database()
	if err := command(); err != nil {
		log.Fatalf("Error runing command :%v", err)
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
	}))

	// md := middlewares.AuthMiddleware()

}

func command() error {
	cmda := &cobra.Command{
		Use:  "app",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmda.AddCommand(cmd.Migrate())

	return cmda.Execute()
}
