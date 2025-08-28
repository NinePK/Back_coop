package controllers

import (
	"strconv"

	"coop_back/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateNotification - สร้างการแจ้งเตือนใหม่
func CreateNotification(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var notification models.Notification
	
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	// บันทึกการแจ้งเตือน
	if err := db.Create(&notification).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create notification"})
	}
	
	return c.Status(201).JSON(notification)
}

// GetNotificationsByUser - ดึงการแจ้งเตือนตาม User ID
func GetNotificationsByUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Params("userId")
	
	var notifications []models.Notification
	
	// ดึงการแจ้งเตือนพร้อมข้อมูลผู้ส่ง เรียงจากล่าสุด
	if err := db.Preload("Sender").
		Where("recipient_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch notifications"})
	}
	
	return c.JSON(notifications)
}

// GetUnreadNotifications - ดึงการแจ้งเตือนที่ยังไม่ได้อ่าน
func GetUnreadNotifications(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Params("userId")
	
	var notifications []models.Notification
	
	if err := db.Preload("Sender").
		Where("recipient_id = ? AND is_read = ?", userID, false).
		Order("created_at DESC").
		Find(&notifications).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch unread notifications"})
	}
	
	return c.JSON(notifications)
}

// GetUnreadCount - ดึงจำนวนการแจ้งเตือนที่ยังไม่ได้อ่าน
func GetUnreadCount(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Params("userId")
	
	var count int64
	
	if err := db.Model(&models.Notification{}).
		Where("recipient_id = ? AND is_read = ?", userID, false).
		Count(&count).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to count unread notifications"})
	}
	
	return c.JSON(fiber.Map{"unreadCount": count})
}

// MarkAsRead - ทำเครื่องหมายว่าอ่านแล้ว
func MarkAsRead(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	notificationID := c.Params("id")
	
	id, err := strconv.ParseInt(notificationID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid notification ID"})
	}
	
	// อัปเดตสถานะเป็นอ่านแล้ว
	if err := db.Model(&models.Notification{}).
		Where("id = ?", id).
		Update("is_read", true).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark as read"})
	}
	
	return c.JSON(fiber.Map{"message": "Notification marked as read"})
}

// MarkAllAsRead - ทำเครื่องหมายว่าอ่านแล้วทั้งหมด
func MarkAllAsRead(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Params("userId")
	
	// อัปเดตการแจ้งเตือนทั้งหมดของ user นี้เป็นอ่านแล้ว
	if err := db.Model(&models.Notification{}).
		Where("recipient_id = ? AND is_read = ?", userID, false).
		Update("is_read", true).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark all as read"})
	}
	
	return c.JSON(fiber.Map{"message": "All notifications marked as read"})
}

// DeleteNotification - ลบการแจ้งเตือน
func DeleteNotification(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	notificationID := c.Params("id")
	
	id, err := strconv.ParseInt(notificationID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid notification ID"})
	}
	
	if err := db.Delete(&models.Notification{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete notification"})
	}
	
	return c.JSON(fiber.Map{"message": "Notification deleted successfully"})
}

// CreateDocumentNotification - ฟังก์ชันช่วยสำหรับสร้างการแจ้งเตือนเมื่อยื่นเอกสาร
func CreateDocumentNotification(db *gorm.DB, studentID, advisorID int64, documentType, documentTitle string, relatedID int64) error {
	// หาข้อมูลนิสิต
	var student models.User
	if err := db.Preload("Major.Faculty").First(&student, studentID).Error; err != nil {
		return err
	}
	
	// สร้างข้อความแจ้งเตือน
	var title, message string
	switch documentType {
	case "coop07":
		title = "เอกสาร COOP-07: โครงร่างรายงานใหม่"
		message = "นิสิต " + student.Fname + " " + student.Sname + " (" + student.Username + ") ได้ส่งโครงร่างรายงานการปฏิบัติงานแล้ว กรุณาตรวจสอบและให้คำแนะนำ"
	case "coop10":
		title = "เอกสาร COOP-10: ยืนยันส่งรายงานใหม่"
		message = "นิสิต " + student.Fname + " " + student.Sname + " (" + student.Username + ") ได้ยืนยันส่งรายงาน Work Term Report แล้ว หัวข้อ: \"" + documentTitle + "\""
	case "coop11":
		title = "เอกสาร COOP-11: รายละเอียดการปฏิบัติงานใหม่"
		message = "นิสิต " + student.Fname + " " + student.Sname + " (" + student.Username + ") ได้ส่งรายละเอียดการปฏิบัติงานแล้ว กรุณาตรวจสอบข้อมูล"
	case "coop12":
		title = "เอกสาร COOP-12: การประเมินตนเองใหม่"
		message = "นิสิต " + student.Fname + " " + student.Sname + " (" + student.Username + ") ได้ทำการประเมินตนเองแล้ว กรุณาตรวจสอบผลการประเมิน"
	default:
		title = "เอกสารสหกิจศึกษาใหม่"
		message = "นิสิต " + student.Fname + " " + student.Sname + " (" + student.Username + ") ได้ส่งเอกสารใหม่แล้ว"
	}
	
	// สร้างการแจ้งเตือน
	notification := models.Notification{
		RecipientID:  advisorID,
		SenderID:     studentID,
		Type:         "document_submitted",
		Title:        title,
		Message:      message,
		DocumentType: documentType,
		RelatedID:    relatedID,
		Priority:     "normal",
		IsRead:       false,
	}
	
	return db.Create(&notification).Error
}