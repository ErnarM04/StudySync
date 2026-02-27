package services

import (
    "errors"
    "time"

    "github.com/kadyrbayev2005/studysync/internal/utils"
    "golang.org/x/crypto/bcrypt"

    "github.com/golang-jwt/jwt/v5"
)

// jwtKey holds the HMAC secret used to sign and verify access tokens (loaded in InitJWT).
var jwtKey []byte

const (
    RoleAdmin = "admin"
    RoleUser  = "user"
)

// Claims is the JWT payload: who the caller is (UserID, Role) plus standard expiry/issued times.
type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// InitJWT reads JWT_SECRET from the environment (or a dev default) into memory for HS256 signing.
func InitJWT() {
    secret := utils.GetEnv("JWT_SECRET", "your-secure-secret-key")
    jwtKey = []byte(secret)
}

// GenerateJWT builds a signed HS256 token valid for 24 hours carrying user id and role.
func GenerateJWT(userID uint, role string) (string, error) {
    expiration := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expiration),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ParseJWT validates the signature and expiry, then returns the embedded claims or an error.
func ParseJWT(tokenStr string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("invalid token")
}

// HashPassword returns a bcrypt hash suitable for storing on User.PasswordHash.
func HashPassword(pw string) string {
    b, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    return string(b)
}

// CheckPasswordHash compares a plaintext password with a stored bcrypt hash.
func CheckPasswordHash(pw, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
    return err == nil
}