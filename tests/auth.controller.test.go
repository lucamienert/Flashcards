package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucamienert/flashcards/models"
	"github.com/lucamienert/flashcards/utils"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	authController := NewAuthController(db)

	r.POST("/auth/signup", authController.SignUpUser)
	r.POST("/auth/signin", authController.SignInUser)
	r.GET("/auth/refresh", authController.RefreshAccessToken)
	r.POST("/auth/logout", authController.LogoutUser)

	return r
}

func TestSignUpUser(t *testing.T) {
	r := setupRouter()

	tests := []struct {
		name           string
		input          *models.SignUpInput
		expectedStatus int
	}{
		{
			name: "Valid SignUp",
			input: &models.SignUpInput{
				Name:            "Test User",
				Email:           "testuser@example.com",
				Password:        "password123",
				PasswordConfirm: "password123",
				Photo:           "http://example.com/photo.jpg",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Password mismatch",
			input: &models.SignUpInput{
				Name:            "Test User",
				Email:           "testuser@example.com",
				Password:        "password123",
				PasswordConfirm: "wrongpassword",
				Photo:           "http://example.com/photo.jpg",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid email",
			input: &models.SignUpInput{
				Name:            "Test User",
				Email:           "invalid-email",
				Password:        "password123",
				PasswordConfirm: "password123",
				Photo:           "http://example.com/photo.jpg",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := []byte(`{
				"name": "` + tt.input.Name + `",
				"email": "` + tt.input.Email + `",
				"password": "` + tt.input.Password + `",
				"passwordConfirm": "` + tt.input.PasswordConfirm + `",
				"photo": "` + tt.input.Photo + `"
			}`)

			req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestSignInUser(t *testing.T) {
	r := setupRouter()

	// Test setup for the test
	// You can mock the database call to return a user for testing purposes
	mockUser := &models.User{
		Email:    "testuser@example.com",
		Password: utils.HashPassword("password123"), // Mocked password hash
	}

	// Assuming you have a mocked DB call that finds the mockUser for testing.
	// Setup the mock DB call here for the purpose of the test.

	tests := []struct {
		name           string
		input          *models.SignInInput
		expectedStatus int
	}{
		{
			name: "Valid SignIn",
			input: &models.SignInInput{
				Email:    "testuser@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Password",
			input: &models.SignInInput{
				Email:    "testuser@example.com",
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Email",
			input: &models.SignInInput{
				Email:    "invaliduser@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := []byte(`{
				"email": "` + tt.input.Email + `",
				"password": "` + tt.input.Password + `"
			}`)

			req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestLogoutUser(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("POST", "/auth/logout", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}
