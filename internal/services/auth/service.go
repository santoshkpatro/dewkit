package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"dewkit/config"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func NewService() *Service {
	db := config.DB
	return &Service{DB: db}
}

// GenerateSalt creates a random per-user salt
func GenerateSalt() (string, error) {
	b := make([]byte, 16) // 128-bit salt
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// HashPassword hashes password using salt + secret (pepper)
func HashPassword(password, salt string) string {
	secret := config.GetEnv("SECRET_KEY")
	data := password + ":" + salt + ":" + secret
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// VerifyPassword checks password against stored hash + salt
func VerifyPassword(password, salt, storedHash string) bool {
	return HashPassword(password, salt) == storedHash
}

// CreateSuperuser creates a superuser if email does not already exist
func (s *Service) CreateSuperuser(
	fullName string,
	email string,
	password string,
) error {

	// 1. Check if email already exists
	var exists bool
	err := s.DB.Get(
		&exists,
		`SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`,
		email,
	)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists with the given email address")
	}

	// 2. Generate salt
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}

	// 3. Hash password
	passwordHash := HashPassword(password, salt)

	// 4. Insert superuser
	_, err = s.DB.Exec(
		`INSERT INTO users (full_name, email, password_hash, password_salt, role)
		 VALUES ($1, $2, $3, $4, 'superuser')`,
		fullName,
		email,
		passwordHash,
		salt,
	)

	return err
}

// Authenticate verifies user credentials
func (s *Service) Authenticate(email string, password string) error {
	var storedHash, storedSalt string

	err := s.DB.Get(
		&storedHash,
		`SELECT password_hash, password_salt FROM users WHERE email = $1`,
		email,
	)
	if err != nil {
		return err
	}

	if !VerifyPassword(password, storedSalt, storedHash) {
		return errors.New("invalid email or password")
	}

	return nil
}
