package response

type UserResponse struct {
	ID       	uint   	`json:"id"`
	Name 		string 	`json:"name"`
	Email    	string 	`json:"email"`
	Password 	string 	`json:"password"`
	UpdatedAt 	string 	`json:"updated_at"`
	CreatedAt 	string 	`json:"created_at"`
}