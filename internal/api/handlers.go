package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// HandleDeleteUserById handles parsing input to pass to the DeleteUserById operation and sends responses back to the client
func (h *APIHandler) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var userId string
	userId = pathParams["userId"]
	if userId == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'userId'", w)
		return
	}

	response, err := h.DeleteUserById(r.Context(), userId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DeleteUserById returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DeleteUserById was unable to send it's response, err: %s", err)
	}
}

// HandleGetUserById handles parsing input to pass to the GetUserById operation and sends responses back to the client
func (h *APIHandler) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var userId string
	userId = pathParams["userId"]
	if userId == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'userId'", w)
		return
	}

	response, err := h.GetUserById(r.Context(), userId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetUserById returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetUserById was unable to send it's response, err: %s", err)
	}
}

// HandleRegisterUser handles parsing input to pass to the RegisterUser operation and sends responses back to the client
func (h *APIHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := UserRegistration{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'UserRegistration'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.RegisterUser(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("RegisterUser returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("RegisterUser was unable to send it's response, err: %s", err)
	}
}

// HandleUpdateUserById handles parsing input to pass to the UpdateUserById operation and sends responses back to the client
func (h *APIHandler) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var userId string
	userId = pathParams["userId"]
	if userId == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'userId'", w)
		return
	}

	reqBody := User{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'User'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.UpdateUserById(r.Context(), userId, reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("UpdateUserById returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("UpdateUserById was unable to send it's response, err: %s", err)
	}
}

