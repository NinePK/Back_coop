package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CustomDomainMiddleware2(c *fiber.Ctx) error {
	// log.Println("c.getHost: " + c.Get("Host"))
	// log.Println("client: " + access_domain)

	// access_domain := c.Get("Host")

	// forwardedProto := c.Get("X-Forwarded-Proto")
	// clientIP := c.Get("X-Forwarded-For")

	// // if clientIP == "" {
	// // 	clientIP = c.Get("X-Real-IP") // Check X-Real-IP header
	// // }
	// if clientIP == "" {
	// 	clientIP = c.IP() // Fallback to c.IP() if no headers are available
	// }

	// if forwardedProto == "" {
	// 	forwardedProto = "http"
	// }

	// incoming_domain := forwardedProto + "://" + clientIP

	log.Println("Logging .....")
	origin := c.Get("Origin")
	log.Println(origin, " ---> ", c.OriginalURL())

	if origin != "https://coop.ict.up.ac.th" && origin != "http://localhost:6007" {
		return c.Status(fiber.StatusForbidden).SendString("Access denied")
	}
	return c.Next()
}

// Custom middleware to check the domain
func CustomDomainMiddleware(c *fiber.Ctx) error {
	// log.Println("c.getHost: " + c.Get("Host"))
	// log.Println("client: " + access_domain)

	// access_domain := c.Get("Host")

	forwardedProto := c.Get("X-Forwarded-Proto")
	clientIP := c.Get("X-Forwarded-For")

	// if clientIP == "" {
	// 	clientIP = c.Get("X-Real-IP") // Check X-Real-IP header
	// }
	if clientIP == "" {
		clientIP = c.IP() // Fallback to c.IP() if no headers are available
	}

	if forwardedProto == "" {
		forwardedProto = "http"
	}

	// incoming_domain := forwardedProto + "://" + clientIP

	origin := c.Get("Origin")
	// log.Println(origin, " ---> ", access_domain, " ---> ", c.OriginalURL(), " ---> ", incoming_domain)

	// return c.Next()
	if origin == os.Getenv("domain_allow") || origin == os.Getenv("local_domain") || clientIP == "127.0.0.1" || clientIP == "171.4.238.185" {
		return c.Next()
	} else {
		log.Println("Access denied from clientIP: " + origin + " and clientIP: " + clientIP)
		return c.Status(fiber.StatusForbidden).SendString("Access denied")
	}
}

func getIP(c *fiber.Ctx) error {

	forwardedProto := c.Get("X-Forwarded-Proto")
	clientIP := c.Get("X-Forwarded-For")

	if clientIP == "" {
		clientIP = c.Get("X-Real-IP") // Check X-Real-IP header
	}
	if clientIP == "" {
		clientIP = c.IP() // Fallback to c.IP() if no headers are available
	}

	if forwardedProto == "" {
		forwardedProto = "http"
	}

	access_domain := forwardedProto + "://" + clientIP

	return c.Status(fiber.StatusAccepted).SendString(access_domain)

	// log.Println("client: " + access_domain)

	// if access_domain == os.Getenv("domain_allow") {
	// 	return c.Next()
	// } else {
	// 	return c.Status(fiber.StatusForbidden).SendString("Access denied")
	// }
}
