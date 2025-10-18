package entity

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) IsValidUsername() bool {
	if len(u.Username) == 0 {
		return false
	}
	if len(u.Username) > 100 {
		return false
	}
	return true
}

func (u *User) IsActive() bool {
	if u.ID > 0 {
		return true
	}
	return false
}
