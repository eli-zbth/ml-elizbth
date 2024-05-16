package parse

import (
	"ml-elizabeth/app/shared/utils/constants"
	"net/url"

	gonanoid "github.com/matoous/go-nanoid/v2"
)


func DeleteDomain(longUrl string) (string) {
	u, _ := url.Parse(longUrl)

	path := u.RequestURI()
	path = path[1:]
	return path
}

func GenerateId() (string,error) {
    id, err := gonanoid.New(constants.ShortUrlLength)
	return id,err
}




