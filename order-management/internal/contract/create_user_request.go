package contract

type CreateUserRequest struct {
	Id          string
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	CreatedAt   string
	Role        string
}
