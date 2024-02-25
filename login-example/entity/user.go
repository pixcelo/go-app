package entity

import (
	"bytes"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            UserID    `db:"id"`
	Email         string    `db:"email"`
	Salt          string    `db:"salt"`
	State         UserState `db:"state"`
	Password      Password  `db:"password"`
	ActivateToken string    `db:"activate_token"`
	UpdatedAt     time.Time `db:"update_at"`
	CreatedAt     time.Time `db:"created_at"`
}

type Users []*User

type UserID uint64

type Password string

func (p Password) String() string {
	return "xxxxxxxxx"
}

func (p Password) GoString() string {
	return "xxxxxxxxx"
}

type UserState string

const (
	userActive   = UserState("active")
	UserInactive = UserState("inactive")
)

// 暗号化：パスワード+ソルトをハッシュ化する
func (u *User) CreateHashedPassword(pw, salt string) (Password, error) {
	var b bytes.Buffer
	b.Write([]byte(pw))
	b.Write([]byte(salt))
	hashed, err := bcrypt.GenerateFromPassword(b.Bytes(), bcrypt.DefaultCost)
	return Password(hashed), err
}

// パスワードの検証
func (u User) Authenticate(pw string) error {
	var b bytes.Buffer
	b.Write([]byte(pw))
	b.Write([]byte(u.Salt))
	return bcrypt.CompareHashAndPassword([]byte(u.Password), b.Bytes())
}
