package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type authKey struct{}

func (m *Middleware) AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read the auth header
		// verify token
		claims, err := m.verifyClaimsFromAuthHeader(r)
		if err != nil {
			writeJSONError(w, err, http.StatusUnauthorized)
			return
		}

		if claims["role"] != "admin" {
			writeJSONError(w, errors.New("user is not admin"), http.StatusUnauthorized)
			return
		}

		// pass the payloadclaims down the context
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), authKey{}, claims)))
	})
}

func (m *Middleware) UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read the auth header
		// verify token
		claims, err := m.verifyClaimsFromAuthHeader(r)
		if err != nil {
			writeJSONError(w, err, http.StatusUnauthorized)
			return
		}

		if claims["role"] != "user" {
			writeJSONError(w, errors.New("user is not user"), http.StatusUnauthorized)
			return
		}

		// pass the payloadclaims down the context
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), authKey{}, claims)))
	})
}

func (m *Middleware) verifyClaimsFromAuthHeader(r *http.Request) (jwt.MapClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := fields[1]
	claims, err := m.tokenCreator.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func writeJSONError(w http.ResponseWriter, err error, statusCode int) {
	type ErrorResponse struct {
		Error string `json:"error"`
	}

	response := ErrorResponse{Error: err.Error()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}
