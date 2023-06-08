package exception

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type PanicType1 struct {
	Message string
}

type PanicType2 struct {
	Message string
}

func ExceptionAdvisorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				log.Printf("Panic occurred: %v", err)
				debug.PrintStack()

				// Handle different panic cases
				switch err.(type) {
				case PanicType1:
					// Handle PanicType1 and return HTTP status code 300
					w.WriteHeader(http.StatusMultipleChoices)
					fmt.Fprint(w, "PanicType1 occurred")
				case PanicType2:
					// Handle PanicType2 and return HTTP status code 400
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprint(w, "PanicType2 occurred")
				default:
					// Handle any other panic and return HTTP status code 500
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "Internal Server Error")
				}
			}
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
