package middlewares

import "net/http"

// securityHeaders is a middleware that wraps your handlers.
// Its job is to tell the Browser (Frontend) exactly how to behave
// to keep the user safe from common web attacks.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. X-DNS-Prefetch-Control
		// WHY: Stops browsers from "guessing" which links a user might click.
		// If on, the browser pre-resolves IP addresses of all links, which can leak
		// user privacy. "off" disables this.
		w.Header().Set("X-DNS-Prefetch-Control", "off")

		// 2. X-XSS-Protection (Standard key: "X-XSS-Protection")
		// WHY: Tells the browser to stop the page from loading if it detects
		// a Cross-Site Scripting (XSS) attack. "1; mode=block" is the strongest setting.
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// 3. X-Content-Type-Options (Standard key: "X-Content-Type-Options")
		// WHY: Prevents "MIME-sniffing". Without this, a browser might try to
		// execute a simple .txt file as a .js script if it looks like code.
		// "nosniff" forces the browser to trust the Content-Type header we send.
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// 4. Strict-Transport-Security (HSTS)
		// WHY: Forces the browser to ONLY use HTTPS for this site for the next year.
		// It prevents "Man-in-the-Middle" attacks where someone tries to
		// downgrade your connection to plain HTTP.
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// 5. Content-Security-Policy (CSP)
		// WHY: The most powerful header. It tells the browser exactly which
		// sources (scripts, images, styles) are allowed to load.
		// "default-src 'self'" means: "Only load stuff from my own domain."
		w.Header().Set("Content-Security-Policy", "default-src 'self'")

		// 6. Referrer-Policy
		// WHY: When a user clicks a link to leave your site, this controls
		// how much info is sent to the next site about where they came from.
		w.Header().Set("Referrer-Policy", "no-referrer")

		w.Header().Set("X-Powered-By", "Django") // mentioning incorrect headers will mislead any attacker with ill intentions

		w.Header().Set("Server", "")
		w.Header().Set("Ix-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone= () ")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// basic middleware skeleton
// func securityHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

// 	})
// }
