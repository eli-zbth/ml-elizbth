package parse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestDeleteDomain(t*testing.T){

    t.Run("Should delete domain from url with http", func(t *testing.T) {

        url := "http://www.domain.cl/params_with_http"
        result := DeleteDomain(url)
        assert.Equal(t,"params_with_http",result)
	})

    t.Run("Should delete domain from url with https", func(t *testing.T) {

        url := "https://www.domain.cl/params_with_https"
        result := DeleteDomain(url)
        assert.Equal(t,"params_with_https",result)
	})

    t.Run("Should delete domain from localhost", func(t *testing.T) {
        url := "http://localhost:8080/paras"
        result := DeleteDomain(url)
        assert.Equal(t,result,"paras")
	})

    t.Run("Should delete domain from url with query params", func(t *testing.T) {
        url := "https://www.domain.com/page?key1=value1&key2=value2"
        result := DeleteDomain(url)
        assert.Equal(t,result,"page?key1=value1&key2=value2")
	})

    t.Run("Should delete domain from url with path params", func(t *testing.T) {
        url := "https://www.domain.com/paramname/paramvalue"
        result := DeleteDomain(url)
        assert.Equal(t,result,"paramname/paramvalue")
	})

}
