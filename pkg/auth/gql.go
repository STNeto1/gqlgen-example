package auth

import (
	"context"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	ID      uint
	Name    string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			// // Allow unauthenticated users in
			// if err != nil || c == nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			userId, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// get the user from the database
			user := getUserByID(userId)

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

func validateAndGetUserID(c *http.Cookie) (uint, error) {
	return 1, nil
}

func getUserByID(id uint) User {
	return User{
		ID:      id,
		Name:    "stneto1",
		IsAdmin: true,
	}
}
