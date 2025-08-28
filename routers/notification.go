package routers

import (
	"coop_back/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterNotificationRoutes(router fiber.Router) {
	notification := router.Group("/notification")
	
	// สร้างการแจ้งเตือนใหม่
	notification.Post("/", controllers.CreateNotification)
	
	// ดึงการแจ้งเตือนตาม User ID
	notification.Get("/user/:userId", controllers.GetNotificationsByUser)
	
	// ดึงการแจ้งเตือนที่ยังไม่ได้อ่าน
	notification.Get("/user/:userId/unread", controllers.GetUnreadNotifications)
	
	// ดึงจำนวนการแจ้งเตือนที่ยังไม่ได้อ่าน
	notification.Get("/user/:userId/unread/count", controllers.GetUnreadCount)
	
	// ทำเครื่องหมายว่าอ่านแล้ว
	notification.Put("/:id/read", controllers.MarkAsRead)
	
	// ทำเครื่องหมายว่าอ่านแล้วทั้งหมด
	notification.Put("/user/:userId/read-all", controllers.MarkAllAsRead)
	
	// ลบการแจ้งเตือน
	notification.Delete("/:id", controllers.DeleteNotification)
}