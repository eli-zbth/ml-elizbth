package manage

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"urlShortenerApi/app/domain/models/response"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUrl(t *testing.T) {
	// customer := &models.Url{

	// 	url: "http://fakeurl.com",
	// }

	e := echo.New()
	// e.Validator, _ = validations.NewCustomValidator(validator.New())

	tests := map[string]struct {
		body         string
		errExpected  error
		useCaseMock  *useCaseMock
		respExpected response.CreateUrlResponse
	}{
		"when the request body is invalid": {
			body:        `}}{{{}}`,
			errExpected: echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request"),
			useCaseMock: new(useCaseMock),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := NewManageHandler(e, tc.useCaseMock)
			req := httptest.NewRequest(http.MethodPost, "/manage/create", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.CreateUrlRequest(c)
			assert.Equal(t, tc.errExpected, err)

			if err == nil {
				var resp response.CreateUrlResponse
				b, _ := io.ReadAll(rec.Body)
				_ = json.Unmarshal(b, &resp)
				assert.EqualValues(t, tc.respExpected, resp)
			}

			tc.useCaseMock.AssertExpectations(t)
		})
	}
}
