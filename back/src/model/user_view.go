package model

import "time"

type UserView struct {
	Email       string    `json:"email,omitempty"`
	Name        string    `json:"name,omitempty"`
	LoginType   string    `json:"login_type,omitempty"`
	PhotoURL    string    `json:"photoUrl,omitempty"`
	FriendCount int       `json:"friend_count,omitempty"`
	Expires     time.Time `json:"expires,omitempty"`
}

type FacebookResponse struct {
	Email   string `facebook:"email"`
	Name    string `facebook:"name"`
	Picture struct {
		Data struct {
			URL string `facebook:"url"`
		} `facebook:"data"`
	} `facebook:"picture"`
	ID string `facebook:"id"`
}
