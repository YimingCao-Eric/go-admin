package util

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// SecretKey is the signing key used for JWT token generation and validation
// In production, this should be stored in environment variables, not hardcoded
const SecretKey = "JWT_SECRET"

// GenerateJWT creates a new JWT token for a given user ID (issuer)
// Parameters:
//   - issur: the user ID string that will be set as the token issuer
//
// Returns:
//   - string: the signed JWT token
//   - error: any error that occurred during token generation
func GenerateJWT(issur string) (string, error) {
	// Create JWT claims with user ID and expiration time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issur,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with secret key using HS256 algorithm
	return claims.SignedString([]byte(SecretKey))
}

// ParseJWT validates a JWT token and extracts the issuer (user ID)
// Parameters:
//   - cookie: the JWT token string from the cookie
//
// Returns:
//   - string: the issuer (user ID) from the token claims
//   - error: validation error or parsing error
func ParseJWT(cookie string) (string, error) {
	// Parse and validate JWT token with custom claims
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // Return secret key for validation
	})

	// Check if token is invalid or parsing failed
	if err != nil || !token.Valid {
		return "", err
	}

	// Extract claims from token
	claims := token.Claims.(*jwt.StandardClaims)

	// Return the issuer (user ID) from the claims
	return claims.Issuer, nil
}
