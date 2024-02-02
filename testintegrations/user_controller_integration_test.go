package testintegrations

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"integration-test/routes"
	"net/http/httptest"
	"strings"

	"testing"
)

func InitUsers() *gin.Engine {
	g, arangoDB := InitIntegration()
	routes.InitializeHttpRoute(g, arangoDB)

	return g
}

func TestUserController_GetUser(t *testing.T) {
	tests := []struct {
		name string

		err                error
		expectedStatusCode int
		total              int
	}{
		{
			name:               "Success",
			total:              100,
			err:                nil,
			expectedStatusCode: 200,
		},
	}

	g := InitUsers()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestUrl := "/api/v1/user"
			requestCall := createGetRequest(requestUrl)
			recorder := httptest.NewRecorder()
			g.ServeHTTP(recorder, requestCall)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
		})
	}
}

func TestUserController_FindUserById(t *testing.T) {
	tests := []struct {
		name               string
		applicationId      string
		expectedStatusCode int
	}{
		{
			name:               "Success",
			applicationId:      "ed914adc-180e-4a74-abc2-82f2ca14cb8f",
			expectedStatusCode: 200,
		},
		{
			name:               "Failed",
			applicationId:      "40ab3d79-5b3e-43dc-9663-9303a04f277d",
			expectedStatusCode: 404,
		},
	}

	g := InitUsers()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestUrl := "/api/v1/user/" + tt.applicationId
			requestCall := createGetRequest(requestUrl)
			recorder := httptest.NewRecorder()
			g.ServeHTTP(recorder, requestCall)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
		})
	}
}

func TestUserController_Create(t *testing.T) {
	tests := []struct {
		name               string
		request            string
		err                error
		expectedStatusCode int
	}{
		{
			name:               "Success",
			request:            "{\n\t\t\t\"name\":\"test\",\n\t\t\t\"email\": \"email@email.com\",\n\t\t\t\"phone_number\": \"091231321\",\n\t\t\t\"address\": \"1234\"\n\t\t}",
			err:                nil,
			expectedStatusCode: 201,
		},
		{
			name:               "Failed Test",
			request:            "{\n    \"name\":\"test\"\n}",
			err:                nil,
			expectedStatusCode: 400,
		},
	}

	g := InitUsers()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestUrl := "/api/v1/user"
			requestCall := createPostRequest(requestUrl, strings.NewReader(tt.request))
			recorder := httptest.NewRecorder()
			g.ServeHTTP(recorder, requestCall)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
		})
	}
}

func TestUserController_Update(t *testing.T) {
	tests := []struct {
		name               string
		applicationId      string
		request            string
		err                error
		expectedStatusCode int
	}{
		{
			name:               "Success",
			applicationId:      "ed914adc-180e-4a74-abc2-82f2ca14cb8f",
			request:            "{\n\t\t\t\"name\":\"test\",\n\t\t\t\"email\": \"email@email.com\",\n\t\t\t\"phone_number\": \"091231321\",\n\t\t\t\"address\": \"1234\"\n\t\t}",
			err:                nil,
			expectedStatusCode: 500,
		},
		{
			name:               "Failed Test",
			applicationId:      "ed914adc-180e-4a74-abc2-82f2ca14cb8f",
			request:            "{\n    \"name\":\"test\"\n}",
			err:                nil,
			expectedStatusCode: 400,
		},
	}

	g := InitUsers()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestUrl := "/api/v1/user/" + tt.applicationId
			requestCall := createPutRequest(requestUrl, strings.NewReader(tt.request))
			recorder := httptest.NewRecorder()
			g.ServeHTTP(recorder, requestCall)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
		})
	}
}

func TestUserController_Delete(t *testing.T) {
	tests := []struct {
		name               string
		applicationId      string
		err                error
		expectedStatusCode int
	}{
		{
			name:               "Success",
			applicationId:      "ed914adc-180e-4a74-abc2-82f2ca14cb8f",
			err:                nil,
			expectedStatusCode: 200,
		},
		{
			name:               "Failed Test",
			applicationId:      "ed914adc-180e-4a74-abc2-82f2ca14cb8e",
			err:                nil,
			expectedStatusCode: 404,
		},
	}

	g := InitUsers()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestUrl := "/api/v1/user/" + tt.applicationId
			requestCall := createDeleteRequest(requestUrl)
			recorder := httptest.NewRecorder()
			g.ServeHTTP(recorder, requestCall)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
		})
	}
}
