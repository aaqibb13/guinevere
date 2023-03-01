package models

type Admin struct {
	Id 			string 	`json:"id,omitempty"`
	Name 		string 	`json:"name"`
	Email		string	`json:"email"`
	Role		string	`json:"role"`
	Password 	string 	`json:"password"`
	CreatedAt	int64	`json:"createdAt"`	
	UpdatedAt	int64	`json:"updatedAt"`
}

type RegisterAdminReq struct {
	Email 		string 	`json:"email"`
	Password 	string	`json:"password"`
}