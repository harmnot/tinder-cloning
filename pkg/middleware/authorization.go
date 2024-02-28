package middleware

import (
	"context"
	"net/http"
	"strings"
	"tinder-cloning/pkg/util"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from the Authorization header
		// format: Bearer {token}
		tokens := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tokens) != 2 {
			util.RenderJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		accountData, err := util.VerifyTokenJWT(tokens[1])
		if err != nil {
			util.RenderJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// If everything is OK, proceed with the next handler
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "accountData", accountData)))
	})
}
