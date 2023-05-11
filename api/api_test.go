package api_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hemisto/miniurl/api"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi_AddUrl(t *testing.T) {
	tests := []struct {
		name               string
		payload            string
		handler            api.Handler
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "return status ok and value",
			payload:            `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			handler:            &handlerMockSuccess{str: "testValue"},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash": "testValue"}`,
		},
		{
			name:               "should return bad request and error when hash returns error",
			payload:            `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			handler:            &handlerMockFail{str: "testValue"},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"msg": "error occured while generating hash"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(tc.payload))
			rr := httptest.NewRecorder()

			//ToDo: router
			r := httprouter.New()
			h := tc.handler
			api.Bind(r, h)
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode)

			body, err := io.ReadAll(rr.Result().Body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expectedBody, string(body))
		})
	}

}

type handlerMockSuccess struct {
	str string
}

func (h *handlerMockSuccess) Hash(url string) (hash string, err error) {
	return h.str, nil
}

type handlerMockFail struct {
	str string
}

func (h *handlerMockFail) Hash(url string) (hash string, err error) {
	return h.str, errors.New("error occured while generating hash")
}
