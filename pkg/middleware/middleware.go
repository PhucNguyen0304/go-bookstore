package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// Middleware check token
func VerifyToken(next http.HandlerFunc, jwtSecret []byte) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get Token from header Authorization
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
            return
        }

        // Cut bearer 
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        // Verify token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Kiểm tra signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Check expirer
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        if exp, ok := claims["exp"].(float64); ok {
            if time.Now().Unix() > int64(exp) {
                http.Error(w, "Token has expired", http.StatusUnauthorized)
                return
            }
        }

        // Token is accurancy, next
        next(w, r)
    }
}