package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"http-demo/internal/domain"
	"http-demo/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserServiceMock struct{}

func (s UserServiceMock) CreateUser(user domain.User) (domain.User, error) {
	if user.Name == "bad_user" {
		return domain.User{}, errors.New("Failed to create user")
	}
	newUser := domain.User{
		Id:   1,
		Name: "Oleg",
		Role: "Admin",
	}

	return newUser, nil
}

func (s UserServiceMock) ChangePassword(user domain.User) error {
	return nil
}

func (s UserServiceMock) GetUser(userId int) (domain.User, error) {
	newUser := domain.User{
		Id:   1,
		Name: "Oleg",
		Role: "Admin",
	}

	return newUser, nil
}

func (s UserServiceMock) DeleteUser(userId int) (domain.User, error) {
	return domain.User{}, nil
}

func Test_Create(t *testing.T) {
	service := UserServiceMock{}
	handler := handler.NewUserHandler(service)

	newUser := domain.User{
		Id:   1,
		Name: "Oleg",
		Role: "Admin",
	}

	badUser := domain.User{
		Name: "bad_user",
	}

	testCases := []struct {
		caseName           string
		user               domain.User
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			caseName:           "User successfully created",
			user:               newUser,
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   "User id: 1, name: Oleg, role: Admin, created",
		},
		{
			caseName:           "Failed user creation",
			user:               badUser,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "",
		},
	}

	for _, ts := range testCases {
		t.Run(ts.caseName, func(t *testing.T) {
			newBodyBz, err := json.Marshal(&ts.user)

			if err != nil {
				t.Errorf("Failed to marshal user")
			}

			body := bytes.NewBuffer(newBodyBz)

			r := httptest.NewRequest(http.MethodPost, "/create", body)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(handler.Create)
			h.ServeHTTP(w, r)

			resBody := w.Body.String() // resBody := string(w.Body.Bytes())
			statusCode := w.Code

			// Для передачи параметров в Get
			// r.URL.Query()

			if statusCode != ts.expectedStatusCode {
				t.Errorf("Incorrect status code, expected %d, got: %d", ts.expectedStatusCode, statusCode)
			}

			if resBody != ts.expectedResponse {
				t.Errorf("Incorrect response, expected %s, got: %s", ts.expectedResponse, resBody)
			}
		})
	}
}
