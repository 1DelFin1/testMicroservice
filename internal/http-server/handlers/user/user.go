package user

import "time"

type Data struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
