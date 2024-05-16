package request


type RedirectRequest struct {
	Id string `param:"id"  validate:"required,gte=3"`
}

