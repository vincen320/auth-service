package userservice

type GetResponses struct {
	Status  int                           `json:"status,omitempty"`
	Message string                        `json:"message,omitempty"`
	Data    UserDataGetResponseAppService `json:"data,omitempty"`
}
