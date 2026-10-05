package types

import ("time")

type User struct {
	ID    	 	  string `json:"id"`
	Name  	 	  string    `json:"name"`
	Username 	  string	`json:"username"`
	Email    	  string    `json:"email"`
	Pass     	  string    `json:"password"`
	RefreshToken  string	`json:"refresh_token"`
	Avatar		  string 	`json:"avatar_url,omitempty"`
	Bio 		  string 	`json:"bio,omitempty"`
	IsOnline 	  bool 		`json:"is_online"`
	IsActive 	  bool 		`json:"is_active"`
	TwoFactor 	  bool 		`json:"two_factor_enabled"`
	Otp			  bool 		`json:"otp_enabled"`
	OtpSecret 	  string 	`json:"otp_secret"`
	Settings 	  string 	`json:"settings"`
	LastSeen	  time.Time `json:"last_seen_at"`
	CreatedAt 	  time.Time `json:"created_at"`
}


type PublicUser struct {
	Name 	 string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar 	 string `json:"avatar_url,omitempty"`
	Bio      string `json:"bio,omitempty"`
	LastSeen time.Time `json:"last_seen_at"`
	IsOnline bool   `json:"is_online"`
}