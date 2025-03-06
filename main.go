package main

import (
	"Beckend_Student2025/cmd"
	config "Beckend_Student2025/configs"
	admin "Beckend_Student2025/controllers/admins"
	"Beckend_Student2025/controllers/staffs"
	"Beckend_Student2025/controllers/users"
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

	// Info
	r.GET("/staff/:id", staffs.GetInfoStaff)
	r.GET("/admin/:id", admin.GetInfoAdmin)
	r.GET("/user/:id", users.GetInfoUser)

	// User
	r.GET("/user/:id", users.GetUserByID)
	r.POST("/user/create", users.CreateUser)
	r.DELETE("/user/:id", users.DeleteUser)

	// Staff
	r.GET("/staff/:id", staffs.GetStaffByID)
	r.POST("/staff/create", staffs.CreateStaff)
	r.DELETE("/staff/:id", staffs.DeleteStaff)
	r.PATCH("/staff/:id", staffs.UpdateStaff)

	// Admin
	r.GET("/admin/:id", admin.GetAdminByID)
	r.POST("/admin/create", admin.CreateAdmin)
	r.DELETE("/admin/:id", admin.DeleteAdmin)
	r.PATCH("/admin/:id", admin.UpdateAdmin)

	r.Run()

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
