package dto

type ApiResponse struct {
	Data       any    `json:"data"`
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode uint   `json:"statusCode"`
}
