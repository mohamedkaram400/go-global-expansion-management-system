package generic_api_response

type APIResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}