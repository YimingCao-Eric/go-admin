package util

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// SecretKey is the signing key used for JWT token generation and validation
// WARNING: This is hardcoded for development - in production, use environment variables
// The secret key must be kept secure and never exposed in version control
const SecretKey = "JWT_SECRET"

// GenerateJWT creates a new JWT token for a given user ID
// Uses HS256 signing algorithm and sets token expiration to 24 hours
// The user ID is stored in the "iss" (issuer) claim of the token
func GenerateJWT(issuer string) (string, error) {
	// Create JWT with claims containing user ID and expiration time
	// Expiration set to 24 hours from current time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign token with secret key using HS256 algorithm
	return claims.SignedString([]byte(SecretKey))
}

// ParseJWT validates a JWT token and extracts the user ID from the issuer claim
// Validates token signature, expiration, and other standard claims
// Returns error if token is invalid, expired, or signature verification fails
func ParseJWT(cookie string) (string, error) {
	// Parse and validate JWT token with standard claims
	// Validation function provides secret key for signature verification
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Return secret key for signature verification
		return []byte(SecretKey), nil
	})

	// Check if token parsing or validation failed
	if err != nil || !token.Valid {
		return "", err
	}

	// Extract and type-assert claims
	claims := token.Claims.(*jwt.StandardClaims)

	// Return user ID from issuer claim
	return claims.Issuer, nil
}
