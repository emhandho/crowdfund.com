package user

import "time"

type User struct {
	ID             int       `db:"id"`
	Name           string    `db:"name"`
	Occupation     string    `db:"occupation"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	AvatarFileName string    `db:"avatar_file_name"`
	Role           string    `db:"role"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
