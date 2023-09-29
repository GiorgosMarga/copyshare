package main

import (
	"fmt"
	"net/http"
)

func (app *Application) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *Application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func (app *Application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := app.sessionManager.GetInt(r.Context(), "user")
		if id == 0 {
			app.clientError(w, http.StatusUnauthorized)
			return

		}
		exists, err := app.user.Exists(id)
		if !exists {
			app.clientError(w, http.StatusUnauthorized)
			return
		}
		if err != nil {
			app.serverError(w, fmt.Errorf("internal server error"))
			return
		}
		w.Header().Set("Cache-control", "no-store")
		next.ServeHTTP(w, r)
	})
}

// func (app *Application) authenticate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id := app.sessionManager.GetInt(r.Context(), "user")
// 		if id == 0 {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		exists, err := app.user.Exists(id)
// 		if err != nil {
// 			app.serverError(w, err)
// 			return
// 		}
// 		if exists {
// 			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
// 			r = r.WithContext(ctx)
// 		}
// 		next.ServeHTTP(w, r)

// 	})
// }

func (app *Application) addJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (app *Application) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// Add other CORS headers as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Continue processing the request
		next.ServeHTTP(w, r)
	})
}
