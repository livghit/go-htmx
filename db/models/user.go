package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/livghit/go-htmx/db"
)

// User represents a registered application user.
type User struct {
	ID        int64
	Username  string
	Password  string // bcrypt hash — never expose in JSON/templates
	Email     string
	CreatedAt time.Time
}

// ErrNotFound is returned when a user lookup finds no rows.
var ErrNotFound = errors.New("user not found")

// GetByID fetches a user by primary key.
func GetByID(id int64) (*User, error) {
	u := &User{}
	err := db.DB.QueryRow(
		`SELECT id, username, password, email, created_at FROM users WHERE id = ?`, id,
	).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return u, nil
}

// GetByUsername fetches a user by username (case-sensitive).
func GetByUsername(username string) (*User, error) {
	u := &User{}
	err := db.DB.QueryRow(
		`SELECT id, username, password, email, created_at FROM users WHERE username = ?`, username,
	).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	return u, nil
}

// GetByEmail fetches a user by email address.
func GetByEmail(email string) (*User, error) {
	u := &User{}
	err := db.DB.QueryRow(
		`SELECT id, username, password, email, created_at FROM users WHERE email = ?`, email,
	).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return u, nil
}

// Create inserts a new user and returns the created record.
// hashedPw must already be a bcrypt hash — never pass a plaintext password.
func Create(username, email, hashedPw string) (*User, error) {
	result, err := db.DB.Exec(
		`INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`,
		username, email, hashedPw, time.Now(),
	)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get last insert id: %w", err)
	}
	return GetByID(id)
}

// Update saves changes to username and email for an existing user.
func (u *User) Update() error {
	_, err := db.DB.Exec(
		`UPDATE users SET username = ?, email = ? WHERE id = ?`,
		u.Username, u.Email, u.ID,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}

// Delete removes the user from the database.
func (u *User) Delete() error {
	_, err := db.DB.Exec(`DELETE FROM users WHERE id = ?`, u.ID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
