package dto

type Success struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Failed struct {
	Success bool   `json:"success"`
	Error   string `json:"message"`
}
