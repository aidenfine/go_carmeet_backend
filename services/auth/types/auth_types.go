package auth_types

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	UserId  string `json:"userId"`
}
