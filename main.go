package main

import (
	"Beckend_Student2025/cmd"
	config "Beckend_Student2025/configs"
	admin "Beckend_Student2025/controllers/admins"
	"Beckend_Student2025/controllers/auth"
	"Beckend_Student2025/controllers/checkins"
	"Beckend_Student2025/controllers/events"
	"Beckend_Student2025/controllers/staffs"
	"Beckend_Student2025/controllers/tickets"
	"Beckend_Student2025/controllers/users"
	"Beckend_Student2025/middlewares"
	"Beckend_Student2025/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("EMAIL_USER") == "" || os.Getenv("EMAIL_PASSWORD") == "" {
		log.Fatal("SMTP credentials not set in .env")
	}

	config.Database()

	if err := command(); err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	err = utils.SendEmail("phloem.contact@gmail.com", "Test Subject", "This is a test email")
	if err != nil {
		fmt.Println("❌ Failed to send email:", err)
	} else {
		fmt.Println("✅ Email sent successfully!")
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

	md := middlewares.AuthMiddleware()

	r.POST("/user/forgot", users.ForgotStudentID)

	r.POST("/user/login", auth.LoginUser)
	r.POST("/staff/login", auth.LoginStaff)
	r.POST("/admin/login", auth.LoginAdmin)

	r.GET("/staff/info", staffs.GetInfoStaff)
	r.GET("/admin/info", admin.GetInfoAdmin)
	r.GET("/user/info", users.GetInfoUser)

	r.GET("/user/:id", users.GetUserByID)
	r.GET("/user/list", users.UserList)
	r.POST("/user/create", users.CreateUser)
	r.DELETE("/user/delete/:id", users.DeleteUser)

	r.GET("/staff/:id", staffs.GetStaffByID)
	r.GET("/staff/list", staffs.StaffList)
	r.POST("/staff/create", staffs.CreateStaff)
	r.DELETE("/staff/delete/:id", staffs.DeleteStaff)
	r.PATCH("/staff/update/:id", staffs.UpdateStaff)

	r.GET("/admin/:id", admin.GetAdminByID)
	r.POST("/admin/create", admin.CreateAdmin)
	r.DELETE("/admin/delete/:id", admin.DeleteAdmin)
	r.PATCH("/admin/update/:id", admin.UpdateAdmin)

	r.POST("/event/create", events.CreateEvent)
	r.PATCH("/event/update/:id", events.UpdateEvent)
	r.GET("/event/:id", events.GetEventByID)
	r.GET("/event/list", events.EventList)

	r.GET("/ticket/list", md, tickets.TicketList)
	r.POST("ticket/create", md, tickets.CreateTicket)

	r.POST("/checkin/create", md, checkins.CheckinCreate)

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
