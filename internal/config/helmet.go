package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

func HelmetConfig() fiber.Handler {
	return helmet.New(helmet.Config{
		// Enable Content Security Policy (CSP) to prevent XSS attacks
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; frame-ancestors 'none'; base-uri 'self'; form-action 'self';",
		// Enable Cross-Origin Resource Sharing (CORS) to allow cross-origin requests
		CrossOriginEmbedderPolicy: "require-corp",
		// Enable Cross-Origin Resource Policy (CORP) to prevent cross-origin resource sharing
		CrossOriginResourcePolicy: "same-origin",
		// Enable Referrer Policy to control the amount of referrer information sent with requests
		ReferrerPolicy: "no-referrer",
		// Enable X-DNS-Prefetch-Control to control DNS prefetching
		XDNSPrefetchControl: "on",
	})
}
