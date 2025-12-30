package pkg

import "github.com/gorilla/handlers"

func Cors() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Location"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	return headersOk, methodsOk, originsOk
}
