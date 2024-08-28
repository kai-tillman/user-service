package api

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)


// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// Delete user by ID
func (h *APIHandler) DeleteUserById(ctx context.Context, userId string) (Response, error) {
	// TODO: implement the DeleteUserById function to return the following responses

	// return NewResponse(204, {}, "", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"deleteUserById operation has not been implemented yet"}, "application/json", nil), nil
}

// Get user by ID
func (h *APIHandler) GetUserById(ctx context.Context, userId string) (Response, error) {
	// TODO: implement the GetUserById function to return the following responses

	// return NewResponse(200, User{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getUserById operation has not been implemented yet"}, "application/json", nil), nil
}

// Register a new user
func (h *APIHandler) RegisterUser(ctx context.Context, reqBody UserRegistration) (Response, error) {
	// TODO: implement the RegisterUser function to return the following responses

	// return NewResponse(201, User{}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"registerUser operation has not been implemented yet"}, "application/json", nil), nil
}

// Update user by ID
func (h *APIHandler) UpdateUserById(ctx context.Context, userId string, reqBody User) (Response, error) {
	// TODO: implement the UpdateUserById function to return the following responses

	// return NewResponse(200, User{}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"updateUserById operation has not been implemented yet"}, "application/json", nil), nil
}

