package request

type EditShortUrlRequest struct {
	ShortUrl string `json:"short_url" validate:"required,url"`
	NewUrl string `json:"new_value" validate:"required,gte=3"`
}

type EditRedirectURLRequest struct {
	ShortUrl string `json:"short_url" validate:"required,url"`
	RedirectUrl string `json:"redirect_url" validate:"required,url"`
}

type EditUrlStatusRequest struct {
	ShortUrl string `json:"short_url" validate:"required,url"`
	IsActive *bool `json:"is_active" validate:"required,boolean"`
}
