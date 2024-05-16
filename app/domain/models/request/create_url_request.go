package request

type CreateUrlRequest struct {
	Url string `json:"url" validate:"required,url"`
	CustomId string `json:"custom_url" validate:"omitempty,gte=3"`
}
