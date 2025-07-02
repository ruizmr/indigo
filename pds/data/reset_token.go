package data

import (
	"time"

	"github.com/bluesky-social/indigo/models"
)

// PasswordReset stores a one-time password reset token for a user.
// The token is never persisted in plaintext; instead we store a SHA-256 hash
// of the token issued to the user. A successful reset marks the row as Used
// which prevents replay attacks.
//
// Tokens are considered invalid once Used == true OR ExpiresAt has passed.
// Cleanup of expired tokens can be performed by the application via a simple
// DELETE query (not implemented here).
//
// NOTE: We intentionally keep the model minimal for the MVP. Additional
// auditing fields (IP address, etc.) can be layered later.

type PasswordReset struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	ExpiresAt time.Time `gorm:"index"`
	Used      bool      `gorm:"index"`

	UserID    models.Uid `gorm:"index"`
	TokenHash string     `gorm:"uniqueIndex"` // hex-encoded SHA-256 digest of the raw token
}
