package models


type User struct {
	ID 	 		int      `json:"id"`
	Name 	    string   `json:"name"`
	Username 	string   `json:"username"`
	Bio 		string   `json:"bio"`
	Email 		string 	 `json:"email"`
	Password 	string 	 `json:"password"`
	Avatar		string 	 `json:"avatar_url"`
}
